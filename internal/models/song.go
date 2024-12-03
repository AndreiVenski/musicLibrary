package models

type SongRequest struct {
	Group string `json:"group" db:"group" validate:"required"`
	Song  string `json:"song" db:"song" validate:"required"`
}

type SongFullDataRequest struct {
	Group       string `json:"group" db:"group" validate:"required"`
	Song        string `json:"song" db:"song" validate:"required"`
	ReleaseDate string `json:"releaseDate" db:"releaseDate"`
	Text        string `json:"text" db:"text"`
	Link        string `json:"link" db:"link"`
}

type SongResponse struct {
	Group       string `json:"group" db:"group" validate:"required"`
	Song        string `json:"song" db:"song" validate:"required"`
	ReleaseDate string `json:"releaseDate" db:"releaseDate" validate:"required"`
	Text        string `json:"text" db:"text" validate:"required"`
	Link        string `json:"link" db:"link" validate:"required"`
}

type SongDetails struct {
	ReleaseDate string `json:"releaseDate" db:"releaseDate" validate:"required"`
	Text        string `json:"text" db:"text" validate:"required"`
	Link        string `json:"link" db:"link" validate:"required"`
}
