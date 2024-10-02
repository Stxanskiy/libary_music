package _interface

import (
	"context"
	"libary_music/internal/activity/model"
)

type SongRepo interface {
	AddSong(ctx context.Context, params *model.Song) (int, error)
	GetSongByID(ctx context.Context, id int) (*model.Song, error)
	UpdateSong(ctx context.Context, song *model.Song) (int, error)
	DeleteSong(ctx context.Context, id int) (int, error)
	ListSongsWithPagination(ctx context.Context, limit, offset int) ([]model.Song, error)
}
