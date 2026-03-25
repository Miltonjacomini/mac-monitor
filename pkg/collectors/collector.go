package collectors

import (
	"context"
	"mac-monitor/pkg/models"
)

// Collector is the base interface for all metric collectors.
type Collector interface {
	// Collect captures the current system metrics.
	Collect(ctx context.Context) (models.MetricsPayload, error)
	// Name returns the collector's identifier (cpu, mem, net, etc.).
	Name() string
}
