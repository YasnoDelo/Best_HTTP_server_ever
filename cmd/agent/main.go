package main

import (
	"time"

	"github.com/YasnoDelo/Best_HTTP_server_ever/internal/agent"
)

func main() {

	// Настраиваем агента для сбора и отправки метрик
	baseURL := "http://localhost:8080"
	pollInterval := 2 * time.Second
	reportInterval := 10 * time.Second

	a := agent.NewAgent(baseURL, pollInterval, reportInterval)
	a.Run()
}
