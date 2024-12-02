package http

import (
	"github.com/gofiber/fiber/v2"
	"songsLibrary/config"
	"songsLibrary/internal/libraryService"
	"songsLibrary/internal/models"
	logger1 "songsLibrary/pkg/logger"
	"songsLibrary/pkg/utils"
)

type libHandler struct {
	libUC  libraryService.UseCase
	cfg    *config.Config
	logger logger1.Logger
}

func NewLibHandler(libUC libraryService.UseCase, cfg *config.Config, logger logger1.Logger) libraryService.Handler {
	return &libHandler{
		libUC:  libUC,
		cfg:    cfg,
		logger: logger,
	}
}

func (h *libHandler) GetLibraryInfo(ctx *fiber.Ctx) error {
	return nil
}

func (h *libHandler) GetMusicTextInfo(ctx *fiber.Ctx) error {
	return nil
}

func (h *libHandler) DeleteMusic(ctx *fiber.Ctx) error {
	return nil
}

func (h *libHandler) UpdateMusicInfo(ctx *fiber.Ctx) error {
	return nil
}

func (h *libHandler) AddNewMusic(ctx *fiber.Ctx) error {
	songData := &models.SongRequest{}
	err := utils.ReadFromRequest(ctx, songData)
	if err != nil {
		return ctx.JSON(fiber.Map{})
	}

	songResp, err := h.libUC.AddSong(ctx.UserContext(), songData)
	if err != nil {
		return ctx.JSON(fiber.Map{})
	}

	return ctx.JSON(fiber.Map{"song": songResp})
}
