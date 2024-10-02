package main

import (
	"fmt"
	"gitlab.com/nevasik7/lg"
	"libary_music/config"
	"libary_music/internal/server"
	"net/http"
)

// @title Music Library API
// @version 1.0.0
// @description API для управления музыкальной библиотекой.
// @BasePath /
func main() {
	// Загружаем конфигурацию из .env файла.
	cfg := config.MustLoad()

	//проверка сервера на работоспослбность
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Music Library API is running!")
	})

	// Инициализируем новый сервер с конфигурацией.
	srv, err := server.New(cfg)
	if err != nil {
		lg.Panicf("Ошибка инициализации сервера: %v", err)
	}

	// Запуск сервера.
	if err := srv.Run(); err != nil {
		lg.Fatalf("Сервер не смог запуститься: %v", err)
	}
}
