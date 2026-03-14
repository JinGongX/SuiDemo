package services

type AppConfigService struct {
	store *SuiStore
}

// 创建新的AppConfigService实例
func NewAppConfigService(store *SuiStore) *AppConfigService {
	return &AppConfigService{
		store: store,
	}
}

// 根据key获取配置项
func (r *AppConfigService) GetAppConfig(key string) (string, error) {
	row := r.store.DB.QueryRow(`
		SELECT value
		FROM appconfig
		WHERE key = ?
	`, key)

	var value string
	err := row.Scan(&value)
	if err != nil {
		return "en", err
	}
	return value, nil
}

// 设置或更新配置项
func (r *AppConfigService) SetAppConfig(key, value string) error {
	_, err := r.store.DB.Exec(`
		INSERT INTO appconfig (key, type, value, description)
		VALUES (?, 'user', ?, '')
		ON CONFLICT(key) DO UPDATE SET value=excluded.value
	`, key, value)
	return err
}

// 获取应用语言
func (s *AppConfigService) GetLanguage() string {
	lang, err := s.GetAppConfig("language")
	if err != nil || lang == "" {
		return "en"
	}
	return lang
}

// 设置应用语言
func (s *AppConfigService) SetLanguage(lang string) error {
	return s.SetAppConfig("language", lang)
}
