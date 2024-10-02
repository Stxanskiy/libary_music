package uc

import (
	"context"
	_interface "libary_music/internal/activity/interface"
	"libary_music/internal/activity/model"
)

// представляет интерфейс работы бизнес-логики для работы с песнями
type SongUC struct {
	songRepo _interface.SongRepo
}

// создает новый экземпляр SongUC
func NewSongUC(songRepo _interface.SongRepo) *SongUC {
	return &SongUC{songRepo: songRepo}
}

// Добавление новой песни в библиотеку
func (uc *SongUC) AddSong(ctx context.Context, song *model.Song) (int, error) {
	return uc.songRepo.AddSong(ctx, song)
}

func (uc *SongUC) GetSongByID(ctx context.Context, id int) (*model.Song, error) {
	return uc.GetSongByID(ctx, id)
}

// Обновление песни
func (uc *SongUC) UpdateSong(ctx context.Context, song *model.Song) error {
	return uc.UpdateSong(ctx, song)
}

// Удалене песни
func (uc *SongUC) DeleteSong(ctx context.Context, id int) error {
	return uc.DeleteSong(ctx, id)
}

func (uc *SongUC) ListSongsWithPagination(ctx context.Context, limit, offset int) ([]model.Song, error) {
	return uc.ListSongsWithPagination(ctx, limit, offset)
}
