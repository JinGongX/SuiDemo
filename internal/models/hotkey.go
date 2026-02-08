package models

type Hotkey struct {
	ID        int    `json:"id"`        // 热键ID
	KeyCode   uint32 `json:"keycode"`   // 键码
	Modifiers uint32 `json:"modifiers"` // 修饰键
}
