package repository

const (
	isExistsMusicQuery = `SELECT EXISTS (SELECT 1 FROM songs WHERE group=$1 AND song=$2)`

	writeSongQuery = `INSERT INTO songs (group, song, releaseDate, text, link) 
						  VALUES ($1, $2, $3, $4, $5) RETURNING *`

	updateSongQuery = `UPDATE songs 
						  SET releaseDate = COALESCE($1, releaseDate), 
						      text = COALESCE($2, text),
						      link = COALESCE($3, link)
						  WHERE "group" = $4 AND song = $5 
						  RETUTNING *
                		 `
)
