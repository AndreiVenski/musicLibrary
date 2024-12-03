package repository

const (
	isExistsMusicQuery = `SELECT EXISTS (SELECT 1 FROM songs WHERE "group"=$1 AND song=$2)`

	writeSongQuery = `INSERT INTO songs (group, song, release_date, text, link) 
						  VALUES ($1, $2, $3, $4, $5) RETURNING *`

	updateSongQuery = `UPDATE songs 
						  SET release_date = COALESCE($1, release_date), 
						      text = COALESCE($2, text),
						      link = COALESCE($3, link)
						  WHERE "group" = $4 AND song = $5 
						  RETURNING *
                		 `

	updateSongByIDQuery = `UPDATE songs 
						  SET "group" = COALESCE($1, "group")
						  	  song = COALESCE($2, song)
						      release_date = COALESCE($3, release_date), 
						      text = COALESCE($4, text),
						      link = COALESCE($5, link)
						  WHERE songid=$6
						  RETURNING *
                		 `

	deleteSongQuery = `DELETE FROM songs WHERE "group"=$1 AND song=$2`

	deleteSongByIDQuery = `DELETE FROM songs WHERE songid=$1`

	getLibraryInfoQuery = `SELECT songid, "group", song, release_date, text, link
							FROM songs
							WHERE "group" ILIKE COALESCE('%' || $1 || '%', '%') AND 
							      song ILIKE COALESCE('%' || $2 || '%', '%') AND
								  release_date ILIKE COALESCE('%' || $3 || '%', '%') AND
								  text ILIKE COALESCE('%' || $4 || '%', '%') AND
								  link ILIKE COALESCE('%' || $5 || '%', '%')
							ORDER BY songid
							LIMIT $6 OFFSET $7
					`

	getSongVerseQuery = `SELECT split_part(text, E'\n\n', $1) AS verse, $1 AS verse_id
						FROM songs
						WHERE "group" = $2 AND song = $3;
						`
)
