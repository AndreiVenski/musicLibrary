package http

import (
	"github.com/gofiber/fiber/v2"
	"songsLibrary/internal/libraryService"
)

func MapLibRoutes(api fiber.Router, handlers libraryService.Handler) {
	api.Get("/info", handlers.GetLibraryInfo)
	api.Get("/music/text", handlers.GetMusicTextInfo)
	api.Post("/music", handlers.AddNewMusic)
	api.Put("/music/update", handlers.UpdateMusicInfo)
	api.Put("/music/update/id", handlers.UpdateMusicInfoByID)
	api.Delete("/music", handlers.DeleteMusic)
	api.Delete("/music/id", handlers.DeleteMusicByID)
}
