package movie

import (
	"context"
	"github.com/charlesoller/omni-api/internal/database"
	"github.com/charlesoller/omni-api/internal/db"
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