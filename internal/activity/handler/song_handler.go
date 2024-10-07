package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"libary_music/internal/activity/model"
	"libary_music/internal/activity/uc"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// SongHandler представляет собой обработчик для песен.
type SongHandler struct {
	UC uc.SongUC
}

// NewSongHandler создает новый экземпляр SongHandler.
func NewSongHandler(uc uc.SongUC) *SongHandler {
	return &SongHandler{UC: uc}
}

// @Summary Добавление песни
// @Description Добавляет песню в бд и возвращает ее ID
// @Tags Песни
// @Accept json
// @Produce json
// @Param song body Song true  "Данные песни"
// @Success 201 {object} SongResponse
// @Failure 400 {object} ErrorRespnse
// @Router /song [post]
func (h *SongHandler) AddSong(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Group string `json:"group"`
		Song  string `json:"song"`
	}

	// Парсим входящий запрос.
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	// Создаем песню на основе входящих данных.
	newSong := &model.Song{
		GroupName: request.Group,
		Title:     request.Song,
	}

	// Вызываем use case для добавления новой песни.
	id, err := h.UC.AddSong(context.Background(), newSong)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка добавления песни: %v", err), http.StatusInternalServerError)
		return
	}

	// Формируем успешный ответ.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Песня успешно добавлена",
		"song_id": id,
	})
}

/*func (h *SongHandler) AddSong(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Group string `json:"group"`
		Song  string `json:"song"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Формаь запроса неверный", http.StatusBadRequest)
		return
	}

}*/

/*func (h *SongHandler) AddSong(w http.ResponseWriter, r *http.Request) {
	var newSong model.Song
	if err := json.NewDecoder(r.Body).Decode(&newSong); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	id, err := h.UC.AddSong(context.Background(), &newSong)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to add song: %v", err), http.StatusInternalServerError)
		return
	}

	// Формируем успешный ответ.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Song created successfully",
		"song_id": id,
	})
}*/

// @Summary Получение песни
// @Description GetSongByID получает песню по ID.
// @Tags Песни
// @Accept json
// @Produce json
// @Param song id "ID- Песни"
// @Success 201 {id} SongResponse
// @Failure 400 {id} ErrorResponse
// @Router /song/{id} [GET]

func (h *SongHandler) GetSongByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	song, err := h.UC.GetSongByID(context.Background(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get song: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(song)
}

// @Summary Удаление песни
// @Description DeleteSong удаляет песню по ID.
// @Tags Песни
// @Accept json
// @Produce json
// @Param song id "Удаление песни"
// @Success 201 {id} SongResponse
// @Failure 400 {id} ErrorResponse
// @Router /song/{id} [delete]

func (h *SongHandler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, err := h.UC.DeleteSong(context.Background(), id); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete song: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary Обновление песни
// @Description UpdateSong обновляет данные песни по ее id.
// @Tags Песни
// @Accept json
// @Produce json
// @Param song body Song true "Данные песни"
// @Success 201 {object} SongResponse
// @Failure 401 {object} ErrorResponse
// @Router /song/{id} [PUT]

func (h *SongHandler) UpdateSong(w http.ResponseWriter, r *http.Request) {
	var song model.Song
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if _, err := h.UC.UpdateSong(context.Background(), &song); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update song: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Song updated successfully")
}

// @Summary Получение песен
// @Description ListSongsWithPagination возвращает список песен с поддержкой пагинации и Фильтрации.
// @Tags Песни
// @Accept json
// @Produce json
// @Param "Данные песен"
// @Success 201 {objects} SongResponse
// @Failure 400 {objects} ErrorResponse
// @Router /song [GET]

func (h *SongHandler) ListSongsWithPagination(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10 // Значение по умолчанию.
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	songs, err := h.UC.ListSongsWithPagination(context.Background(), limit, offset)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get songs: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}
