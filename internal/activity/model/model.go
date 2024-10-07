package model

type MusicBand struct {
	MusicBandID int    `json:"music_band_id"`
	Name        string `json:"name"`
}

// Song определяет структуру песни.
type Song struct {
	SongID      int    `json:"song_id"`
	MusicBandID int    `json:"music_band_id"`
	GroupName   string `json:"group_name"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Lyrics      string `json:"lyrics"`
	Link        string `json:"link"`
}

/*type Song struct {
	SongID      int       `json:"song_id"`
	MusicBandID int       `json:"music_band_id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	Lyrics      string    `json:"lyrics"`
	Link        string    `json:"link"`
}*/

type Verse struct {
	VerseID  int    `json:"verse_id"`
	SongID   int    `json:"song_id"`
	Content  string `json:"content"` // Текст куплета
	Position int    `json:"position"`
}

type SongS struct {
	ID        int    `json:"id"`
	GroupName string `json:"group_name"`
	Title     string `json:"title"`
}

// SongResponse используется для возврата данных песни в API.
type SongResponse struct {
	ID        int    `json:"id"`
	GroupName string `json:"group_name"`
	Title     string `json:"title"`
}

// ErrorResponse используется для возврата ошибок в API.
type ErrorResponse struct {
	Message string `json:"message"`
}
