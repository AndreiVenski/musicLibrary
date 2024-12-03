package models

type GroupField struct {
	Group string `json:"group" db:"group"`
}

type SongField struct {
	Song string `json:"song" db:"song"`
}

type ReleaseDateField struct {
	ReleaseDate string `json:"releaseDate" db:"releaseDate"`
}

type TextField struct {
	Text string `json:"text" db:"text"`
}

type LinkField struct {
	Link string `json:"link" db:"link"`
}

type SongRequest struct {
	GroupField `validate:"required"`
	SongField  `validate:"required"`
}

type SongDetails struct {
	ReleaseDateField `validate:"required"`
	TextField        `validate:"required"`
	LinkField        `validate:"required"`
}

type SongFullDataRequest struct {
	SongRequest
	ReleaseDateField
	TextField
	LinkField
}

type SongFullDataWithLimitAndOffsetRequest struct {
	GroupField
	SongField
	ReleaseDateField
	TextField
	LinkField
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type SongFullDataRequestWithID struct {
	ID int `json:"songId" db:"songId" validate:"required"`
	GroupField
	SongField
	ReleaseDateField
	TextField
	LinkField
}

type SongResponse struct {
	ID               int `json:"songId" db:"songId"`
	GroupField       `validate:"required"`
	SongField        `validate:"required"`
	ReleaseDateField `validate:"required"`
	TextField        `validate:"required"`
	LinkField        `validate:"required"`
}
