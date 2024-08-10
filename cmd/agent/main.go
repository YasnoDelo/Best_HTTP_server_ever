package main

import (
	"time"

	"github.com/YasnoDelo/Best_HTTP_server_ever/agent"
)

func main() {
	baseURL := "http://localhost:8080"
	pollInterval := 2 * time.Second
	reportInterval := 10 * time.Second

	a := agent.NewAgent(baseURL, pollInterval, reportInterval)
	a.Run()
}
