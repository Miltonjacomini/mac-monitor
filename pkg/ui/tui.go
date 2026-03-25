package ui

import (
	"fmt"
	"mac-monitor/pkg/engine"
	"mac-monitor/pkg/models"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			MarginBottom(1)
	
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1).
			MarginBottom(1)

	sectionStyle = lipgloss.NewStyle().
			MarginBottom(1).
			Padding(1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD"))
)

type model struct {
	engine      *engine.Engine
	cpuProgress progress.Model
	memProgress progress.Model
	width       int
	height      int
	lastMetrics map[string]models.MetricsPayload
}

func NewModel(eng *engine.Engine) model {
	return model{
		engine:      eng,
		cpuProgress: progress.New(progress.WithDefaultGradient()),
		memProgress: progress.New(progress.WithDefaultGradient()),
	}
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.cpuProgress.Width = msg.Width - 10
		m.memProgress.Width = msg.Width - 10
		return m, nil
	case tickMsg:
		m.lastMetrics = m.engine.GetLatestMetrics()
		return m, tick()
	}
	return m, nil
}

func (m model) View() string {
	if m.lastMetrics == nil {
		return "Initialing collectors..."
	}

	var s strings.Builder

	s.WriteString(titleStyle.Render("mac-monitor v0.1.0"))
	s.WriteString("\n")

	// CPU Section
	if cpuData, ok := m.lastMetrics["cpu"]; ok {
		cpu := cpuData.Data["cpu"].(models.CPUMetrics)
		s.WriteString(headerStyle.Render(" CPU Usage "))
		s.WriteString(fmt.Sprintf(" %d MHz | Total: %.1f%%\n", cpu.Frequency, cpu.TotalUsage))
		s.WriteString(m.cpuProgress.ViewAs(cpu.TotalUsage / 100))
		s.WriteString("\n\n")
	}

	// Memory Section
	if memData, ok := m.lastMetrics["memory"]; ok {
		mem := memData.Data["memory"].(models.MemoryMetrics)
		// Note: We should ideally have the total physical RAM.
		// For now, let's just show raw numbers for now.

		// Actually, let's just show raw numbers for now.
		s.WriteString(headerStyle.Render(" Memory "))
		s.WriteString(fmt.Sprintf(" Used: %.2f GB | Wired: %.2f GB | Pressure: %.0f\n", 
			float64(mem.Used)/1024/1024/1024, 
			float64(mem.Wired)/1024/1024/1024,
			mem.Pressure))
	}

	// Network Ports Section
	if netData, ok := m.lastMetrics["network"]; ok {
		net := netData.Data["network"].(models.NetworkMetrics)
		s.WriteString("\n")
		s.WriteString(headerStyle.Render(" Open Network Ports "))
		s.WriteString(fmt.Sprintf(" Count: %d\n", len(net.OpenPorts)))
		
		// Sort and show top 5
		ports := append([]models.NetworkPort{}, net.OpenPorts...)
		sort.Slice(ports, func(i, j int) bool {
			return ports[i].Port < ports[j].Port
		})

		for i := 0; i < 5 && i < len(ports); i++ {
			p := ports[i]
			s.WriteString(fmt.Sprintf(" [%s] %d \t %s\n", p.Protocol, p.Port, p.Process))
		}
	}

	// Disk Section
	if diskData, ok := m.lastMetrics["disk"]; ok {
		disk := diskData.Data["disk"].(models.DiskMetrics)
		s.WriteString("\n")
		s.WriteString(headerStyle.Render(" Disk Volumes "))
		for _, v := range disk.Volumes {
			if v.MountPoint == "/" || v.MountPoint == "/System/Volumes/Data" {
				percent := float64(v.Used) / float64(v.Total)
				s.WriteString(fmt.Sprintf(" %s [%s]: %.1f%% used (%.1f/%.1f GB)\n", 
					v.MountPoint, v.FileSystem, percent*100, 
					float64(v.Used)/1024/1024/1024, 
					float64(v.Total)/1024/1024/1024))
			}
		}
	}

	s.WriteString("\nPress 'q' to quit.")
	return s.String()
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
