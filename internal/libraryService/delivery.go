package libraryService

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GetLibraryInfo(ctx *fiber.Ctx) error
	GetMusicTextInfo(ctx *fiber.Ctx) error
	DeleteMusic(ctx *fiber.Ctx) error
	UpdateMusicInfo(ctx *fiber.Ctx) error
	AddNewMusic(ctx *fiber.Ctx) error
}
