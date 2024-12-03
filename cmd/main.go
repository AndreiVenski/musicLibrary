package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/pressly/goose"
	"os"
	_ "songsLibrary/api/docs"
	"songsLibrary/config"
	server2 "songsLibrary/internal/server"
	"songsLibrary/pkg/db/pg_conn"
	logger1 "songsLibrary/pkg/logger"
)

// @title Song library API
// @version 1.0
// @description This is API for song library
// @contact.name Andrei Venski
// @contact.url https://github.com/andrew967
// @contact.email venskiandrei32@gmail.com
// @BasePath /lib
func main() {

	cfg, err := config.InitConfig(".env")
	if err != nil {
		log.Error("Config init failed", err)
		os.Exit(1)
	}

	logger := logger1.NewApiLogger(cfg)
	logger.InitLogger()

	db, err := pg_conn.NewPsqlDB(cfg)
	if err != nil {
		logger.Error("DB init failed", err)
		os.Exit(1)
	}
	defer db.Close()

	if err = goose.Up(db.DB, "migrations"); err != nil {
		logger.Error("Migrations up failed", err)
		os.Exit(1)
	}

	fiberApp := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: false,
	})
	server := server2.NewServer(db, cfg, fiberApp, logger)

	if err = server.Run(); err != nil {
		os.Exit(0)
	}

}
