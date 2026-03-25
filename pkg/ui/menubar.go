package ui

import (
	"fmt"
	"mac-monitor/pkg/engine"
	"mac-monitor/pkg/models"
	"time"

	"github.com/getlantern/systray"
)

// MenuBar handles the macOS top bar integration.
type MenuBar struct {
	engine *engine.Engine
}

func NewMenuBar(eng *engine.Engine) *MenuBar {
	return &MenuBar{engine: eng}
}

func (mb *MenuBar) Run() {
	systray.Run(mb.onReady, mb.onExit)
}

func (mb *MenuBar) onReady() {
	systray.SetTitle("M")
	systray.SetTooltip("mac-monitor")

	mCPU := systray.AddMenuItem("CPU: --", "Current CPU Usage")
	mMem := systray.AddMenuItem("Mem: --", "Current Memory Usage")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit mac-monitor")

	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			case <-time.After(1 * time.Second):
				metrics := mb.engine.GetLatestMetrics()
				
				if cpuData, ok := metrics["cpu"]; ok {
					cpu := cpuData.Data["cpu"].(models.CPUMetrics)
					title := fmt.Sprintf("CPU: %.0f%%", cpu.TotalUsage)
					systray.SetTitle(title)
					mCPU.SetTitle(fmt.Sprintf("CPU Usage: %.1f%%", cpu.TotalUsage))
				}

				if memData, ok := metrics["memory"]; ok {
					mem := memData.Data["memory"].(models.MemoryMetrics)
					mMem.SetTitle(fmt.Sprintf("Mem Used: %.2f GB (Pressure: %.0f)", 
						float64(mem.Used)/1024/1024/1024, mem.Pressure))
				}
			}
		}
	}()
}

func (mb *MenuBar) onExit() {
	// Cleanup if needed
}
