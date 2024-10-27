package movie

import (
	"context"

	"github.com/charlesoller/omni-api/internal/database"
	"github.com/charlesoller/omni-api/internal/db"
	"github.com/pgvector/pgvector-go"
)

type MovieRepository struct {
	db *database.Store
}

func NewMovieRepository(db *database.Store) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) GetAllMovies(ctx context.Context, offset int32, limit int32) ([]db.Movie, error) {
	p := db.GetAllMoviesParams {
		Limit: limit,
		Offset: offset,
	}
	
	movies, err := r.db.Q.GetAllMovies(ctx, p)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *MovieRepository) GetMovie (ctx context.Context, id int32) (*db.Movie, error) {
	movie, err := r.db.Q.GetMovie(ctx, id)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *MovieRepository) GetMovieDetails (ctx context.Context, id int32) (*db.GetMovieDetailsRow, error) {
	movie, err := r.db.Q.GetMovieDetails(ctx, id)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *MovieRepository) GetSimilarMovies (ctx context.Context, embedding *pgvector.Vector) ([]db.Movie, error) {
	movies, err := r.db.Q.FindSimilarMovies(ctx, *embedding)
	if err != nil {
		return nil, err
	}
	return movies, nil
}