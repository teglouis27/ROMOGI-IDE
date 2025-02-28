package main

import (
	"context"
	"embed"
	"log"

	"IDE_latest/backend/handlers"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
)

var assets embed.FS

type App struct {
	AuthHandler *handlers.AuthHandler
}

// NewApp initializes the application with necessary components.
func NewApp() *App {
	db, err := handlers.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	return &App{
		AuthHandler: handlers.NewAuthHandler(db),
	}
}

func (a *App) Startup(ctx context.Context) {
	log.Println("Application starting up")
	// Optionally start additional services here
}

func (a *App) Shutdown(ctx context.Context) {
	handlers.CloseDB()
	log.Println("Application shutting down")
	// Ensure all cleanup is handled here, such as closing any open connections
}

func main() {

	app := NewApp()

	err := wails.Run(&options.App{
		Title:            "Romogi IDE",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		OnShutdown:       app.Shutdown,
		LogLevel:         logger.DEBUG,
		Bind:             []interface{}{app.AuthHandler},
	})

	if err != nil {
		log.Fatal(err)
	}
}
