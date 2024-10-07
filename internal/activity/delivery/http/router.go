package http

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"libary_music/docs"
	"libary_music/internal/activity/handler"
	"libary_music/internal/activity/repo"
	"libary_music/internal/activity/uc"
	"libary_music/pkg/api/musiclibary"
	"libary_music/pkg/storage"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"gitlab.com/nevasik7/lg"
)

// RouterInit инициализирует маршруты с использованием `chi`.
func RouterInit(db *storage.DB) *chi.Mux {

	// Создаем новый роутер с `chi`.
	r := chi.NewRouter()

	// Инициализация репозиториев.
	songRepo := repo.NewSongRepo(db)
	verseRepo := repo.NewVerseRepo(db)
	if songRepo == nil || verseRepo == nil {
		fmt.Println("Ошибка: репозитории не были созданы")
		return nil
	}

	musicClient := musiclibary.NewMusicLibraryClient(os.Getenv("URL_API_NEW_MUSIC_LIBARY_CLIENT"))

	// Иници3ализация Use Case.
	songUC := uc.NewSongUC(songRepo, musicClient)
	verseUC := uc.NewVerseUC(verseRepo)

	// Инициализация хендлеров.
	songHandler := handler.NewSongHandler(*songUC)
	verseHandler := handler.NewVerseHandler(*verseUC)
	if songHandler == nil || verseHandler == nil {
		fmt.Println("Ошибка: обработчики не были созданы")
		return nil
	}

	//swagger
	docs.SwaggerInfo.BasePath = "http://localhost:8080/swagger/doc.json"

	lg.Trace("Обработчики успешно инициализированы")

	// Определяем маршруты для песен.
	r.Post("/song", songHandler.AddSong)
	r.Get("/song/{id}", songHandler.GetSongByID)
	r.Put("/song/{id}", songHandler.UpdateSong)
	r.Delete("/song/{id}", songHandler.DeleteSong)

	// Маршрут для получения списка песен с поддержкой пагинации.
	r.Get("/song", songHandler.ListSongsWithPagination)

	// Определяем маршруты для куплетов.
	r.Post("/verse", verseHandler.AddVerse)
	r.Put("/verse/{id}", verseHandler.UpdateVerse)
	r.Get("/verse/{id}", verseHandler.GetSongVerse)

	// Проверка работоспособности сервера.
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Music Library API is running!"))
	})

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
