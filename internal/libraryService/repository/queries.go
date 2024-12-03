package repository

const (
	isExistsMusicQuery = `SELECT EXISTS (SELECT 1 FROM songs WHERE "group"=$1 AND song=$2)`

	writeSongQuery = `INSERT INTO songs (group, song, releaseDate, text, link) 
						  VALUES ($1, $2, $3, $4, $5) RETURNING *`

	updateSongQuery = `UPDATE songs 
						  SET releaseDate = COALESCE($1, releaseDate), 
						      text = COALESCE($2, text),
						      link = COALESCE($3, link)
						  WHERE "group" = $4 AND song = $5 
						  RETUTNING *
                		 `

	updateSongByIDQuery = `UPDATE songs 
						  SET group = COALESCE($1, "group")
						  	  song = COALESCE($2, song)
						      releaseDate = COALESCE($3, releaseDate), 
						      text = COALESCE($4, text),
						      link = COALESCE($5, link)
						  WHERE songId=$6
						  RETUTNING *
                		 `

	deleteSongQuery = `DELETE FROM songs WHERE group=$1 AND song=$2`

	deleteSongByIDQuery = `DELETE FROM songs WHERE songId=$1`

	getLibraryInfoQuery = `SELECT songId, "group", song, releaseDate, text, link
							FROM songs
							WHERE "group" ILIKE COALESCE('%' || $1 || '$%', '%') AND 
							      song ILIKE COALESCE('%' || $2 || '$%', '%') AND
								  releaseDate ILIKE COALESCE('%' || $3 || '$%', '%') AND
								  text ILIKE COALESCE('%' || $4 || '$%', '%') AND
								  link ILIKE COALESCE('%' || $5 || '$%', '%') AND
							ORDER BY id
							LIMIT $6 OFFSET $7
					`

	getSongVerseQuery = `WITH split_lyrics AS (
						  SELECT 
							  UNNEST(string_to_array(text, '\n\n')) WITH ORDINALITY AS (verse, verse_number)
						  FROM songs
						  WHERE "group" = $1 AND song = $2
					       )
							SELECT verse
							FROM split_lyrics
							WHERE verse_number = $3
						`
)
