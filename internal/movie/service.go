package movie

import (
	"context"

	"github.com/charlesoller/omni-api/internal/db"
)

type MovieService struct {
	r *MovieRepository
}

func NewMovieService(repo *MovieRepository) *MovieService {
	return &MovieService{
		r: repo,
	}
}

func (s *MovieService) GetAllMovies(ctx context.Context, page int32, limit int32) ([]db.Movie, error) {
	offset := (page - 1) * limit

	movies, err := s.r.GetAllMovies(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (s *MovieService) GetMovie(ctx context.Context, id int32) (*db.Movie, error) {
	movie, err := s.r.GetMovie(ctx, id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}