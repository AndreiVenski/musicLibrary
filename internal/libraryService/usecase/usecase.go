package usecase

import (
	"songsLibrary/config"
	"songsLibrary/internal/libraryService"
	logger1 "songsLibrary/pkg/logger"
)

type libUC struct {
	libRepo libraryService.Usecase
	logger  logger1.Logger
	cfg     *config.Config
}

func NewLibUseCase(libRepo libraryService.Repository, logger logger1.Logger, cfg *config.Config) libraryService.Usecase {
	return &libUC{
		libRepo: libRepo,
		logger:  logger,
		cfg:     cfg,
	}
}
