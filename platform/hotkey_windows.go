//go:build windows
// +build windows

package platform

import (
	"fmt"
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
