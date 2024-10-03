package handler

import (
	"encoding/json"
	"libary_music/internal/activity/model"
	"libary_music/internal/activity/uc"
	"net/http"
)

type VerseHandler struct {
	UC uc.VerseUC
}

// NewVerseHandler создает новый экземпляр VerseHandler.
func NewVerseHandler(uc uc.VerseUC) *VerseHandler {
	return &VerseHandler{UC: uc}
}

// AddVerse обрабатывает POST-запрос для добавления нового куплета.
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

// ОБновление куплетов
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
