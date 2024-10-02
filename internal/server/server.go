package server

import (
	// Импорт модуля конфигурации
	"fmt"
	"libary_music/config"
	"net/http"
)

// Server представляет структуру HTTP-сервера.
type Server struct {
	cfg        *config.Config
	httpServer *http.Server
}

// New создает новый сервер с использованием конфигурации.
func New(cfg *config.Config) (*Server, error) {
	// Создаем новый HTTP-сервер.
	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port), // Указываем адрес и порт из конфигурации.
	}

	return &Server{httpServer: srv}, nil
}

// Run запускает HTTP-сервер.
func (s *Server) Run() error {
	fmt.Printf("Starting server on %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
