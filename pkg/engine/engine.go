package engine

import (
	"context"
	"mac-monitor/pkg/collectors"
	"mac-monitor/pkg/models"
	"sync"
	"time"
)

// Engine orchestrates the metrics collection cycle.
type Engine struct {
	collectors []collectors.Collector
	interval   time.Duration
	stopChan   chan struct{}
	mu         sync.RWMutex
	metrics    map[string]models.MetricsPayload
}

// NewEngine creates a new metrics collection engine.
func NewEngine(interval time.Duration) *Engine {
	return &Engine{
		interval: interval,
		stopChan: make(chan struct{}),
		metrics:  make(map[string]models.MetricsPayload),
	}
}

// RegisterCollector adds a new collector to the engine.
func (e *Engine) RegisterCollector(c collectors.Collector) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.collectors = append(e.collectors, c)
}

// Start begins the collection loop.
func (e *Engine) Start(ctx context.Context) {
	ticker := time.NewTicker(e.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			e.collectAll(ctx)
		case <-e.stopChan:
			return
		case <-ctx.Done():
			return
		}
	}
}

// Stop halts the collection loop.
func (e *Engine) Stop() {
	close(e.stopChan)
}

// GetLatestMetrics returns the most recently collected metrics.
func (e *Engine) GetLatestMetrics() map[string]models.MetricsPayload {
	e.mu.RLock()
	defer e.mu.RUnlock()
	
	// Return a copy to avoid race conditions
	copyMap := make(map[string]models.MetricsPayload)
	for k, v := range e.metrics {
		copyMap[k] = v
	}
	return copyMap
}

func (e *Engine) collectAll(ctx context.Context) {
	var wg sync.WaitGroup
	e.mu.RLock()
	currentCollectors := make([]collectors.Collector, len(e.collectors))
	copy(currentCollectors, e.collectors)
	e.mu.RUnlock()

	for _, c := range currentCollectors {
		wg.Add(1)
		go func(col collectors.Collector) {
			defer wg.Done()
			payload, err := col.Collect(ctx)
			if err != nil {
				// For now, just log or skip on error
				return
			}
			e.mu.Lock()
			e.metrics[col.Name()] = payload
			e.mu.Unlock()
		}(c)
	}
	wg.Wait()
}
