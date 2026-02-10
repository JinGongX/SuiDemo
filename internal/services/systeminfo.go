package services

import (
	"runtime"
)

type SystemInfo struct{}

func NewSystemInfo() *SystemInfo {
	return &SystemInfo{}
}

func (s *SystemInfo) GetOS() string {
	return runtime.GOOS // 返回 "darwin" / "windows" / "linux"
}

// 根据系统初始化默认热键
func OSinithotkeys() string {
	var sqlhotkeys = `
			INSERT INTO hotkeys (keycode, modifiers, description, target)
			VALUES 
				(80, 1, '剪贴板窗口快捷键', '1'),
				(77, 1, '偏好设置快捷键 ', '2'); `

	if runtime.GOOS == "darwin" {
		sqlhotkeys = `
			INSERT INTO hotkeys (keycode, modifiers, description, target)
			VALUES 
				(34, 768, '剪贴板窗口快捷键', '1'),
				(46, 768, '偏好设置快捷键 ', '2'); `
	}
	return sqlhotkeys
}
