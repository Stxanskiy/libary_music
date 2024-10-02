package uc

import (
	"context"
	_interface "libary_music/internal/activity/interface"
	"libary_music/internal/activity/model"
)

type VerseUC struct {
	verseRepo _interface.VerseRepo
}

func NewRepoUC(verseRepo _interface.VerseRepo) *VerseUC {
	return &VerseUC{verseRepo: verseRepo}
}

func (uc *VerseUC) AddVerse(ctx context.Context, verse *model.Verse) (int, error) {
	return uc.verseRepo.AddVerse(ctx, verse)
}
