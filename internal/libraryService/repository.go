package libraryService

import (
	"context"
	"songsLibrary/internal/models"
)

type Repository interface {
	IsSongExists(ctx context.Context, song *models.SongRequest) (bool, error)
	CreateSong(ctx context.Context, songData *models.SongRequest, songDetails *models.SongDetails) (*models.SongResponse, error)
	UpdateSong(ctx context.Context, songData *models.SongFullDataRequest) (*models.SongResponse, error)
}
