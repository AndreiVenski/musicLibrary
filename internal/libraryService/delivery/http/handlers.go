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
// @Failure 400 {object} map[string]interface{} "Invalid request data"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /info [post]
func (h *libHandler) GetLibraryInfo(ctx *fiber.Ctx) error {
	h.logger.Infof("GetLibraryInfo request starts")

	filter := &models.SongFullDataWithLimitAndOffsetRequest{}
	if err := utils.ReadFromRequest(ctx, filter); err != nil {
		h.logger.Errorf("Invalid request data: %v", err)
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}
	h.logger.Debugf("Extracted parameters: %+v", filter)
	if filter.Limit == 0 {
		filter.Limit = 5
		h.logger.Debugf("Limit not provided, default to %d", filter.Limit)
	}

	songsResp, err := h.libUC.GetLibraryInfo(ctx.UserContext(), filter)
	if err != nil {
		if httpErrors.IsServiceError(err) {
			h.logger.Errorf("Failed to get songs with filter: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not get songs with filter")
		}
		h.logger.Warnf("Bad request when get library info: %v", err)
		return respondWithError(ctx, fiber.StatusBadRequest, err.Error())
	}
	h.logger.Debugf("Processing result: %+v", songsResp)
	h.logger.Infof("Successfully get library info")
	return ctx.Status(fiber.StatusOK).JSON(songsResp)
}

// @Summary Get music text information
// @Description Retrieve verse information for a specific song.
// @Tags Library
// @Accept json
// @Produce json
// @Param verseReq body models.VerseRequest true "Verse request"
// @Success 200 {object} models.VerseResponse "Verse information"
// @Failure 400 {object} map[string]interface{} "Invalid request data"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /music/text [post]
func (h *libHandler) GetMusicTextInfo(ctx *fiber.Ctx) error {
	h.logger.Infof("Get GetMusicTextInfo request")
	verseReq := &models.VerseRequest{}
	if err := utils.ReadFromRequest(ctx, verseReq); err != nil {
		h.logger.Errorf("Invalid request data")
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}
	h.logger.Debugf("Extracted parameters: %+v", verseReq)
	verseResp, err := h.libUC.GetSongVerse(ctx.UserContext(), verseReq)
	if err != nil {
		if httpErrors.IsServiceError(err) {
			h.logger.Errorf("Failed to get verse info: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not get verse information")
		}
		h.logger.Warnf("Bad request when getting verse info: %v", err)
		return respondWithError(ctx, fiber.StatusBadRequest, err.Error())
	}
	h.logger.Debugf("Processing result: %+v", verseResp)
	h.logger.Info("Successfully getted verse: %s", verseResp.Verse)
	return ctx.Status(fiber.StatusOK).JSON(verseResp)
}

// @Summary Delete a song
// @Description Delete a song from the library using song details.
// @Tags Library
// @Accept json
// @Produce json
// @Param songData body models.SongRequest true "Song to delete"
// @Success 200 {object} map[string]interface{} "Song deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request data"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /music [delete]
func (h *libHandler) DeleteMusic(ctx *fiber.Ctx) error {
	h.logger.Infof("Get DeleteMusic request")
	songData := &models.SongRequest{}
	err := utils.ReadFromRequest(ctx, songData)
	if err != nil {
		h.logger.Errorf("Invalid request data")
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}
	h.logger.Debugf("Extracted parameters: %+v", songData)
	err = h.libUC.DeleteSong(ctx.UserContext(), songData)
	if err != nil {
		if httpErrors.IsServiceError(err) {
			h.logger.Errorf("Failed to delete song: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not delete song")
		}
		h.logger.Warnf("Bad request when deleting song: %v", err)
		return respondWithError(ctx, fiber.StatusBadRequest, err.Error())
	}
	h.logger.Info("Successfully deleted song with name: %s", songData.Song)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "song deleted"})
}

// @Summary Delete a song by ID
// @Description Delete a song from the library using its unique ID.
// @Tags Library
// @Accept json
// @Produce json
// @Param songID body models.SongID true "Song ID to delete"
// @Success 200 {object} map[string]interface{} "Song deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request data"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /music/id [delete]
func (h *libHandler) DeleteMusicByID(ctx *fiber.Ctx) error {
	h.logger.Infof("Get DeleteMusicByID request")
	songID := &struct {
		ID int `json:"songId"`
	}{}
	err := utils.ReadFromRequest(ctx, songID)
	if err != nil {
		h.logger.Errorf("Invalid request data")
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}
	h.logger.Debugf("Extracted parameters: %+v", songID)
	err = h.libUC.DeleteSongByID(ctx.UserContext(), songID.ID)
	if err != nil {
		if httpErrors.IsServiceError(err) {
			h.logger.Errorf("Failed to delete song by id: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not delete song by id")
		}
		h.logger.Warnf("Bad request when deleting song by id: %v", err)
		return respondWithError(ctx, fiber.StatusBadRequest, err.Error())
	}

	h.logger.Info("Successfully deleted song with id: %v", songID.ID)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "song deleted"})
}

// @Summary Update music information
// @Description Update details of a song in the library.
// @Tags Library
// @Accept json
// @Produce json
// @Param updateData body models.SongFullDataRequest true "Song data to update"
// @Success 200 {object} models.SongResponse "Updated song information"
// @Failure 400 {object} map[string]interface{} "Invalid request data"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /music/update [put]
func (h *libHandler) UpdateMusicInfo(ctx *fiber.Ctx) error {
	h.logger.Infof("Get UpdateMusicInfo request")
	updateData := &models.SongFullDataRequest{}
	err := utils.ReadFromRequest(ctx, updateData)
	if err != nil {
		h.logger.Errorf("Invalid request data")
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}
	h.logger.Debugf("Extracted parameters: %+v", updateData)
	updatedSong, err := h.libUC.UpdateSongDetails(ctx.UserContext(), updateData)
	if err != nil {
		if httpErrors.IsServiceError(err) {
			h.logger.Errorf("Failed to update song details: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not update song details")
		}
		h.logger.Warnf("Bad request when updating song details: %v", err)
		return respondWithError(ctx, fiber.StatusBadRequest, err.Error())
	}
	h.logger.Debugf("Processing result: %+v", updatedSong)
	h.logger.Info("Successfully updated song with id : %v", updatedSong.ID)
	return ctx.Status(fiber.StatusOK).JSON(updatedSong)
}

// @Summary Update music information by ID
// @Description Update details of a song in the library using its unique ID.
// @Tags Library
// @Accept json
// @Produce json
// @Param updateData body models.SongFullDataRequestWithID true "Song data to update"
// @Success 200 {object} models.SongResponse "Updated song information"
// @Failure 400 {object} map[string]interface{} "Invalid request data"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /music/update/id [put]
func (h *libHandler) UpdateMusicInfoByID(ctx *fiber.Ctx) error {
	h.logger.Infof("Get UpdateMusicInfoByID request")
	updateData := &models.SongFullDataRequestWithID{}
	err := utils.ReadFromRequest(ctx, updateData)
	if err != nil {
		h.logger.Errorf("Invalid request data")
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}
	h.logger.Debugf("Extracted parameters: %+v", updateData)
	updatedSong, err := h.libUC.UpdateSongByID(ctx.UserContext(), updateData)
	if err != nil {
		if httpErrors.IsServiceError(err) {
			h.logger.Errorf("Failed to update song: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not update song")
		}
		h.logger.Warnf("Bad request when updating song by id: %v", err)
		return respondWithError(ctx, fiber.StatusBadRequest, err.Error())
	}
	h.logger.Debugf("Processing result: %+v", updatedSong)
	h.logger.Info("Successfully updated song with id: %v", updateData.ID)
	return ctx.Status(fiber.StatusOK).JSON(updatedSong)
}

// @Summary Add a new song
// @Description Add a new song to the music library.
// @Tags Library
// @Accept json
// @Produce json
// @Param songData body models.SongRequest true "Song data to add"
// @Success 200 {object} models.SongResponse "Added song information"
// @Failure 400 {object} map[string]interface{} "Invalid request data"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /music [post]
func (h *libHandler) AddNewMusic(ctx *fiber.Ctx) error {
	h.logger.Infof("Get AddNewMusic request")
	songData := &models.SongRequest{}
	err := utils.ReadFromRequest(ctx, songData)
	if err != nil {
		h.logger.Errorf("Invalid request data %v", err)
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}
	h.logger.Debugf("Extracted parameters: %+v", songData)
	songResp, err := h.libUC.AddSong(ctx.UserContext(), songData)
	if err != nil {
		if httpErrors.IsServiceError(err) {
			h.logger.Errorf("Failed to add song: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not add song to db")
		}
		h.logger.Warnf("Bad request when adding new song: %v", err)
		return respondWithError(ctx, fiber.StatusBadRequest, err.Error())
	}
	h.logger.Debugf("Processing result: %+v", songResp)
	h.logger.Infof("Successfully added new song with ID: %v", songResp.ID)
	return ctx.Status(fiber.StatusOK).JSON(songResp)
}
