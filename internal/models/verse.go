package models

type VerseRequest struct {
	Group   string `json:"group" db:"group"`
	Song    string `json:"song" db:"song"`
	VerseID int    `json:"verseID" db:"verseID"`
}
type VerseResponse struct {
	Verse   string `json:"verse" db:"verse"`
	VerseID int    `json:"verseID" db:"verse_id"`
}
