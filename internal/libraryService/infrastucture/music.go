package infrastucture

import (
	"songsLibrary/config"
	"songsLibrary/internal/libraryService"
)

type libMusic struct {
	baseURL    string
	httpClient libraryService.HttpClient
	cfg        *config.Config
}

func NewLibMusic(baseURL string, httpClient libraryService.HttpClient, cfg *config.Config) libraryService.Music {
	return &libMusic{
		baseURL:    baseURL,
		httpClient: httpClient,
		cfg:        cfg,
	}
}
