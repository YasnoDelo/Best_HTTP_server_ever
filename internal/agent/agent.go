package agent

import (
	"fmt"
	"time"

	"github.com/YasnoDelo/Best_HTTP_server_ever/httpclient"
	"github.com/YasnoDelo/Best_HTTP_server_ever/metrics"
)

type Agent struct {
	metrics        *metrics.Metrics
	httpClient     *httpclient.HTTPClient
	pollInterval   time.Duration
	reportInterval time.Duration
}

func NewAgent(baseURL string, pollInterval, reportInterval time.Duration) *Agent {
	return &Agent{
		metrics:        metrics.NewMetrics(),
		httpClient:     httpclient.NewHTTPClient(baseURL),
		pollInterval:   pollInterval,
		reportInterval: reportInterval,
	}
}

func (a *Agent) Run() {
	tickerPoll := time.NewTicker(a.pollInterval)
	tickerReport := time.NewTicker(a.reportInterval)

	for {
		select {
		case <-tickerPoll.C:
			a.metrics.UpdateMetrics()
		case <-tickerReport.C:
			a.reportMetrics()
		}
	}
}

func (a *Agent) reportMetrics() {
	// Send all metrics to the server
	for _, metric := range map[string]interface{}{
		"Alloc":         a.metrics.Alloc,
		"BuckHashSys":   a.metrics.BuckHashSys,
		"Frees":         a.metrics.Frees,
		"GCCPUFraction": a.metrics.GCCPUFraction,
		"GCSys":         a.metrics.GCSys,
		"HeapAlloc":     a.metrics.HeapAlloc,
		"HeapIdle":      a.metrics.HeapIdle,
		"HeapInuse":     a.metrics.HeapInuse,
		"HeapObjects":   a.metrics.HeapObjects,
		"HeapReleased":  a.metrics.HeapReleased,
		"HeapSys":       a.metrics.HeapSys,
		"LastGC":        a.metrics.LastGC,
		"Lookups":       a.metrics.Lookups,
		"MCacheInuse":   a.metrics.MCacheInuse,
		"MCacheSys":     a.metrics.MCacheSys,
		"MSpanInuse":    a.metrics.MSpanInuse,
		"MSpanSys":      a.metrics.MSpanSys,
		"Mallocs":       a.metrics.Mallocs,
		"NextGC":        a.metrics.NextGC,
		"NumForcedGC":   a.metrics.NumForcedGC,
		"NumGC":         a.metrics.NumGC,
		"OtherSys":      a.metrics.OtherSys,
		"PauseTotalNs":  a.metrics.PauseTotalNs,
		"StackInuse":    a.metrics.StackInuse,
		"StackSys":      a.metrics.StackSys,
		"Sys":           a.metrics.Sys,
		"TotalAlloc":    a.metrics.TotalAlloc,
		"PollCount":     a.metrics.PollCount,
		"RandomValue":   a.metrics.RandomValue,
	} {
		metricType := "gauge"
		if _, ok := metric.(metrics.Counter); ok {
			metricType = "counter"
		}

		if err := a.httpClient.SendMetric(metricType, metric.(metrics.Gauge)); err != nil {
			fmt.Printf("Error sending metric: %v\n", err)
		}
	}
}
