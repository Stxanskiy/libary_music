package uc

import (
	"context"
	_interface "libary_music/internal/activity/interface"
	"libary_music/internal/activity/model"
)

type VerseUC struct {
	verseRepo _interface.VerseRepo
}

func NewVerseUC(verseRepo _interface.VerseRepo) *VerseUC {
	return &VerseUC{verseRepo: verseRepo}
}

func (uc *VerseUC) AddVerse(ctx context.Context, params *model.Verse) (int, error) {
	return uc.verseRepo.AddVerse(ctx, params)
}

func (uc *VerseUC) UpdateVerse(ctx context.Context, params *model.Verse) (int, error) {
	return uc.verseRepo.UpdateVerse(ctx, params)
}

func (uc *VerseUC) GetSongVerse(ctx context.Context, songID int, limit, offset int) ([]model.Verse, error) {
	return uc.verseRepo.GetSongVerse(ctx, songID, limit, offset)
}
