package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"songsLibrary/internal/libraryService"
	"songsLibrary/internal/models"
)

type libRepo struct {
	db *sqlx.DB
}

func NewLibRepository(db *sqlx.DB) libraryService.Repository {
	return &libRepo{
		db: db,
	}
}

func (r *libRepo) IsSongExists(ctx context.Context, song *models.SongRequest) (bool, error) {
	var exists bool
	err := r.db.GetContext(ctx, &exists, isExistsMusicQuery, song.Group, song.Song)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, errors.Wrap(err, "libRepo.IsSongExists.Get")
	}
	return exists, nil
}

func (r *libRepo) CreateSong(ctx context.Context, songData *models.SongRequest, songDetails *models.SongDetails) (*models.SongResponse, error) {
	songResp := &models.SongResponse{}
	if err := r.db.QueryRowxContext(ctx, writeSongQuery,
		songData.Song, songData.Group,
		songDetails.ReleaseDate, songDetails.Text,
		songDetails.Link,
	).StructScan(songResp); err != nil {
		return nil, errors.Wrap(err, "libRepo.CreateSong.StructScan")
	}

	return songResp, nil
}

func (r *libRepo) UpdateSong(ctx context.Context, songData *models.SongFullDataRequest) (*models.SongResponse, error) {
	songResp := &models.SongResponse{}
	if err := r.db.QueryRowxContext(ctx, updateSongQuery, songData.ReleaseDate,
		songData.Text, songData.Link,
		songData.Group, songData.Song,
	).StructScan(songResp); err != nil {
		return nil, errors.Wrap(err, "libRepo.UpdateSong.StructScan")
	}

	return songResp, nil
}
