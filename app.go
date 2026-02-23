// app.go
package main

import (
	"context"
	"embed"
	"galaxies-burn-rate/internal/data"
)

// App struct
type App struct {
	ctx        context.Context
	gameDataFS embed.FS
}

// NewApp creates a new App application struct
func NewApp(fs embed.FS) *App {
	return &App{
		gameDataFS: fs,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetGameData is exposed to the Vue frontend
func (a *App) GetGameData() map[string]string {
	return data.LoadEmbeddedGameData(a.gameDataFS)
}
