// main.go
package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"galaxies-burn-rate/internal/data"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed game_data/*.yaml
var gameDataFS embed.FS

func main() {
	// 1. App initialization. We pass the embedded FS so the App struct can hold onto it.
	app := NewApp(gameDataFS)

	// 2. Load settings using your decoupled data package
	settings := data.LoadClientSettings(gameDataFS)

	// 3. Configure the Wails Window
	err := wails.Run(&options.App{
		Title:      "Galaxies: Burn Rate",
		Width:      settings.Display.Width,
		Height:     settings.Display.Height,
		Fullscreen: settings.Display.Fullscreen,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
