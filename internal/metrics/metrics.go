package metrics

import (
	"math/rand"
	"runtime"
	"sync/atomic"
)

// Gauge represents a float64 metric.
type Gauge float64

// Counter represents an int64 metric.
type Counter int64

// Metrics holds the runtime metrics.
type Metrics struct {
	Alloc         Gauge
	BuckHashSys   Gauge
	Frees         Gauge
	GCCPUFraction Gauge
	GCSys         Gauge
	HeapAlloc     Gauge
	HeapIdle      Gauge
	HeapInuse     Gauge
	HeapObjects   Gauge
	HeapReleased  Gauge
	HeapSys       Gauge
	LastGC        Gauge
	Lookups       Gauge
	MCacheInuse   Gauge
	MCacheSys     Gauge
	MSpanInuse    Gauge
	MSpanSys      Gauge
	Mallocs       Gauge
	NextGC        Gauge
	NumForcedGC   Gauge
	NumGC         Gauge
	OtherSys      Gauge
	PauseTotalNs  Gauge
	StackInuse    Gauge
	StackSys      Gauge
	Sys           Gauge
	TotalAlloc    Gauge

	PollCount   Counter
	RandomValue Gauge
}

// NewMetrics creates a new Metrics object.
func NewMetrics() *Metrics {
	return &Metrics{}
}

// UpdateMetrics updates the metrics values.
func (m *Metrics) UpdateMetrics() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	m.Alloc = Gauge(memStats.Alloc)
	m.BuckHashSys = Gauge(memStats.BuckHashSys)
	m.Frees = Gauge(memStats.Frees)
	m.GCCPUFraction = Gauge(memStats.GCCPUFraction)
	m.GCSys = Gauge(memStats.GCSys)
	m.HeapAlloc = Gauge(memStats.HeapAlloc)
	m.HeapIdle = Gauge(memStats.HeapIdle)
	m.HeapInuse = Gauge(memStats.HeapInuse)
	m.HeapObjects = Gauge(memStats.HeapObjects)
	m.HeapReleased = Gauge(memStats.HeapReleased)
	m.HeapSys = Gauge(memStats.HeapSys)
	m.LastGC = Gauge(memStats.LastGC)
	m.Lookups = Gauge(memStats.Lookups)
	m.MCacheInuse = Gauge(memStats.MCacheInuse)
	m.MCacheSys = Gauge(memStats.MCacheSys)
	m.MSpanInuse = Gauge(memStats.MSpanInuse)
	m.MSpanSys = Gauge(memStats.MSpanSys)
	m.Mallocs = Gauge(memStats.Mallocs)
	m.NextGC = Gauge(memStats.NextGC)
	m.NumForcedGC = Gauge(memStats.NumForcedGC)
	m.NumGC = Gauge(memStats.NumGC)
	m.OtherSys = Gauge(memStats.OtherSys)
	m.PauseTotalNs = Gauge(memStats.PauseTotalNs)
	m.StackInuse = Gauge(memStats.StackInuse)
	m.StackSys = Gauge(memStats.StackSys)
	m.Sys = Gauge(memStats.Sys)
	m.TotalAlloc = Gauge(memStats.TotalAlloc)

	atomic.AddInt64((*int64)(&m.PollCount), 1)
	m.RandomValue = Gauge(rand.Float64())
}
