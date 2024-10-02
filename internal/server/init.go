package server

import (
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) MapHandlers() error {
	var (
		r *chi.Mux
	)

	r = chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	}))

	/*	authRouter := r.Route("/api", func(r chi.Router) {})

		// routes
		s.httpServer.Handler = r

		http.InitActivity(authRouter, activityHand, mw)
		http.InitStatistics(authRouter, statisticsHand, mw)
		http.InitAuth(authRouter, authHand, mw)
	*/
	return nil
}
