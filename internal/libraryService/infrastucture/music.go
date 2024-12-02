package infrastucture

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"songsLibrary/config"
	"songsLibrary/internal/libraryService"
	"songsLibrary/internal/models"
	"songsLibrary/pkg/utils"
)

type libMusic struct {
	httpClient libraryService.HttpClient
	cfg        *config.Config
}

func NewLibMusic(httpClient libraryService.HttpClient, cfg *config.Config) libraryService.Music {
	return &libMusic{
		httpClient: httpClient,
		cfg:        cfg,
	}
}

func (m *libMusic) GetSongDetail(ctx context.Context, songData *models.SongRequest) (*models.SongDetails, error) {
	jsonData, err := json.Marshal(songData)
	if err != nil {
		return nil, errors.Wrap(err, "libMusic.GetSongDetail.Marshal")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, m.cfg.MusicService.MusicAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, errors.Wrap(err, "libMusic.GetSongDetail.NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "libMusic.GetSongDetail.NewRequest")
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
	case 400:
		return nil, errors.New("incorrect request")
	case 500:
		return nil, errors.New("API server doesn't work")
	default:
		return nil, errors.New("unknown response status code")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "libMusic.GetSongDetail.ReadAll")
	}

	songDetails := &models.SongDetails{}
	if err = json.Unmarshal(body, songDetails); err != nil {
		return nil, errors.Wrap(err, "libMusic.GetSongDetail.Unmarshal")
	}

	if err = utils.ValidateStruct(songDetails); err != nil {
		return nil, errors.New("validate failed")
	}
	return songDetails, nil
}
