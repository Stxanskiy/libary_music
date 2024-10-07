package main

import (
	"gitlab.com/nevasik7/lg"
	"libary_music/config"
	"libary_music/internal/server"
)

// @title Music Library API
// @version 1.0.0
// @description API для управления музыкальной библиотекой.
// @BasePath /
func main() {
	cfg := config.MustLoad()

	lg.Init()

	srv, err := server.New(cfg)
	if err != nil {
		lg.Panicf("Ошибка инициализации сервера: %v", err)
	}

	if err = srv.Run(); err != nil {
		lg.Fatalf("Сервер не смог запуститься: %v", err)
	}
}
