package main

import (
	"context"
	"mac-monitor/pkg/engine"
	"mac-monitor/pkg/models"
)

// App struct
type App struct {
	ctx    context.Context
	engine *engine.Engine
}

// NewApp creates a new App application struct
func NewApp(eng *engine.Engine) *App {
	return &App{
		engine: eng,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetMetrics returns the latest metrics from the engine
func (a *App) GetMetrics() map[string]models.MetricsPayload {
	return a.engine.GetLatestMetrics()
}
