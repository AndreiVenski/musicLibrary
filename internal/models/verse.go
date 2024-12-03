package models

type VerseRequest struct {
	Group   string `json:"group" db:"group"`
	Song    string `json:"song" db:"song"`
	VerseID int    `json:"verseID" db:"verseID"`
}
type VerseResponse struct {
	VerseID int    `json:"verseID" db:""`
	Verse   string `json:"verse" db:""`
}
