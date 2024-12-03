package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"songsLibrary/internal/libraryService"
	"songsLibrary/internal/models"
	"songsLibrary/pkg/httpErrors"
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, httpErrors.NotFoundSongError
		}
		return nil, errors.Wrap(err, "libRepo.UpdateSong.StructScan")
	}

	return songResp, nil
}

func (r *libRepo) UpdateSongByID(ctx context.Context, songData *models.SongFullDataRequestWithID) (*models.SongResponse, error) {
	songResp := &models.SongResponse{}
	if err := r.db.QueryRowxContext(ctx, updateSongByIDQuery, songData.Group, songData.Song,
		songData.ReleaseDate, songData.Text, songData.Link, songData.ID).StructScan(songResp); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, httpErrors.NotFoundSongError
		}
		return nil, errors.New("libRepo.UpdateSongByID.StructScan")
	}
	return songResp, nil
}

func (r *libRepo) DeleteSong(ctx context.Context, songData *models.SongRequest) (bool, error) {
	result, err := r.db.ExecContext(ctx, deleteSongQuery, songData.Group, songData.Song)
	if err != nil {
		return false, errors.Wrap(err, "libRepo.DeleteSong.ExecContext")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, errors.Wrap(err, "libRepo.DeleteSong.ExecContext")
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (r *libRepo) DeleteSongByID(ctx context.Context, songID int) (bool, error) {
	result, err := r.db.ExecContext(ctx, deleteSongByIDQuery, songID)
	if err != nil {
		return false, errors.Wrap(err, "libRepo.DeleteSongByID.ExecContext")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {

		return false, errors.Wrap(err, "libRepo.DeleteSongByID.ExecContext")
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (r *libRepo) GetLibraryInfo(ctx context.Context, songData *models.SongFullDataWithLimitAndOffsetRequest) ([]*models.SongResponse, error) {
	var songsResp []*models.SongResponse

	err := r.db.SelectContext(ctx, &songsResp, getLibraryInfoQuery, songData.Group, songData.Song,
		songData.ReleaseDate, songData.Text, songData.Link, songData.Limit, songData.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "libRepo.GetLibraryInfo.SelectContext")
	}

	return songsResp, nil
}

func (r *libRepo) GetSongVerse(ctx context.Context, verseData *models.VerseRequest) (*models.VerseResponse, error) {
	verse := &models.VerseResponse{}
	if verseData.VerseID == 0 {
		verseData.VerseID = 1
	}
	if err := r.db.GetContext(ctx, verse, getSongVerseQuery, verseData.VerseID, verseData.Group, verseData.Song); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, httpErrors.NotFoundSongOrVerseError
		}
		return nil, errors.Wrap(err, "libRepo.GetSongVerse.Get")
	}

	return verse, nil
}
