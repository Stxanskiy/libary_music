package server

import (
	"fmt"
	"gitlab.com/nevasik7/lg"
	"libary_music/config"
	http2 "libary_music/internal/activity/delivery/http"
	"libary_music/pkg/storage"
	"net/http"
	"os"
)

// Server представляет структуру HTTP-сервера.
type Server struct {
	cfg        *config.Config
	httpServer *http.Server
}

// New создает новый сервер с использованием конфигурации.
func New(cfg *config.Config) (*Server, error) {
	lg.Init()
	// Подключение к базе данных.
	dbURL := os.Getenv("DBURL")
	if dbURL == "" {
		lg.Panic("Ошибка: переменная окружения URL не задана")
		return nil, fmt.Errorf("DBURL not set in environment variables")
	}

	lg.Tracef("Подключение к базе данных по адресу: %s\n", dbURL)
	db, err := storage.NewDB(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("Ошибка подключения к БД: %w", err)
	}
	lg.Trace("Подключение к базе данных успешно")

	// Инициализация маршрутизатора.
	r := http2.RouterInit(db)
	if r == nil {
		fmt.Println("Ошибка инициализации маршрутизатора")
		return nil, fmt.Errorf("инициализация маршрутизатора завершилась неудачей")
	}

	// Создание HTTP-сервера.
	srv := &http.Server{
		Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
		Handler: r,
	}

	lg.Tracef("Сервер инициализирован на %s\n", srv.Addr)
	return &Server{
		cfg:        cfg,
		httpServer: srv,
	}, nil
}

// Run запускает HTTP-сервер.
func (s *Server) Run() error {
	lg.Infof("Запуск сервера на %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
