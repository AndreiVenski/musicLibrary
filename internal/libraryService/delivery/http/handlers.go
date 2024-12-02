package http

import (
	"songsLibrary/config"
	"songsLibrary/internal/libraryService"
	logger1 "songsLibrary/pkg/logger"
)

type libHandler struct {
	libUC  libraryService.Usecase
	cfg    *config.Config
	logger logger1.Logger
}

func NewLibHandler(libUC libraryService.Usecase, cfg *config.Config, logger logger1.Logger) libraryService.Handler {
	return libHandler{
		libUC:  libUC,
		cfg:    cfg,
		logger: logger,
	}
}
