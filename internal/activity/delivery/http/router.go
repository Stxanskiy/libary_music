package http

import (
	"github.com/gorilla/mux"

	"net/http"
)

// InitRouter инициализирует маршруты с использованием `mux`.
func InitRouter(songHandler *handler.) *mux.Router {
	r := mux.NewRouter()

	// Основные маршруты для CRUD операций.
	r.HandleFunc("/song", songHandler.AddSong).Methods(http.MethodPost)
	r.HandleFunc("/song/{id:[0-9]+}", songHandler.GetSongByID).Methods(http.MethodGet)
	r.HandleFunc("/song/{id:[0-9]+}", songHandler.UpdateSong).Methods(http.MethodPut)
	r.HandleFunc("/song/{id:[0-9]+}", songHandler.DeleteSong).Methods(http.MethodDelete)

	// Маршрут для получения списка песен с поддержкой пагинации.
	r.HandleFunc("/songs", songHandler.ListSongsWithPagination).Methods(http.MethodGet)

	return r
}
