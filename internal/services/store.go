package services

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type SuiStore struct {
	DB *sql.DB
}

// 获取安全的数据库路径
func getSafeDBPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appDir := filepath.Join(configDir, "SuiNest")
	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		return "", err
	}
	return filepath.Join(appDir, "suidemo.db"), nil
}

// 初始化数据库
func NewSuiStore() (*SuiStore, error) {
	dbPath, err := getSafeDBPath()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	err = Migrate(db)
	if err != nil {
		return nil, err
	}
	return &SuiStore{DB: db}, nil
}

func (cs *SuiStore) Close() {
	cs.DB.Close()
}

func (cs *SuiStore) Start() error {
	// 这里可以初始化数据库或其它启动逻辑
	return nil
}

func (cs *SuiStore) Stop() error {
	cs.Close()
	return nil
}

func Migrate(db *sql.DB) error {
	fmt.Println("🚀 开始数据库迁移...")
	// 创建表
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS hotkeys (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    keycode INTEGER NOT NULL,
	    modifiers INTEGER NOT NULL,
	    description TEXT,
		target TEXT 
	);

	CREATE TABLE IF NOT EXISTS  appconfig (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT UNIQUE,
		type TEXT,
		value TEXT,
		description TEXT,
		state INTEGER DEFAULT 0
	);
	`)
	if err != nil {
		return err
	}
	//  初始化默认数据
	err = InitDefaultData(db)
	if err != nil {
		return err
	}
	fmt.Println("✅ 数据库迁移完成")
	return nil
}

func InitDefaultData(db *sql.DB) error {
	// 初始化配置
	if _, err := db.Exec(OSinitAppConfig()); err != nil {
		return err
	}
	// 初始化热键
	if _, err := db.Exec(OSinithotkeys()); err != nil {
		return err
	}
	return nil
}
