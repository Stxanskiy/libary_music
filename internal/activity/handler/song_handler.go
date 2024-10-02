package handler

import (
	"context"
	"encoding/json"
	"fmt"
	_interface "libary_music/internal/activity/interface"
	"libary_music/internal/activity/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// SongHandler представляет собой обработчик для песен.
type SongHandler struct {
	SongRepo _interface.SongRepo
}

// NewSongHandler создает новый экземпляр SongHandler.
func NewSongHandler(songRepo _interface.SongRepo) *SongHandler {
	return &SongHandler{SongRepo: songRepo}
}

// AddSong добавляет новую песню.
func (h *SongHandler) AddSong(w http.ResponseWriter, r *http.Request) {
	var newSong model.Song
	if err := json.NewDecoder(r.Body).Decode(&newSong); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	id, err := h.SongRepo.AddSong(context.Background(), &newSong)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to add song: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Song added with ID: %d", id)
}

// GetSongByID получает песню по ID.
func (h *SongHandler) GetSongByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	song, err := h.SongRepo.GetSongByID(context.Background(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get song: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(song)
}

// DeleteSong удаляет песню по ID.
func (h *SongHandler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, err := h.SongRepo.DeleteSong(context.Background(), id); err != nil {
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

	if _, err := h.SongRepo.UpdateSong(context.Background(), &song); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update song: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Song updated successfully")
}

// ListSongsWithPagination возвращает список песен с поддержкой пагинации.
func (h *SongHandler) ListSongsWithPagination(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры limit и offset из URL.
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10 // Устанавливаем значение по умолчанию, если параметр отсутствует.
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	songs, err := h.SongRepo.ListSongsWithPagination(context.Background(), limit, offset)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get songs: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(songs)
}
