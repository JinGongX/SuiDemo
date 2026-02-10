package services

import (
	"changeme/internal/models"
	hotkey "changeme/platform"
	"fmt"
	"sync"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type HotkeyService struct {
	store *SuiStore
}

func NewHotkeyService(store *SuiStore) *HotkeyService {
	return &HotkeyService{
		store: store,
	}
}

// 快捷键修改
func (cs *HotkeyService) UpHotkey(id int, key int, modifier int) error {
	var oldKey, oldModifier uint32
	// 先查询旧的 keycode 和 modifier
	cs.store.DB.QueryRow("SELECT keycode, modifiers FROM hotkeys WHERE id = ?", id).Scan(&oldKey, &oldModifier)

	_, err := cs.store.DB.Exec(`
        UPDATE hotkeys 
        SET keycode = ?, modifiers = ? 
        WHERE id = ?
    `, key, modifier, id)
	LoadAndRegisterHotkeysFrom(cs.store, id)
	// 注销旧的热键
	hotkey.UnregisterHotKey(oldKey, oldModifier)
	return err
}

func (cs *HotkeyService) GetHotkeys() ([]models.Hotkey, error) {
	rows, err := cs.store.DB.Query("SELECT id, keycode, modifiers FROM hotkeys")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotkeys []models.Hotkey
	for rows.Next() {
		var hk models.Hotkey
		if err := rows.Scan(&hk.ID, &hk.KeyCode, &hk.Modifiers); err != nil {
			return nil, err
		}
		hotkeys = append(hotkeys, hk)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return hotkeys, nil
}

type hotkeyEntry struct {
	ID       int
	Hotkey   models.Hotkey
	Window   application.Window
	Callback func()
}

var (
	hotkeys = make(map[int]*hotkeyEntry)
	mutex   sync.RWMutex
)

func RegisterHotkeyEntry(id int, keycode, modifiers uint32, win application.Window, cb func()) {
	mutex.Lock()
	defer mutex.Unlock()

	hotkeys[id] = &hotkeyEntry{
		ID: id,
		Hotkey: models.Hotkey{
			KeyCode:   keycode,
			Modifiers: modifiers,
		},
		Window:   win,
		Callback: cb,
	}
	//fmt.Printf("✅ 注册热键 Entry: id=%d (%d, %d)\n", id, keycode, modifiers)
	hotkey.RegisterHotKeyWithCallback(keycode, modifiers, cb)
}

func RegistertargetHotkeyEntry(id int, keycode, modifiers uint32, cb func()) {
	mutex.Lock()
	defer mutex.Unlock()

	hotkeys[id] = &hotkeyEntry{
		ID: id,
		Hotkey: models.Hotkey{
			KeyCode:   keycode,
			Modifiers: modifiers,
		},
		Window:   nil, // 没有窗口绑定
		Callback: cb,
	}
	//fmt.Printf("✅ 注册热键 Entry: id=%d (%d, %d)\n", id, keycode, modifiers)
	hotkey.RegisterHotKeyWithCallback(keycode, modifiers, cb)
}

// 从数据库加载热键并注册
func LoadAndRegisterHotkeysFrom(m *SuiStore, id int) {
	var hk models.Hotkey
	err := m.DB.QueryRow("SELECT keycode, modifiers FROM hotkeys  where id = ?", id).Scan(&hk.KeyCode, &hk.Modifiers)
	if err != nil {
		fmt.Println("Error querying hotkeys:", err)
		return
	}
	mutex.RLock()
	entry, ok := hotkeys[id]
	mutex.RUnlock()

	if !ok {
		fmt.Println("❌ 未找到之前注册的窗口或回调，无法重新注册热键")
		return
	}
	hotkey.UnregisterHotKey(hk.KeyCode, hk.Modifiers) // 先注销旧的热键
	RegisterHotkeyEntry(id, hk.KeyCode, hk.Modifiers, entry.Window, entry.Callback)
}

// 注册窗口回调函数
func RegisterWindowAndCallback(id int, win application.Window, cb hotkey.HotKeyCallback) {
	mutex.Lock()
	hotkeys[id] = &hotkeyEntry{
		ID: id,
		Hotkey: models.Hotkey{
			KeyCode:   0, // 默认值，实际使用时不需要
			Modifiers: 0, // 默认值，实际使用时不需要
		},
		Window:   win,
		Callback: cb,
	}
	mutex.Unlock()
}
