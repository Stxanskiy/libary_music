package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"libary_music/internal/activity/model"
	"libary_music/pkg/storage"
)

type verseRepo struct {
	pg *pgxpool.Pool
}

func NewVerseRepo(pg *storage.DB) *verseRepo {
	return &verseRepo{pg: pg.Pool}

}

const (
	addVerseSQL = `
		INSERT INTO public.verse(verse_id, content, position) 
		VALUES ($1, $2, $3)
		returning verse_id;`

	updateVerseSQL = ` UPDATE public.verse
	SET content = $1, position = $2
	WHERE verse_id = $3
	RETURNING verse_id;`
)

func (r *verseRepo) AddVerse(ctx context.Context, params *model.Verse) (int, error) {
	var songID int
	err := r.pg.QueryRow(ctx, addVerseSQL, params.VerseID, params.Content, params.Position).Scan(params.SongID)
	if err != nil {
		return 0, err
	}
	return songID, nil

}

func (r *verseRepo) UpdateVerse(ctx context.Context, params *model.Verse) (int, error) {
	var verseID int
	err := r.pg.QueryRow(ctx, updateVerseSQL, params.Content, params.Position).Scan(&verseID)
	if err != nil {
		return 0, err
	}
	return verseID, nil
}
