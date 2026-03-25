package models

import "time"

// MetricsPayload is the normalized payload for all system metrics.
type MetricsPayload struct {
	Timestamp time.Time
	Data      map[string]interface{}
}

// ProcessInfo contains basic info about a running process.
type ProcessInfo struct {
	PID     int
	Name    string
	CPU     float64
	Memory  uint64
}

// CPUMetrics represents the state of CPU resources.
type CPUMetrics struct {
	TotalUsage float64
	PerCore    []float64
	Frequency  int64 // MHz
	TopProcs   []ProcessInfo
}

// MemoryMetrics represents the state of memory resources.
type MemoryMetrics struct {
	Used       uint64
	Wired      uint64
	Compressed uint64
	Pressure   float64 // 0-100%
	SwapUsed   uint64
}

// NetworkPort represents an open network port.
type NetworkPort struct {
	Port     int
	Protocol string // "TCP" | "UDP"
	PID      int
	Process  string
}

// NetworkMetrics represents the state of network resources.
type NetworkMetrics struct {
	Interfaces []InterfaceMetrics
	OpenPorts  []NetworkPort
}

// VolumeMetrics represents stats for a disk volume.
type VolumeMetrics struct {
	MountPoint string
	FileSystem string
	Total      uint64
	Used       uint64
	Free       uint64
}

// DiskMetrics represents the state of disk resources.
type DiskMetrics struct {
	Volumes []VolumeMetrics
	ReadBytesRate  float64
	WriteBytesRate float64
}

// InterfaceMetrics contains stats for a specific network interface.
type InterfaceMetrics struct {
	Name      string
	BytesIn   uint64
	BytesOut  uint64
	PacketsIn uint64
	PacketsOut uint64
}
