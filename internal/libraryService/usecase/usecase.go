package usecase

import (
	"context"
	"github.com/pkg/errors"
	"songsLibrary/config"
	"songsLibrary/internal/libraryService"
	"songsLibrary/internal/models"
	logger1 "songsLibrary/pkg/logger"
)

type libUC struct {
	libRepo  libraryService.Repository
	libMusic libraryService.Music
	logger   logger1.Logger
	cfg      *config.Config
}

func NewLibUseCase(libRepo libraryService.Repository, logger logger1.Logger, cfg *config.Config, libMusic libraryService.Music) libraryService.UseCase {
	return &libUC{
		libRepo:  libRepo,
		logger:   logger,
		cfg:      cfg,
		libMusic: libMusic,
	}
}

func (uc *libUC) AddSong(ctx context.Context, songData *models.SongRequest) (*models.SongResponse, error) {
	exists, err := uc.libRepo.IsSongExists(ctx, songData)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("song exists in database")
	}

	songDetails, err := uc.libMusic.GetSongDetail(ctx, songData)
	if err != nil {
		return nil, err
	}

	songResp, err := uc.libRepo.CreateSong(ctx, songData, songDetails)
	if err != nil {
		return nil, err
	}

	return songResp, err
}

func (uc *libUC) UpdateSongDetails(ctx context.Context, songFullData *models.SongFullDataRequest) (*models.SongResponse, error) {
	songResp, err := uc.libRepo.UpdateSong(ctx, songFullData)
	if err != nil {
		return nil, err
	}

	return songResp, nil
}

func (uc *libUC) UpdateSongByID(ctx context.Context, songFullDataWithID *models.SongFullDataRequestWithID) (*models.SongResponse, error) {
	songResp, err := uc.libRepo.UpdateSongByID(ctx, songFullDataWithID)
	if err != nil {
		return nil, err
	}
	return songResp, nil
}

func (uc *libUC) DeleteSong(ctx context.Context, songData *models.SongRequest) error {
	deleted, err := uc.libRepo.DeleteSong(ctx, songData)
	if err != nil {
		return err
	}

	if !deleted {
		return errors.New("song not found")
	}

	return nil
}

func (uc *libUC) DeleteSongByID(ctx context.Context, songID int) error {
	deleted, err := uc.libRepo.DeleteSongByID(ctx, songID)
	if err != nil {
		return err
	}

	if !deleted {
		return errors.New("song not found")
	}

	return nil
}

func (uc *libUC) GetLibraryInfo(ctx context.Context, filter *models.SongFullDataWithLimitAndOffsetRequest) ([]*models.SongResponse, error) {
	songsResp, err := uc.libRepo.GetLibraryInfo(ctx, filter)
	if err != nil {
		return nil, err
	}

	return songsResp, nil
}

func (uc *libUC) GetSongVerse(ctx context.Context, verseInfo *models.VerseRequest) (*models.VerseResponse, error) {
	verseResp, err := uc.libRepo.GetSongVerse(ctx, verseInfo)
	if err != nil {
		return nil, err
	}
	return verseResp, nil
}
