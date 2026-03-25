package engine

import (
	"context"
	"mac-monitor/pkg/models"
	"sync/atomic"
	"testing"
	"time"
)

type mockCollector struct {
	collectCount int32
}

func (m *mockCollector) Collect(ctx context.Context) (models.MetricsPayload, error) {
	atomic.AddInt32(&m.collectCount, 1)
	return models.MetricsPayload{
		Timestamp: time.Now(),
		Data:      map[string]interface{}{"count": m.collectCount},
	}, nil
}

func (m *mockCollector) Name() string {
	return "mock"
}

func TestEngine_CollectAll(t *testing.T) {
	eng := NewEngine(100 * time.Millisecond)
	mock := &mockCollector{}
	eng.RegisterCollector(mock)

	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	eng.Start(ctx)

	count := atomic.LoadInt32(&mock.collectCount)
	if count < 2 {
		t.Errorf("expected at least 2 collections, got %d", count)
	}

	metrics := eng.GetLatestMetrics()
	if _, ok := metrics["mock"]; !ok {
		t.Error("expected metrics for 'mock' collector")
	}
}
