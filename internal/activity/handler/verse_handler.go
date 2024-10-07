package handler

import (
	"encoding/json"
	"libary_music/internal/activity/model"
	"libary_music/internal/activity/uc"
	"net/http"
	"strconv"
	"strings"
)

type VerseHandler struct {
	UC uc.VerseUC
}

func NewVerseHandler(uc uc.VerseUC) *VerseHandler {
	return &VerseHandler{UC: uc}
}

// AddVerse добавляет новый куплет в песню.
// @Summary Добавляет новый куплет в песню
// @Description Добавляет новый куплет в песню
// @Tags Verses
// @Accept json
// @Produce json
// @Param verse body model.Verse true "Данные куплета"
// @Success 201 {object} model.Verse
// @Failure 400 {object} model.ErrorResponse
// @Router /verse [post]
func (h *VerseHandler) AddVerse(w http.ResponseWriter, r *http.Request) {
	var verse model.Verse
	if err := json.NewDecoder(r.Body).Decode(&verse); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if _, err := h.UC.AddVerse(r.Context(), &verse); err != nil {
		http.Error(w, "Failed to add verse", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateVerse обновляет куплет по его ID.
// @Summary Обновляет куплет по ID
// @Description Обновляет куплет по его ID
// @Tags Verses
// @Accept json
// @Produce json
// @Param id path int true "ID куплета"
// @Param verse body model.Verse true "Данные куплета"
// @Success 204 {object} model.Verse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /verse/{id} [put]
func (h *VerseHandler) UpdateVerse(w http.ResponseWriter, r *http.Request) {
	var verse model.Verse
	if err := json.NewDecoder(r.Body).Decode(&verse); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if _, err := h.UC.UpdateVerse(r.Context(), &verse); err != nil {
		http.Error(w, "Failed to update verse", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetSongVerse получает куплеты по ID песни.
// @Summary Получение куплетов по ID песни
// @Description Получает куплеты по ID песни
// @Tags Verses
// @Accept json
// @Produce json
// @Param song_id path int true "ID песни"
// @Success 200 {object} model.Verse
// @Failure 404 {object} model.ErrorResponse
// @Router /song/{song_id}/verses [get]
func (h *VerseHandler) GetSongVerse(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	songIDStr := strings.Split(url, "/")

	songID, err := strconv.Atoi(songIDStr[2])
	if err != nil {
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10 // Значение по умолчанию
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0 // Значение по умолчанию
	}

	verses, err := h.UC.GetSongVerse(r.Context(), songID, limit, offset)
	if err != nil {
		http.Error(w, "Failed to get verses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(verses)
}
