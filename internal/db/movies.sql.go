// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: movies.sql

package db

import (
	"context"
	"time"

	pgvector_go "github.com/pgvector/pgvector-go"
)

const findSimilarMovies = `-- name: FindSimilarMovies :many
SELECT id, title, original_title, overview, release_date, runtime, budget, revenue, popularity, vote_average, vote_count, status, tagline, homepage, original_language, adult, backdrop_path, poster_path, collection_id, embedding
FROM movies
ORDER BY embedding <=> $1
LIMIT 5
`

func (q *Queries) FindSimilarMovies(ctx context.Context, embedding pgvector_go.Vector) ([]Movie, error) {
	rows, err := q.query(ctx, q.findSimilarMoviesStmt, findSimilarMovies, embedding)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.OriginalTitle,
			&i.Overview,
			&i.ReleaseDate,
			&i.Runtime,
			&i.Budget,
			&i.Revenue,
			&i.Popularity,
			&i.VoteAverage,
			&i.VoteCount,
			&i.Status,
			&i.Tagline,
			&i.Homepage,
			&i.OriginalLanguage,
			&i.Adult,
			&i.BackdropPath,
			&i.PosterPath,
			&i.CollectionID,
			&i.Embedding,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllMovies = `-- name: GetAllMovies :many
SELECT id, title, original_title, overview, release_date, runtime, budget, revenue, popularity, vote_average, vote_count, status, tagline, homepage, original_language, adult, backdrop_path, poster_path, collection_id, embedding FROM movies
ORDER BY id
LIMIT $1 OFFSET $2
`

type GetAllMoviesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetAllMovies(ctx context.Context, arg GetAllMoviesParams) ([]Movie, error) {
	rows, err := q.query(ctx, q.getAllMoviesStmt, getAllMovies, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.OriginalTitle,
			&i.Overview,
			&i.ReleaseDate,
			&i.Runtime,
			&i.Budget,
			&i.Revenue,
			&i.Popularity,
			&i.VoteAverage,
			&i.VoteCount,
			&i.Status,
			&i.Tagline,
			&i.Homepage,
			&i.OriginalLanguage,
			&i.Adult,
			&i.BackdropPath,
			&i.PosterPath,
			&i.CollectionID,
			&i.Embedding,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMovie = `-- name: GetMovie :one
SELECT id, title, original_title, overview, release_date, runtime, budget, revenue, popularity, vote_average, vote_count, status, tagline, homepage, original_language, adult, backdrop_path, poster_path, collection_id, embedding FROM movies
WHERE id = $1
`

func (q *Queries) GetMovie(ctx context.Context, id int32) (Movie, error) {
	row := q.queryRow(ctx, q.getMovieStmt, getMovie, id)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.OriginalTitle,
		&i.Overview,
		&i.ReleaseDate,
		&i.Runtime,
		&i.Budget,
		&i.Revenue,
		&i.Popularity,
		&i.VoteAverage,
		&i.VoteCount,
		&i.Status,
		&i.Tagline,
		&i.Homepage,
		&i.OriginalLanguage,
		&i.Adult,
		&i.BackdropPath,
		&i.PosterPath,
		&i.CollectionID,
		&i.Embedding,
	)
	return i, err
}

const updateMovieEmbedding = `-- name: UpdateMovieEmbedding :exec
UPDATE movies 
SET embedding = $2
WHERE id = $1
`

type UpdateMovieEmbeddingParams struct {
	ID        int32              `json:"id"`
	Embedding pgvector_go.Vector `json:"embedding"`
}

func (q *Queries) UpdateMovieEmbedding(ctx context.Context, arg UpdateMovieEmbeddingParams) error {
	_, err := q.exec(ctx, q.updateMovieEmbeddingStmt, updateMovieEmbedding, arg.ID, arg.Embedding)
	return err
}

const upsertMovie = `-- name: UpsertMovie :one
INSERT INTO movies (
  id, title, original_title, overview, release_date, runtime, budget, revenue,
  popularity, vote_average, vote_count, status, tagline, homepage,
  original_language, adult, backdrop_path, poster_path, collection_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8,
  $9, $10, $11, $12, $13, $14,
  $15, $16, $17, $18, $19
)
ON CONFLICT (id) DO UPDATE SET
  title = EXCLUDED.title,
  original_title = EXCLUDED.original_title,
  overview = EXCLUDED.overview,
  release_date = EXCLUDED.release_date,
  runtime = EXCLUDED.runtime,
  budget = EXCLUDED.budget,
  revenue = EXCLUDED.revenue,
  popularity = EXCLUDED.popularity,
  vote_average = EXCLUDED.vote_average,
  vote_count = EXCLUDED.vote_count,
  status = EXCLUDED.status,
  tagline = EXCLUDED.tagline,
  homepage = EXCLUDED.homepage,
  original_language = EXCLUDED.original_language,
  adult = EXCLUDED.adult,
  backdrop_path = EXCLUDED.backdrop_path,
  poster_path = EXCLUDED.poster_path,
  collection_id = EXCLUDED.collection_id
RETURNING id, title, original_title, overview, release_date, runtime, budget, revenue, popularity, vote_average, vote_count, status, tagline, homepage, original_language, adult, backdrop_path, poster_path, collection_id, embedding
`

type UpsertMovieParams struct {
	ID               int32     `json:"id"`
	Title            string    `json:"title"`
	OriginalTitle    string    `json:"original_title"`
	Overview         string    `json:"overview"`
	ReleaseDate      time.Time `json:"release_date"`
	Runtime          int32     `json:"runtime"`
	Budget           int64     `json:"budget"`
	Revenue          int64     `json:"revenue"`
	Popularity       float64   `json:"popularity"`
	VoteAverage      float64   `json:"vote_average"`
	VoteCount        int32     `json:"vote_count"`
	Status           string    `json:"status"`
	Tagline          string    `json:"tagline"`
	Homepage         string    `json:"homepage"`
	OriginalLanguage string    `json:"original_language"`
	Adult            bool      `json:"adult"`
	BackdropPath     string    `json:"backdrop_path"`
	PosterPath       string    `json:"poster_path"`
	CollectionID     int32     `json:"collection_id"`
}

func (q *Queries) UpsertMovie(ctx context.Context, arg UpsertMovieParams) (Movie, error) {
	row := q.queryRow(ctx, q.upsertMovieStmt, upsertMovie,
		arg.ID,
		arg.Title,
		arg.OriginalTitle,
		arg.Overview,
		arg.ReleaseDate,
		arg.Runtime,
		arg.Budget,
		arg.Revenue,
		arg.Popularity,
		arg.VoteAverage,
		arg.VoteCount,
		arg.Status,
		arg.Tagline,
		arg.Homepage,
		arg.OriginalLanguage,
		arg.Adult,
		arg.BackdropPath,
		arg.PosterPath,
		arg.CollectionID,
	)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.OriginalTitle,
		&i.Overview,
		&i.ReleaseDate,
		&i.Runtime,
		&i.Budget,
		&i.Revenue,
		&i.Popularity,
		&i.VoteAverage,
		&i.VoteCount,
		&i.Status,
		&i.Tagline,
		&i.Homepage,
		&i.OriginalLanguage,
		&i.Adult,
		&i.BackdropPath,
		&i.PosterPath,
		&i.CollectionID,
		&i.Embedding,
	)
	return i, err
}
