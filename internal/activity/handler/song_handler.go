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

// AddSong добавляет новый куплет в песню.
// @Summary Добавляение песни
// @Description Добавляет новую песню в музыкальную библиотеку
// @Tags Song
// @Accept json
// @Produce json
// @Param verse body model.SongS true "Данные песни"
// @Success 201 {object} model.SongResponse
// @Failure 400 {object} model.ErrorResponse
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

// GetSongByID Получение песни по {id}
// @Summary Получение песни
// @Description Получает песню по ее {id}
// @Tags Song
// @Accept json
// @Produce json
// @Param song_id path int true "ID песни"
// @Success 200 {object} model.Song
// @Failure 500 {object} model.ErrorResponse
// @Router /song/{song_id} [get]
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

// DeleteSong Удаляет песню по ее {id}
// @Summary Удаляет песню по ее {id}
// @Description Удаление песни из музыкальной библиотеки по ее {id}
// @Tags Song
// @Accept json
// @Produce json
// @Param verse body model.Song true "Данные куплета"
// @Success 201 {object} model.SongResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /verse [delete]
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

// UpdateSong обновляет песню по его ID.
// @Summary Обновляет песню по ID
// @Description Обновляет песню из музыкальной библиотеки по его {id}
// @Tags Song
// @Accept json
// @Produce json
// @Param id path int true "ID Песни"
// @Param verse body model.Song true "Данные Песни"
// @Success 204 {object} model.SongResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /song/{id} [put]
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

// ListSongsWithPagination возвращает список песен с поддержкой пагинации.
// @Summary Получение списка песен
// @Description Возвращает список песен с пагинацией
// @Tags Song
// @Accept json
// @Produce json
// @Param limit query int false "Количество песен на странице"-(limit)
// @Param offset query int false "Смещение"-(offset)
// @Success 200 {array} model.SongResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /song [get]
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
