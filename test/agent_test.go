package agent_test

import (
	"testing"
	"time"

	"github.com/YasnoDelo/Best_HTTP_server_ever/agent"
)

type MockHTTPClient struct {
	sentMetrics []string
}

func (m *MockHTTPClient) SendMetric(metricType, metricName string, value interface{}) error {
	m.sentMetrics = append(m.sentMetrics, metricName)
	return nil
}

func TestAgentRun(t *testing.T) {
	// Создаем mock HTTP клиент
	mockClient := &MockHTTPClient{}

	// Создаем агента с mock клиентом и небольшими интервалами
	a := agent.NewAgent("", 1*time.Millisecond, 2*time.Millisecond)
	a.SetHTTPClient(mockClient)

	// Запускаем агента в отдельной горутине
	go a.Run()

	// Ждем, чтобы агент успел выполнить хотя бы один цикл
	time.Sleep(10 * time.Millisecond)

	// Проверяем, что метрики были отправлены
	if len(mockClient.sentMetrics) == 0 {
		t.Error("Expected some metrics to be sent, but got none")
	}
}
