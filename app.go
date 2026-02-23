// app.go
package main

import (
	"context"
	"embed"
	"galaxies-burn-rate/internal/data"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx        context.Context
	gameDataFS embed.FS
}

func NewApp(fs embed.FS) *App {
	return &App{
		gameDataFS: fs,
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// CreateNewGame triggers the extraction pattern for a new save
func (a *App) CreateNewGame(captainName string, shipName string) string {
	saveID, err := data.CreateNewSave(a.gameDataFS, captainName, shipName)
	if err != nil {
		return ""
	}
	return saveID
}

// LoadGameData fetches all YAMLs from a specific save folder
func (a *App) LoadGameData(saveID string) map[string]string {
	return data.LoadSaveGame(saveID)
}

// GetAvailableSaves exposes the directory scanner to the frontend
func (a *App) GetAvailableSaves() []data.SaveMeta {
	return data.GetAvailableSaves()
}

// --- WINDOW MANAGEMENT ---

// SetDisplayMode handles windowed vs fullscreen
func (a *App) SetDisplayMode(mode string) {
	if mode == "fullscreen" || mode == "borderless" {
		runtime.WindowFullscreen(a.ctx)
	} else {
		runtime.WindowUnfullscreen(a.ctx)
	}
}

// SetResolution resizes the window explicitly
func (a *App) SetResolution(width int, height int) {
	runtime.WindowSetSize(a.ctx, width, height)
	runtime.WindowCenter(a.ctx)
}

func (a *App) QuitGame() {
	runtime.Quit(a.ctx)
}
