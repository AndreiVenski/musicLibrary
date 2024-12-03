package libraryService

import (
	"context"
	"songsLibrary/internal/models"
)

type Repository interface {
	IsSongExists(ctx context.Context, song *models.SongRequest) (bool, error)
	CreateSong(ctx context.Context, songData *models.SongRequest, songDetails *models.SongDetails) (*models.SongResponse, error)
	UpdateSong(ctx context.Context, songData *models.SongFullDataRequest) (*models.SongResponse, error)
	UpdateSongByID(ctx context.Context, songData *models.SongFullDataRequestWithID) (*models.SongResponse, error)
	DeleteSong(ctx context.Context, songData *models.SongRequest) (bool, error)
	DeleteSongByID(ctx context.Context, songID int) (bool, error)
	GetLibraryInfo(ctx context.Context, songData *models.SongFullDataWithLimitAndOffsetRequest) ([]*models.SongResponse, error)
	GetSongVerse(ctx context.Context, verseData *models.VerseRequest) (*models.VerseResponse, error)
}
