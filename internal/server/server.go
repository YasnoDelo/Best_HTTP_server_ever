package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type Metric struct {
	Type  string
	Value interface{}
}

type Server struct {
	metrics map[string]Metric
	mutex   sync.Mutex
}

// NewServer создает новый экземпляр сервера
func NewServer() *Server {
	return &Server{
		metrics: make(map[string]Metric),
	}
}

// ServeHTTP позволяет серверу реализовать интерфейс http.Handler
func (server *Server) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	switch {
	case request.URL.Path == "/":
		server.serveImage(responseWriter, request)
	case request.URL.Path == "/hi":
		server.handleWelcome(responseWriter, request)
	case strings.HasPrefix(request.URL.Path, "/update/"):
		server.handleMetricUpdate(responseWriter, request)
	default:
		http.Error(responseWriter, "404 Not Found", http.StatusNotFound)
	}
}

// handleWelcome обрабатывает запросы к корневому маршруту
func (server *Server) handleWelcome(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
	fmt.Fprintf(responseWriter, "Welcome to the Best HTTP Server Ever!")
}

// serveImage обрабатывает запрос на получение изображения
func (s *Server) serveImage(responseWriter http.ResponseWriter, request *http.Request) {
	// Получаем текущую рабочую директорию
	workingDir, err := os.Getwd()
	if err != nil {
		http.Error(responseWriter, "Failed to get working directory", http.StatusInternalServerError)
		return
	}

	// Удаляем последнюю директорию "cmd" для получения корневой директории проекта
	workingDir = filepath.Dir(workingDir)
	// fmt.Print(workingDir)

	// Формируем путь к изображению относительно корневой директории
	imagePath := filepath.Join(workingDir, "img", "Firefox_wallpaper.png")
	// fmt.Print(imagePath)

	// Проверяем существование файла
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		http.Error(responseWriter, "Image not found", http.StatusNotFound)
		return
	}

	// Открываем изображение
	imageFile, err := os.Open(imagePath)
	if err != nil {
		http.Error(responseWriter, "Failed to open image", http.StatusInternalServerError)
		return
	}
	defer imageFile.Close()

	// Устанавливаем заголовок Content-Type
	responseWriter.Header().Set("Content-Type", "image/png")

	// Отправляем изображение в ответе
	_, err = io.Copy(responseWriter, imageFile)
	if err != nil {
		http.Error(responseWriter, "Failed to serve image", http.StatusInternalServerError)
	}
}

// handleMetricUpdate обрабатывает запросы на обновление метрик
func (server *Server) handleMetricUpdate(responseWriter http.ResponseWriter, request *http.Request) {
	pathParts := strings.Split(strings.Trim(request.URL.Path, "/"), "/")

	if len(pathParts) != 4 || pathParts[0] != "update" {
		http.Error(responseWriter, "Invalid URL format", http.StatusBadRequest)
		return
	}

	metricType := pathParts[1]
	metricName := pathParts[2]
	metricValueStr := pathParts[3]

	server.mutex.Lock()
	defer server.mutex.Unlock()

	var metricValue interface{}
	var parseError error

	switch metricType {
	case "counter":
		metricValue, parseError = strconv.ParseInt(metricValueStr, 10, 64)
	case "gauge":
		metricValue, parseError = strconv.ParseFloat(metricValueStr, 64)
	default:
		http.Error(responseWriter, "Invalid metric type", http.StatusBadRequest)
		return
	}

	if parseError != nil {
		http.Error(responseWriter, "Invalid metric value", http.StatusBadRequest)
		return
	}

	server.metrics[metricName] = Metric{
		Type:  metricType,
		Value: metricValue,
	}

	fmt.Printf("Received metric: %s (%s) with value: %v\n", metricName, metricType, metricValue)
	responseWriter.WriteHeader(http.StatusOK)
}

// GetMetrics возвращает копию всех метрик
func (server *Server) GetMetrics() map[string]Metric {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	metricsCopy := make(map[string]Metric, len(server.metrics))
	for key, value := range server.metrics {
		metricsCopy[key] = value
	}
	return metricsCopy
}
