package _interface

import (
	"context"
	"libary_music/internal/activity/model"
)

type VerseRepo interface {
	AddVerse(ctx context.Context, params *model.Verse) (int, error)
	UpdateVerse(ctx context.Context, params *model.Verse) (int, error)
	GetSongVerse(ctx context.Context, params *model.Verse) (int, error)
}
