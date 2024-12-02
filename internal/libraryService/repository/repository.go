package repository

import (
	"github.com/jmoiron/sqlx"
	"songsLibrary/internal/libraryService"
)

type libRepo struct {
	db *sqlx.DB
}

func NewLibRepository(db *sqlx.DB) libraryService.Repository {
	return &libRepo{
		db: db,
	}
}
