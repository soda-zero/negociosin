package main

import (
	"changeme/backend"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := backend.NewApp()
    defer app.Close()

	// Create application with options
	err := wails.Run(&options.App   {
		Title:     "productisusi",
		Width:     800,
		Height:    600,
		MinWidth:  600,
		MinHeight: 480,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:        backend.NewApp().Startup,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
