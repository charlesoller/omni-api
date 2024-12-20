// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: movie_languages.sql

package db

import (
	"context"
)

const upsertMovieLanguage = `-- name: UpsertMovieLanguage :one
INSERT INTO movie_languages (
  movie_id, language_id
) VALUES (
  $1, $2
)
ON CONFLICT (movie_id, language_id) DO UPDATE SET
  movie_id = EXCLUDED.movie_id,  
  language_id = EXCLUDED.language_id
RETURNING movie_id, language_id
`

type UpsertMovieLanguageParams struct {
	MovieID    int32  `json:"movie_id"`
	LanguageID string `json:"language_id"`
}

func (q *Queries) UpsertMovieLanguage(ctx context.Context, arg UpsertMovieLanguageParams) (MovieLanguage, error) {
	row := q.queryRow(ctx, q.upsertMovieLanguageStmt, upsertMovieLanguage, arg.MovieID, arg.LanguageID)
	var i MovieLanguage
	err := row.Scan(&i.MovieID, &i.LanguageID)
	return i, err
}
