//go:build windows
// +build windows

package platform

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"

	"golang.design/x/hotkey"
)

const (
	CF_UNICODETEXT = 13
	CF_DIB         = 8
	GMEM_MOVEABLE  = 0x0002
	//WM_HOTKEY      = 0x0312
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	procSetForeground = user32.NewProc("SetForegroundWindow")
	procShowWindow    = user32.NewProc("ShowWindow")
	procGetForeground = user32.NewProc("GetForegroundWindow")
)

const (
	SW_RESTORE = 9
)

func SetForegroundWindow(hwnd uintptr) {
	procSetForeground.Call(hwnd)
}

func ShowWindow(hwnd uintptr, nCmdShow int) {
	procShowWindow.Call(hwnd, uintptr(nCmdShow))
}

func GetForegroundWindow() uintptr {
	hwnd, _, _ := procGetForeground.Call()
	return hwnd
}

func ActivateApp() {
	hwnd := GetForegroundWindow() // 或你自己保存的主窗口句柄
	if hwnd != 0 {
		ShowWindow(hwnd, SW_RESTORE)
		SetForegroundWindow(hwnd)
	}
}

// 改用"golang.design/x/hotkey" 调用热键，直接调用windows API会无法监听热键消息
type HotKeyCallback func()

type hotkeyEntry struct {
	hk  *hotkey.Hotkey
	key uint32
	mod uint32
}

var (
	callbacks  = make(map[int]hotkeyEntry) //make(map[int]*hotkey.Hotkey)
	cbHandlers = make(map[int]HotKeyCallback)
	idCounter  = 1
	mutex      sync.RWMutex
	startOnce  sync.Once
)

// RegisterHotKeyWithCallback 注册热键并绑定回调（使用 golang.design/x/hotkey 实现）
func RegisterHotKeyWithCallback(keyCode, modifiers uint32, cb HotKeyCallback) {
	mod := toHotkeyModifiers(modifiers)
	key := toHotkeyKey(keyCode)

	hk := hotkey.New(mod, key)
	err := hk.Register()
	if err != nil {
		fmt.Printf("❌ 热键注册失败: %v\n", err)
		return
	}

	localID := idCounter
	idCounter++

	mutex.Lock()
	callbacks[localID] = hotkeyEntry{hk: hk, key: keyCode, mod: modifiers} // hk
	cbHandlers[localID] = cb
	mutex.Unlock()

	// 启动监听线程
	go func(id int, h *hotkey.Hotkey, cb HotKeyCallback) {
		for {
			<-h.Keydown()
			fmt.Printf("✅ 热键触发 (id=%d)\n", id)
			cb()
		}
	}(localID, hk, cb)
	// 仅启动一次监听线程
	// 启动统一监听线程（只启动一次）
	//startOnce.Do(startListener)

	fmt.Printf("✈️ 注册热键 (id=%d, key=%v, mod=%v)\n", localID, key, mod)

}

// UnregisterHotKey 注销所有已注册热键
func UnregisterHotKey(_keyCode, _modifiers uint32) {
	mutex.Lock()
	defer mutex.Unlock()

	for id, hk := range callbacks {
		if _keyCode == hk.key && _modifiers == hk.mod {
			_ = hk.hk.Unregister() // 注销热键
			delete(callbacks, id)  // 删除回调
			//delete(cbHandlers, id) // 删除回调处理器
			fmt.Printf("🧹 注销热键 (id=%d, key=%v ", id, _keyCode)
		}
	}

}

func toHotkeyModifiers(modifiers uint32) []hotkey.Modifier {
	var mods []hotkey.Modifier
	if modifiers&1 != 0 {
		mods = append(mods, hotkey.ModAlt)
	}
	if modifiers&2 != 0 {
		mods = append(mods, hotkey.ModCtrl)
	}
	if modifiers&4 != 0 {
		mods = append(mods, hotkey.ModShift)
	}
	if modifiers&8 != 0 {
		mods = append(mods, hotkey.ModWin)
	}
	return mods
}

func toHotkeyKey(keyCode uint32) hotkey.Key {
	return hotkey.Key(keyCode)
}

func SimulateCmdC() {

}

func HideDock() {
	// Windows 不需要隐藏 Dock 图标
	// 但可以在这里添加其他逻辑
	fmt.Println("Windows 平台不支持隐藏 Dock 图标")
}

// OCR功能 Windows OCR实现较复杂，暂时使用PowerShell脚本调用Windows OCR API进行识别 （需要Windows 10或更高版本）
const ocrScript = `
[Console]::OutputEncoding=[System.Text.Encoding]::UTF8
$OutputEncoding=[System.Text.Encoding]::UTF8
Add-Type -AssemblyName System.Runtime.WindowsRuntime

[void][Windows.Graphics.Imaging.BitmapDecoder, Windows.Graphics, ContentType=WindowsRuntime]
[void][Windows.Graphics.Imaging.SoftwareBitmap, Windows.Graphics, ContentType=WindowsRuntime]
[void][Windows.Media.Ocr.OcrEngine, Windows.Foundation, ContentType=WindowsRuntime]
[void][Windows.Media.Ocr.OcrResult, Windows.Foundation, ContentType=WindowsRuntime]
[void][Windows.Storage.Streams.InMemoryRandomAccessStream, Windows.Storage, ContentType=WindowsRuntime]
[void][Windows.Storage.Streams.DataWriter, Windows.Storage, ContentType=WindowsRuntime]

function Await($WinRtTask, $ResultType) {
    $methods = [System.WindowsRuntimeSystemExtensions].GetMethods() | Where-Object {
        $_.Name -eq 'AsTask' -and $_.IsGenericMethodDefinition -and $_.GetParameters().Count -eq 1
    }
    $method = $methods[0].MakeGenericMethod($ResultType)
    $netTask = $method.Invoke($null, @($WinRtTask))
    $netTask.Wait(-1) | Out-Null
    $netTask.Result
}

# 用 System.IO 直接读字节，不依赖 StorageFile
$bytes = [System.IO.File]::ReadAllBytes('{{IMAGE_PATH}}')

# 写入 InMemoryRandomAccessStream
$stream = [Windows.Storage.Streams.InMemoryRandomAccessStream]::new()
$writer = [Windows.Storage.Streams.DataWriter]::new($stream)
$writer.WriteBytes($bytes)
$storeTask = $writer.StoreAsync()
$storeAsTask = [System.WindowsRuntimeSystemExtensions].GetMethods() | Where-Object {
    $_.Name -eq 'AsTask' -and $_.IsGenericMethodDefinition -and $_.GetParameters().Count -eq 1
} | Select-Object -First 1
$storeNetTask = $storeAsTask.MakeGenericMethod([uint32]).Invoke($null, @($storeTask))
$storeNetTask.Wait(-1) | Out-Null
$stream.Seek(0)

$decoder = Await ([Windows.Graphics.Imaging.BitmapDecoder]::CreateAsync($stream)) ([Windows.Graphics.Imaging.BitmapDecoder])
$bitmap  = Await ($decoder.GetSoftwareBitmapAsync()) ([Windows.Graphics.Imaging.SoftwareBitmap])
$engine  = [Windows.Media.Ocr.OcrEngine]::TryCreateFromUserProfileLanguages()
$result  = Await ($engine.RecognizeAsync($bitmap)) ([Windows.Media.Ocr.OcrResult])
Write-Output $result.Text
`

func RecognizeImageBase64(base64str string) error {
	if idx := strings.Index(base64str, ","); idx != -1 {
		base64str = base64str[idx+1:]
	}

	data, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		return fmt.Errorf("base64 decode failed: %w", err)
	}

	tmp, err := os.CreateTemp("", "ocr-*.png")
	if err != nil {
		return fmt.Errorf("create temp file failed: %w", err)
	}
	imgPath := tmp.Name()
	tmp.Write(data)
	tmp.Close()
	defer os.Remove(imgPath)

	imgPath = strings.ReplaceAll(imgPath, "/", `\`)
	script := strings.ReplaceAll(ocrScript, "{{IMAGE_PATH}}", imgPath)

	cmd := exec.Command("powershell",
		"-NoProfile",
		"-NonInteractive",
		"-ExecutionPolicy", "Bypass",
		"-Command", script,
	)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		return fmt.Errorf("powershell error: %w\nstderr: %s", err, stderr.String())
	}

	text := strings.TrimSpace(stdout.String())
	if text == "" {
		return errors.New("OCR returned empty string")
	}

	return nil
}
