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

// AddSong добавляет новую песню.
func (h *SongHandler) AddSong(w http.ResponseWriter, r *http.Request) {
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
}

// GetSongByID получает песню по ID.
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

// DeleteSong удаляет песню по ID.
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

// UpdateSong обновляет данные песни.
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
