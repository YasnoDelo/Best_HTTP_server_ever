package agent_test

import (
	"testing"
	"time"

	"github.com/YasnoDelo/Best_HTTP_server_ever/internal/agent"
	"github.com/YasnoDelo/Best_HTTP_server_ever/internal/httpclient"
)

// Мок-объект, реализующий интерфейс HTTPClientInterface
type MockHTTPClient struct {
	sentMetrics []string
}

func (m *MockHTTPClient) SendMetric(metricType, metricName string, value interface{}) error {
	m.sentMetrics = append(m.sentMetrics, metricName)
	return nil
}

// Создаем обёртку для мок-объекта
type MockHTTPClientWrapper struct {
	*httpclient.HTTPClient
	mockClient *MockHTTPClient
}

func (w *MockHTTPClientWrapper) SendMetric(metricType, metricName string, value interface{}) error {
	return w.mockClient.SendMetric(metricType, metricName, value)
}

func TestAgentRun(t *testing.T) {
	// Создаем mock HTTP клиент
	mockClient := &MockHTTPClient{}
	clientWrapper := &MockHTTPClientWrapper{mockClient: mockClient}

	// Создаем агента с mock клиентом и небольшими интервалами
	testAgent := agent.NewAgent("", 1*time.Millisecond, 2*time.Millisecond)
	testAgent.SetHTTPClient(clientWrapper.HTTPClient)

	// Запускаем агента в отдельной горутине
	go testAgent.Run()

	// Ждем, чтобы агент успел выполнить хотя бы один цикл
	time.Sleep(10 * time.Millisecond)

	// Проверяем, что метрики были отправлены
	if len(mockClient.sentMetrics) == 0 {
		t.Error("Expected some metrics to be sent, but got none")
	}
}
