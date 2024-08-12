package agent

import (
	"fmt"
	"time"

	"github.com/YasnoDelo/Best_HTTP_server_ever/internal/httpclient"
	"github.com/YasnoDelo/Best_HTTP_server_ever/internal/metrics"
)

type Agent struct {
	metrics        *metrics.Metrics
	httpClient     *httpclient.HTTPClient
	pollInterval   time.Duration
	reportInterval time.Duration
}

func (a *Agent) SetHTTPClient(client *httpclient.HTTPClient) {
	a.httpClient = client
}

func NewAgent(baseURL string, pollInterval, reportInterval time.Duration) *Agent {
	return &Agent{
		metrics:        metrics.NewMetrics(),
		httpClient:     httpclient.NewHTTPClient(baseURL),
		pollInterval:   pollInterval,
		reportInterval: reportInterval,
	}
}

func (agent *Agent) Run() {
	tickerPoll := time.NewTicker(agent.pollInterval)
	tickerReport := time.NewTicker(agent.reportInterval)

	for {
		select {
		case <-tickerPoll.C:
			fmt.Println("Polling metrics...")
			agent.metrics.UpdateMetrics()
		case <-tickerReport.C:
			fmt.Println("Reporting metrics...")
			agent.reportMetrics()
		}
	}
}

func (agent *Agent) reportMetrics() {
	fmt.Println("Starting to report metrics...")
	// Send all metrics to the server
	for name, metric := range map[string]interface{}{
		"Alloc":         agent.metrics.Alloc,
		"BuckHashSys":   agent.metrics.BuckHashSys,
		"Frees":         agent.metrics.Frees,
		"GCCPUFraction": agent.metrics.GCCPUFraction,
		"GCSys":         agent.metrics.GCSys,
		"HeapAlloc":     agent.metrics.HeapAlloc,
		"HeapIdle":      agent.metrics.HeapIdle,
		"HeapInuse":     agent.metrics.HeapInuse,
		"HeapObjects":   agent.metrics.HeapObjects,
		"HeapReleased":  agent.metrics.HeapReleased,
		"HeapSys":       agent.metrics.HeapSys,
		"LastGC":        agent.metrics.LastGC,
		"Lookups":       agent.metrics.Lookups,
		"MCacheInuse":   agent.metrics.MCacheInuse,
		"MCacheSys":     agent.metrics.MCacheSys,
		"MSpanInuse":    agent.metrics.MSpanInuse,
		"MSpanSys":      agent.metrics.MSpanSys,
		"Mallocs":       agent.metrics.Mallocs,
		"NextGC":        agent.metrics.NextGC,
		"NumForcedGC":   agent.metrics.NumForcedGC,
		"NumGC":         agent.metrics.NumGC,
		"OtherSys":      agent.metrics.OtherSys,
		"PauseTotalNs":  agent.metrics.PauseTotalNs,
		"StackInuse":    agent.metrics.StackInuse,
		"StackSys":      agent.metrics.StackSys,
		"Sys":           agent.metrics.Sys,
		"TotalAlloc":    agent.metrics.TotalAlloc,
		"PollCount":     agent.metrics.PollCount,
		"RandomValue":   agent.metrics.RandomValue,
	} {
		var metricType string
		var metricValue interface{}

		switch v := metric.(type) {
		case metrics.Counter:
			metricType = "counter"
			metricValue = int64(v)
		case metrics.Gauge:
			metricType = "gauge"
			metricValue = float64(v)
		default:
			fmt.Printf("Unknown metric type for %s\n", name)
			continue
		}

		fmt.Printf("Sending metric: %s (%s) with value: %v\n", name, metricType, metricValue)
		if err := agent.httpClient.SendMetric(metricType, name, metricValue); err != nil {
			fmt.Printf("Error sending metric: %v\n", err)
		} else {
			fmt.Printf("Metric %s sent successfully\n", name)
		}
	}
	fmt.Println("Finished reporting metrics.")
}
