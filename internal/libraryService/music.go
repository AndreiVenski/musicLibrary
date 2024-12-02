package libraryService

import (
	"context"
	"songsLibrary/internal/models"
)

type Music interface {
	GetSongDetail(ctx context.Context, songData *models.SongRequest) (*models.SongDetails, error)
}
