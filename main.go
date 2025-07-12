package main

import (
	"changeme/services"
	"embed"
	_ "embed"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

type AppService struct {
	App          *application.App
	secondWindow *application.WebviewWindow
	mainwindow   *application.WebviewWindow
}

func (a *AppService) SetApp(app *application.App) {
	a.App = app
}

func (a *AppService) OpenSecondWindow() {
	if a.secondWindow != nil {
		fmt.Println("[DEBUG] secondWindow is not nil")
		a.secondWindow.Show()
	} else {
		fmt.Println("[ERROR] secondWindow is nil")
	}
}

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {
	appservice := &AppService{}

	suiService, errt := services.NewSuiStore()
	if errt != nil {
		// 处理错误，比如日志或退出
	}

	//fmt.Println(clipboardService)
	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "myproject",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(&GreetService{}),
			application.NewService(appservice),
			application.NewService(suiService),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "Window 1",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
	})
	//add secondWindow
	secondWindow := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title:     "second",
		Name:      "second",
		Width:     440,
		Height:    620,
		Frameless: true, // ✅ 去除边框
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			//Backdrop:                application.MacBackdropTranslucent,
			TitleBar: application.MacTitleBarHidden,
			Backdrop: application.MacBackdropTransparent, // 可选：背景透明
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/second",
	})
	secondWindow.Hide()
	appservice.secondWindow = secondWindow
	appservice.SetApp(app)

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		// for {
		// 	now := time.Now().Format(time.RFC1123)
		// 	app.EmitEvent("time", now)
		// 	time.Sleep(time.Second)
		// }
	}()

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
