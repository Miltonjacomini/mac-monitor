package main

import (
	"context"
	"embed"
	"fmt"
	"mac-monitor/pkg/collectors"
	"mac-monitor/pkg/engine"
	"mac-monitor/pkg/models"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create engine and collectors
	eng := engine.NewEngine(1 * time.Second)
	eng.RegisterCollector(collectors.NewCPUCollector())
	eng.RegisterCollector(collectors.NewMemoryCollector())
	eng.RegisterCollector(collectors.NewNetworkCollector())
	eng.RegisterCollector(collectors.NewDiskCollector())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go eng.Start(ctx)

	// Create an instance of the app structure
	app := NewApp(eng)

	// Define System Tray Menu
	trayMenu := menu.NewMenu()
	mCPU := trayMenu.AddText("CPU Usage: --", nil, nil)
	mMem := trayMenu.AddText("Mem Used: --", nil, nil)
	trayMenu.AddSeparator()
	trayMenu.AddText("Show Dashboard", nil, func(_ *menu.CallbackData) {
		runtime.WindowShow(app.ctx)
	})
	trayMenu.AddText("Quit", nil, func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "mac-monitor Dashboard",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			// Update Tray in background
			go func() {
				for {
					select {
					case <-ctx.Done():
						return
					case <-time.After(1 * time.Second):
						metrics := eng.GetLatestMetrics()
						if cpuData, ok := metrics["cpu"]; ok {
							cpu := cpuData.Data["cpu"].(models.CPUMetrics)
							mCPU.SetLabel(fmt.Sprintf("CPU Usage: %.1f%%", cpu.TotalUsage))
						}
						if memData, ok := metrics["memory"]; ok {
							mem := memData.Data["memory"].(models.MemoryMetrics)
							mMem.SetLabel(fmt.Sprintf("Mem Used: %.2f GB", float64(mem.Used)/1024/1024/1024))
						}
					}
				}
			}()
		},
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title:   "mac-monitor",
				Message: "Nativa de Monitoramento para macOS",
			},
		},
		Menu: trayMenu,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
