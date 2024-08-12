package main

import (
	"log"
	"net/http"

	"github.com/YasnoDelo/Best_HTTP_server_ever/internal/server"
)

func main() {
	// Инициализация нового сервера
	srv := server.NewServer()

	// Определяем порт, на котором будет работать сервер
	addr := ":8080"

	// Логирование старта сервера
	log.Printf("Starting server on %s", addr)

	// Запуск сервера
	if err := http.ListenAndServe(addr, srv); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
