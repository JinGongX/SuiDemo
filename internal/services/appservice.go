package services

import (
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

type AppService struct {
	App          *application.App
	SecondWindow *application.WebviewWindow
	Mainwindow   *application.WebviewWindow
}

var HotkeyShowSecondWindow = 1 // 显示第二个窗口的快捷键ID

func (a *AppService) SetApp(app *application.App) {
	a.App = app

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "main",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		//BackgroundColour: application.NewRGB(27, 38, 54),
		URL: "/",
	})
	//add secondWindow
	secondWindow := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "second",
		Name:   "second",
		Width:  340,
		Height: 320,
		//Frameless: true, // ✅ 去除边框
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			//Backdrop:                application.MacBackdropTranslucent,
			TitleBar: application.MacTitleBarHidden,
			Backdrop: application.MacBackdropTransparent, // 可选：背景透明
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/#/second",
	})
	secondWindow.Hide()
	a.SecondWindow = secondWindow
	// Register the hotkey to show and focus the second window
	RegisterWindowAndCallback(HotkeyShowSecondWindow, a.SecondWindow, func() {
		if a.SecondWindow == nil {
			return
		}
		//screen, _ := secondWindow.GetScreen()                            // 获取屏幕信息
		//secondWindow.SetPosition((screen.X+screen.Size.Width-340)*2, 10) //+10
		//hotkey.ActivateApp()
		a.SecondWindow.Show()
		a.SecondWindow.Focus() // 聚焦窗口
	})

	// 拦截关闭事件，改为隐藏窗口
	a.SecondWindow.RegisterHook(events.Common.WindowClosing, func(e *application.WindowEvent) {
		// Prevent the window from closing
		a.SecondWindow.Hide()
		e.Cancel() // 取消关闭事件
	})

}

func (a *AppService) OpenSecondWindow() {
	if a.SecondWindow != nil {
		a.SecondWindow.Show()
	}
}
