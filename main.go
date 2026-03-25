package main

import (
	"context"
	"flag"
	"fmt"
	"mac-monitor/pkg/collectors"
	"mac-monitor/pkg/engine"
	"mac-monitor/pkg/ui"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	useTUI := flag.Bool("tui", false, "Run in TUI mode")
	flag.Parse()

	eng := engine.NewEngine(1 * time.Second)
	
	eng.RegisterCollector(collectors.NewCPUCollector())
	eng.RegisterCollector(collectors.NewMemoryCollector())
	eng.RegisterCollector(collectors.NewNetworkCollector())
	eng.RegisterCollector(collectors.NewDiskCollector())
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		eng.Start(ctx)
	}()

	if *useTUI {
		p := tea.NewProgram(ui.NewModel(eng), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	} else {
		mb := ui.NewMenuBar(eng)
		mb.Run()
	}
}
