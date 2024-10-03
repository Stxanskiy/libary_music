package http

import (
	"fmt"
	"libary_music/pkg/storage"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *storage.DB {
	// Настраиваем тестовую базу данных.
	dsn := os.Getenv("DBURL")
	db, err := storage.NewDB(dsn)
	if err != nil {
		t.Fatalf("Не удалось подключиться к тестовой базе данных: %v", err)
	}

	return db
}

func TestAddSong(t *testing.T) {
	// Подготавливаем тестовую базу данных.
	db := setupTestDB(t)
	defer db.Close()

	// Инициализируем роутер.
	r := RouterInit(db)

	// Создаем тестовый сервер.
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Делаем запрос на добавление новой песни.
	songData := `{"music_band_id": 1, "title": "Test Song", "release_date": "2023-10-01", "lyrics": "Test lyrics", "link": "https://example.com/song"}`

	resp, err := http.Post(fmt.Sprintf("%s/songs", ts.URL), "application/json", strings.NewReader(songData))
	if err != nil {
		t.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа.
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetSongByID(t *testing.T) {
	// Подготавливаем тестовую базу данных.
	db := setupTestDB(t)
	defer db.Close()

	// Инициализируем роутер.
	r := RouterInit(db)

	// Создаем тестовый сервер.
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Добавляем тестовую песню.
	songData := `{"music_band_id": 1, "title": "Test Song", "release_date": "2023-10-01", "lyrics": "Test lyrics", "link": "https://example.com/song"}`
	_, err := http.Post(fmt.Sprintf("%s/songs", ts.URL), "application/json", strings.NewReader(songData))
	if err != nil {
		t.Fatalf("Ошибка при выполнении запроса: %v", err)
	}

	// Выполняем запрос на получение песни по ID.
	resp, err := http.Get(fmt.Sprintf("%s/songs/1", ts.URL))
	if err != nil {
		t.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа.
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
