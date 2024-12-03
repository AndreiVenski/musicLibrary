package server

import (
	"github.com/gofiber/swagger"
	"songsLibrary/internal/libraryService/delivery/http"
	"songsLibrary/internal/libraryService/infrastucture"
	"songsLibrary/internal/libraryService/repository"
	"songsLibrary/internal/libraryService/usecase"
)

func (s *Server) MapHandlers() {
	libRepo := repository.NewLibRepository(s.db)

	client := infrastucture.NewHTTPClient(s.cfg)
	libMusic := infrastucture.NewLibMusic(client, s.cfg)

	libUseCase := usecase.NewLibUseCase(libRepo, s.logger, s.cfg, libMusic)

	libHandler := http.NewLibHandler(libUseCase, s.cfg, s.logger)

	s.fiber.Get("/swagger/*", swagger.HandlerDefault)

	lib := s.fiber.Group("/lib")

	http.MapLibRoutes(lib, libHandler)
}
