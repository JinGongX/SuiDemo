//go:build darwin
// +build darwin

package platform

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework Carbon -framework Vision  -framework Foundation
#include <stdlib.h>
#include <AppKit/AppKit.h>
#include "system.h"
*/
import "C"
import (
	"fmt"
	"sync"
)

type HotKeyCallback func()

var (
	callbacks = make(map[string]HotKeyCallback)
	mutex     sync.RWMutex
)

//export hotKeyCallback
func hotKeyCallback(keyCode, modifiers C.uint) {
	id := fmt.Sprintf("%d_%d", uint32(keyCode), uint32(modifiers))

	mutex.RLock()
	cb, exists := callbacks[id]
	mutex.RUnlock()

	if exists {
		//fmt.Printf("✅ 触发热键: %s\n", id)
		cb()
	} else {
		fmt.Printf("❌ 找不到回调: %s\n", id)
	}
}

// RegisterHotKeyWithCallback 注册热键并绑定 Go 回调
func RegisterHotKeyWithCallback(keyCode, modifiers uint32, cb HotKeyCallback) {
	id := fmt.Sprintf("%d_%d", keyCode, modifiers)

	mutex.Lock()
	callbacks[id] = cb
	mutex.Unlock()
	//fmt.Printf("✈️ Registering hotkey: %d %d\n", keyCode, modifiers)
	C.RegisterHotKeyDynamic(C.uint(keyCode), C.uint(modifiers))
}

// 激活当前应用程序
func ActivateApp() {
	C.NSAppActivateIgnoringOtherApps() // 激活当前应用程序
}

// UnregisterHotKey 注销热键（Go -> C）
func UnregisterHotKey(keyCode, modifiers uint32) {
	id := fmt.Sprintf("%d_%d", keyCode, modifiers)

	mutex.Lock()
	delete(callbacks, id)
	mutex.Unlock()

	//fmt.Printf("🧹 Unregistering hotkey: %d %d\n", keyCode, modifiers)
	C.UnregisterHotKey(C.uint(keyCode), C.uint(modifiers))
}

func HideDock() {
	C.HideDockIcon()
}
