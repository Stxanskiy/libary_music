package model

import (
	"time"
)

type MusicBand struct {
	MusicBandID int    `json:"music_band_id"`
	Name        string `json:"name"`
}

type Song struct {
	SongID      int       `json:"song_id"`
	MusicBandID int       `json:"music_band_id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	Lyrics      string    `json:"lyrics"`
	Link        string    `json:"link"`
}

type Verse struct {
	VerseID  int    `json:"verse_id"`
	SongID   int    `json:"song_id"`
	Content  string `json:"content"`
	Position int    `json:"position"`
}
