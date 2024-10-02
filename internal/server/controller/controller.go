package controller

/*// VerseController отвечает за обработку запросов к API.
type VerseController struct {
	service service.VerseService
}

// NewVerseController создаёт новый контроллер для куплетов.
func NewVerseController(svc service.VerseService) *VerseController {
	return &VerseController{service: svc}
}

// AddVerse добавляет новый куплет.
func (vc *VerseController) AddVerse(w http.ResponseWriter, r *http.Request) {
	var verse model.Verse
	if err := json.NewDecoder(r.Body).Decode(&verse); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	ctx := context.TODO()
	if err := vc.service.AddVerse(ctx, &verse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(verse)
}

// GetVersesBySongID возвращает список куплетов по песне.
func (vc *VerseController) GetVersesBySongID(w http.ResponseWriter, r *http.Request) {
	var mux chi.Mux
	vars := mux.Vars(r)
	songID, err := strconv.ParseInt(vars["song_id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	ctx := context.TODO()
	verses, err := vc.service.GetVersesBySongID(ctx, songID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(verses)
}

// ListVersesWithPagination возвращает список куплетов с пагинацией.
func (vc *VerseController) ListVersesWithPagination(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))

	ctx := context.TODO()
	verses, err := vc.service.ListVersesWithPagination(ctx, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(verses)
}

// UpdateVerse обновляет куплет.
func (vc *VerseController) UpdateVerse(w http.ResponseWriter, r *http.Request) {
	var verse model.Verse
	if err := json.NewDecoder(r.Body).Decode(&verse); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	ctx := context.TODO()
	if err := vc.service.UpdateVerse(ctx, &verse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(verse)
}

// DeleteVerse удаляет куплет.
func (vc *VerseController) DeleteVerse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	verseID, err := strconv.ParseInt(vars["verse_id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid verse ID", http.StatusBadRequest)
		return
	}

	ctx := context.TODO()
	if err := vc.service.DeleteVerse(ctx, verseID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
*/
