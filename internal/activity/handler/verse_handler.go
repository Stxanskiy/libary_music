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

// @Summary Добавляет новый куплет в песню
// @Description AddVerse обрабатывает POST-запрос для добавления нового куплета.
// @Tags Куплеты
// @Accept json
// @Produce json
// @Param verse body Verse true  "Даные куалета"
// @Success 201 {object} VerseResponse
// @Failure 400 {object} ErrorResponse
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

// @Summary Обновляет данные песни
// @Description UpdateVerse позволяет обновить куплет по ее {id}
// @Tags Куплеты
// @Accept json
// @Produce json
// @Param verse body Verse true  "Даные куалета"
// @Success 201 {object} VerseResponse
// @Failure 400 {object} ErrorResponse
// @Router /verse/{id} [PUT]

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

// @Summary Получение куплетов песни с пагинацией
// @Description GetSongVerse позволяет получить куплеты песни по {id} пенси
// @Tags Куплеты
// @Accept json
// @Produce json
// @Param verse {id}
// @Success 201 {object} VerseResponse
// @Failure 400 {object} ErrorResponse
// @Router /verse/{id} [GET]

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
