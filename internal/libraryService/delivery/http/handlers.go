package http

import (
	"github.com/gofiber/fiber/v2"
	"songsLibrary/config"
	"songsLibrary/internal/libraryService"
	"songsLibrary/internal/models"
	"songsLibrary/pkg/httpErrors"
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

func respondWithError(ctx *fiber.Ctx, code int, message string) error {
	return ctx.Status(code).JSON(fiber.Map{
		"error": message,
	})
}

// @Summary Get library information
// @Description Retrieve information about the music library with optional pagination.
// @Tags Library
// @Accept json
// @Produce json
// @Param filter body models.SongFullDataWithLimitAndOffsetRequest true "Filter parameters"
// @Success 200 {object} models.SongsResponse "List of songs"
// @Failure 400 {object} fiber.Map "Invalid request data"
// @Failure 500 {object} fiber.Map "Internal server error"
// @Router /library/info [get]
func (h *libHandler) GetLibraryInfo(ctx *fiber.Ctx) error {
	filter := &models.SongFullDataWithLimitAndOffsetRequest{}
	if err := utils.ReadFromRequest(ctx, filter); err != nil {
		h.logger.Errorf("Invalid request data")
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}
	if filter.Limit == 0 {
		filter.Limit = 5
	}

	if filter.Offset == 0 {
		filter.Offset = 0
	}

	songsResp, err := h.libUC.GetLibraryInfo(ctx.UserContext(), filter)
	if err != nil {
		if httpErrors.IsServiceError(err) {
			h.logger.Errorf("Failed to get songs with filter: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not get songs with filter")
		}
		return respondWithError(ctx, fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(songsResp)
}

// @Summary Get music text information
// @Description Retrieve verse information for a specific song.
// @Tags Library
// @Accept json
// @Produce json
// @Param verseReq body models.VerseRequest true "Verse request"
// @Success 200 {object} models.VerseResponse "Verse information"
// @Failure 400 {object} fiber.Map "Invalid request data"
// @Failure 500 {object} fiber.Map "Internal server error"
// @Router /music/text [get]
func (h *libHandler) GetMusicTextInfo(ctx *fiber.Ctx) error {
	verseReq := &models.VerseRequest{}
	if err := utils.ReadFromRequest(ctx, verseReq); err != nil {
		h.logger.Errorf("Invalid request data")
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}

	verseResp, err := h.libUC.GetSongVerse(ctx.UserContext(), verseReq)
	if err != nil {
		if httpErrors.IsServiceError(err) {
			h.logger.Errorf("Failed to get verse info: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not get verse information")
		}
		return respondWithError(ctx, fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(verseResp)
}

// @Summary Delete a song
// @Description Delete a song from the library using song details.
// @Tags Library
// @Accept json
// @Produce json
// @Param songData body models.SongRequest true "Song to delete"
// @Success 200 {object} fiber.Map "Song deleted successfully"
// @Failure 400 {object} fiber.Map "Invalid request data"
// @Failure 500 {object} fiber.Map "Internal server error"
// @Router /music [delete]
func (h *libHandler) DeleteMusic(ctx *fiber.Ctx) error {
	songData := &models.SongRequest{}
	err := utils.ReadFromRequest(ctx, songData)
	if err != nil {
		h.logger.Errorf("Invalid request data")
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}

	err = h.libUC.DeleteSong(ctx.UserContext(), songData)
	if err != nil {
		if httpErrors.IsServiceError(err) {
			h.logger.Errorf("Failed to delete song: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not delete song")
		}
		return respondWithError(ctx, fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "song deleted"})
}

// DeleteMusicByID godoc
// @Summary Delete a song by ID
// @Description Delete a song from the library using its unique ID.
// @Tags Library
// @Accept json
// @Produce json
// @Param songID body struct{ ID int `json:"songId"` } true "Song ID to delete"
// @Success 200 {object} fiber.Map "Song deleted successfully"
// @Failure 400 {object} fiber.Map "Invalid request data"
// @Failure 500 {object} fiber.Map "Internal server error"
// @Router /music/id [delete]
func (h *libHandler) DeleteMusicByID(ctx *fiber.Ctx) error {
	songID := &struct {
		ID int `json:"songId""`
	}{}
	err := utils.ReadFromRequest(ctx, songID)
	if err != nil {
		h.logger.Errorf("Invalid request data")
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}

	err = h.libUC.DeleteSongByID(ctx.UserContext(), songID.ID)
	if err != nil {
		if httpErrors.IsServiceError(err) {
			h.logger.Errorf("Failed to delete song by id: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not delete song by id")
		}
		return respondWithError(ctx, fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "song deleted"})
}

// @Summary Update music information
// @Description Update details of a song in the library.
// @Tags Library
// @Accept json
// @Produce json
// @Param updateData body models.SongFullDataRequest true "Song data to update"
// @Success 200 {object} models.SongResponse "Updated song information"
// @Failure 400 {object} fiber.Map "Invalid request data"
// @Failure 500 {object} fiber.Map "Internal server error"
// @Router /music/update [put]
func (h *libHandler) UpdateMusicInfo(ctx *fiber.Ctx) error {
	updateData := &models.SongFullDataRequest{}
	err := utils.ReadFromRequest(ctx, updateData)
	if err != nil {
		h.logger.Errorf("Invalid request data")
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}

	updatedSong, err := h.libUC.UpdateSongDetails(ctx.UserContext(), updateData)
	if err != nil {
		if httpErrors.IsServiceError(err) {
			h.logger.Errorf("Failed to update song details: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not update song details")
		}
		return respondWithError(ctx, fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(updatedSong)
}

// UpdateMusicInfoByID godoc
// @Summary Update music information by ID
// @Description Update details of a song in the library using its unique ID.
// @Tags Library
// @Accept json
// @Produce json
// @Param updateData body models.SongFullDataRequestWithID true "Song data to update"
// @Success 200 {object} models.SongResponse "Updated song information"
// @Failure 400 {object} fiber.Map "Invalid request data"
// @Failure 500 {object} fiber.Map "Internal server error"
// @Router /music/update/id [put]
func (h *libHandler) UpdateMusicInfoByID(ctx *fiber.Ctx) error {
	updateData := &models.SongFullDataRequestWithID{}
	err := utils.ReadFromRequest(ctx, updateData)
	if err != nil {
		h.logger.Errorf("Invalid request data")
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}

	updatedSong, err := h.libUC.UpdateSongByID(ctx.UserContext(), updateData)
	if err != nil {
		if httpErrors.IsServiceError(err) {
			h.logger.Errorf("Failed to update song: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not update song")
		}
		return respondWithError(ctx, fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(updatedSong)
}

// AddNewMusic godoc
// @Summary Add a new song
// @Description Add a new song to the music library.
// @Tags Library
// @Accept json
// @Produce json
// @Param songData body models.SongRequest true "Song data to add"
// @Success 200 {object} models.SongResponse "Added song information"
// @Failure 400 {object} fiber.Map "Invalid request data"
// @Failure 500 {object} fiber.Map "Internal server error"
// @Router /music [post]
func (h *libHandler) AddNewMusic(ctx *fiber.Ctx) error {
	songData := &models.SongRequest{}
	err := utils.ReadFromRequest(ctx, songData)
	if err != nil {
		h.logger.Errorf("Invalid request data %v", err)
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}

	songResp, err := h.libUC.AddSong(ctx.UserContext(), songData)
	if err != nil {
		if httpErrors.IsServiceError(err) {
			h.logger.Errorf("Failed to add song: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not add song to db")
		}
		return respondWithError(ctx, fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(songResp)
}
