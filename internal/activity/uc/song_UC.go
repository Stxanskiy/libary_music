package uc

import (
	"context"
	"fmt"
	_interface "libary_music/internal/activity/interface"
	"libary_music/internal/activity/model"
	"libary_music/pkg/api/musiclibary"
)

// представляет интерфейс работы бизнес-логики для работы с песнями
type SongUC struct {
	songRepo           _interface.SongRepo
	musicLibraryClient *musiclibary.MusicLibraryClient
}

// создает новый экземпляр SongUC

func NewSongUC(repo _interface.SongRepo, client *musiclibary.MusicLibraryClient) *SongUC {
	return &SongUC{
		songRepo:           repo,
		musicLibraryClient: client,
	}
}

/*func NewSongUC(songRepo _interface.SongRepo) *SongUC {
	return &SongUC{songRepo: songRepo}
}*/

// Добавление новой песни в библиотеку
func (uc *SongUC) AddSong(ctx context.Context, params *model.Song) (int, error) {
	// Запрос информации о песне через внешний API.
	apiResponse, err := uc.musicLibraryClient.GetSongDetail(params.GroupName, params.Title)
	if err != nil {
		return 0, fmt.Errorf("ошибка получения информации о песне: %w", err)
	}

	// Обогащение данных о песне.
	params.ReleaseDate = apiResponse.ReleaseDate
	params.Lyrics = apiResponse.Text
	params.Link = apiResponse.Link

	// Вставка новой песни в БД.
	id, err := uc.songRepo.AddSong(ctx, params)
	if err != nil {
		return 0, fmt.Errorf("ошибка добавления песни в БД: %w", err)
	}

	return id, nil
}

/*func (uc *SongUC) AddSong(ctx context.Context, song *model.Song) (int, error) {
	return uc.songRepo.AddSong(ctx, song)
}*/

func (uc *SongUC) GetSongByID(ctx context.Context, id int) (*model.Song, error) {
	return uc.songRepo.GetSongByID(ctx, id)
}

// Обновление песни
func (uc *SongUC) UpdateSong(ctx context.Context, song *model.Song) (int, error) {
	return uc.songRepo.UpdateSong(ctx, song)
}

// Удалене песни
func (uc *SongUC) DeleteSong(ctx context.Context, id int) (int, error) {
	return uc.songRepo.DeleteSong(ctx, id)
}

func (uc *SongUC) ListSongsWithPagination(ctx context.Context, limit, offset int) ([]model.Song, error) {
	return uc.songRepo.ListSongsWithPagination(ctx, limit, offset)
}
