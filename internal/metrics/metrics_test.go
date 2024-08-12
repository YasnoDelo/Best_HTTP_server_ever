package metrics_test

import (
	"testing"

	"github.com/YasnoDelo/Best_HTTP_server_ever/internal/metrics"
)

func TestUpdateMetrics(t *testing.T) {
	m := metrics.NewMetrics()
	m.UpdateMetrics()

	// Проверяем, что значения метрик не равны нулю
	if m.Alloc == 0 {
		t.Error("Alloc should not be 0 after UpdateMetrics")
	}
	if m.PollCount == 0 {
		t.Error("PollCount should be greater than 0 after UpdateMetrics")
	}
}

func TestRandomValue(t *testing.T) {
	m := metrics.NewMetrics()
	m.UpdateMetrics()

	// Проверяем, что RandomValue находится в диапазоне от 0 до 1
	if m.RandomValue < 0 || m.RandomValue > 1 {
		t.Errorf("RandomValue out of range: %f", m.RandomValue)
	}
}
