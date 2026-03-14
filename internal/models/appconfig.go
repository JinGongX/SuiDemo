package models

type Lang struct {
	AppMenu Appmenus `json:"Appmenus"`
}

type Appmenus struct {
	About       string `json:"about"`
	Quit        string `json:"quit"`
	Preferences string `json:"preferences"`
	Second      string `json:"second"`
	Version     string `json:"version"`
}
