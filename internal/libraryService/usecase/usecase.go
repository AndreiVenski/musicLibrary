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

func (uc *libUC) UpdateSong(ctx context.Context, songFullData *models.SongFullDataRequest) (*models.SongResponse, error) {
	exists, err := uc.libRepo.IsSongExists(ctx, &models.SongRequest{Group: songFullData.Group, Song: songFullData.Song})
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("song doesn't exist in database")
	}

	songResp, err := uc.libRepo.

}
