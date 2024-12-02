package libraryService

import (
	"context"
	"songsLibrary/internal/models"
)

type UseCase interface {
	AddSong(ctx context.Context, songData *models.SongRequest) (*models.SongResponse, error)
}
