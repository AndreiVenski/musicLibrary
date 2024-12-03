package libraryService

import (
	"context"
	"songsLibrary/internal/models"
)

type UseCase interface {
	AddSong(ctx context.Context, songData *models.SongRequest) (*models.SongResponse, error)
	UpdateSongDetails(ctx context.Context, songFullData *models.SongFullDataRequest) (*models.SongResponse, error)
	UpdateSongByID(ctx context.Context, songFullDataWithID *models.SongFullDataRequestWithID) (*models.SongResponse, error)
	DeleteSong(ctx context.Context, songData *models.SongRequest) error
	DeleteSongByID(ctx context.Context, songID int) error
	GetLibraryInfo(ctx context.Context, filter *models.SongFullDataWithLimitAndOffsetRequest) ([]*models.SongResponse, error)
	GetSongVerse(ctx context.Context, verseInfo *models.VerseRequest) (*models.VerseResponse, error)
}
