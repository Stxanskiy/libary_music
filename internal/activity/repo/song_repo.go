package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"libary_music/internal/activity/model"
	"libary_music/pkg/storage"
)

// Структура репозитория песен.
type songRepo struct {
	pg *pgxpool.Pool
}

// NewSongRepo создает новый репозиторий песен.
func NewSongRepo(db *storage.DB) *songRepo {
	// Используем db.Pool для получения *pgxpool.Pool из *storage.DB.
	return &songRepo{pg: db.Pool}
}

// SQL-запросы.
const (
	addSongSQL = `
		INSERT INTO public.song (music_band_id, title, release_date, lyrics, link) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING song_id;`

	getSongByIDSQL = `
		SELECT song_id, music_band_id, title, release_date, lyrics, link 
		FROM public.song 
		WHERE song_id = $1;`

	updateSongSQL = `
		UPDATE public.song 
		SET music_band_id = $1, title = $2, release_date = $3, lyrics = $4, link = $5 
		WHERE song_id = $6 
		RETURNING song_id;`

	deleteSongSQL = `
		DELETE FROM public.song 
		WHERE song_id = $1 
		RETURNING song_id;`

	listSongSQL = `
		SELECT song_id, music_band_id, title, release_date, lyrics, link 
		FROM public.song
		ORDER BY release_date DESC 
		limit $1 offset $2;`
)

// Реализация методов интерфейса SongRepo.

func (r *songRepo) AddSong(ctx context.Context, params *model.Song) (int, error) {
	var songID int
	err := r.pg.QueryRow(ctx, addSongSQL, params.MusicBandID, params.Title, params.ReleaseDate, params.Lyrics, params.Link).Scan(&songID)
	if err != nil {
		return 0, err
	}
	return songID, nil
}

func (r *songRepo) GetSongByID(ctx context.Context, id int) (*model.Song, error) {
	var song model.Song
	err := r.pg.QueryRow(ctx, getSongByIDSQL, id).Scan(
		&song.SongID, &song.MusicBandID, &song.Title, &song.ReleaseDate, &song.Lyrics, &song.Link)
	if err != nil {
		return nil, err
	}
	return &song, nil
}

func (r *songRepo) DeleteSong(ctx context.Context, id int) (int, error) {
	var deletedID int
	err := r.pg.QueryRow(ctx, deleteSongSQL, id).Scan(&deletedID)
	if err != nil {
		return 0, err
	}
	return deletedID, nil
}

func (r *songRepo) UpdateSong(ctx context.Context, song *model.Song) (int, error) {
	var songID int
	err := r.pg.QueryRow(ctx, updateSongSQL, song.MusicBandID, song.Title, song.ReleaseDate, song.Lyrics, song.Link, song.SongID).Scan(&songID)
	if err != nil {
		return 0, err
	}
	return songID, nil
}

func (r *songRepo) ListSongsWithPagination(ctx context.Context, limit, offset int) ([]model.Song, error) {
	rows, err := r.pg.Query(ctx, listSongSQL, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []model.Song
	for rows.Next() {
		var song model.Song
		if err := rows.Scan(&song.SongID, &song.MusicBandID, &song.Title, &song.ReleaseDate, &song.Lyrics, &song.Link); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}
