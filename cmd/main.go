package main

import (
	"gitlab.com/nevasik7/lg"
	"libary_music/config"
	"libary_music/internal/server"
)

// @title Онлайн Библиотека Песен
// @version 1.0
// @description Это API для управления музыкальной библиотекой.

// @BasePath /
// @schemes http

func main() {
	lg.Init()
	cfg := config.MustLoad()

	srv, err := server.New(cfg)
	if err != nil {
		lg.Panicf("Ошибка инициализации сервера: %v", err)
	}

	if err = srv.Run(); err != nil {
		lg.Fatalf("Сервер не смог запуститься: %v", err)
	}
}
