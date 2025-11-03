package main

import (
	"context"
	"embed"

	"github.com/OscarMitchell/echo/src/server"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	serverHandler := server.NewServerHandler()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Echo",
		Width:  850,
		Height: 575,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		OnStartup: func(ctx context.Context) {
			serverHandler.OnStartup(ctx)
		},
		OnBeforeClose: func(ctx context.Context) (prevent bool) {
			if serverHandler.IsUp() {
				dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
					Type:    runtime.QuestionDialog,
					Title:   "Quit?",
					Message: "The Server is still running.\nAre you sure you want to quit?",
				})
				if err != nil {
					return false
				}
				return dialog == "yes"
			}
			return false
		},
		OnShutdown: serverHandler.OnShutdown,
		Bind: []any{
			serverHandler,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
