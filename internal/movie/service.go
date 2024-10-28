package movie

import (
	"context"
	"encoding/json"

	"github.com/charlesoller/omni-api/internal/db"
	"github.com/charlesoller/omni-api/internal/models"
	"github.com/pgvector/pgvector-go"
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

func (s *MovieService) GetAllMoviesMatchingSearch(ctx context.Context, page int32, limit int32, search string) ([]db.Movie, error) {
	offset := (page - 1) * limit

	movies, err := s.r.GetAllMoviesMatchingSearch(ctx, offset, limit, search)
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

func (s *MovieService) GetRandomMovie(ctx context.Context) (*db.Movie, error) {
	movie, err := s.r.GetRandomMovie(ctx)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (s *MovieService) GetMovieDetails(ctx context.Context, id int32) (*models.MovieDetails, error) {
	movie, err := s.r.GetMovieDetails(ctx, id)
	if err != nil {
		return nil, err
	}

	var castMembers []db.CastMember
	if err := json.Unmarshal(movie.CastMembers, &castMembers); err != nil {
		return nil, err
	}

	var crewMembers []db.CrewMember
	if err := json.Unmarshal(movie.CrewMembers, &crewMembers); err != nil {
		return nil, err
	}

	var genres []db.Genre
	if err := json.Unmarshal(movie.Genres, &genres); err != nil {
		return nil, err
	}
	
	var countries []db.Country
	if err := json.Unmarshal(movie.Countries, &countries); err != nil {
		return nil, err
	}

	var languages []db.Language
	if err := json.Unmarshal(movie.Languages, &languages); err != nil {
		return nil, err
	}

	var productionCompanies []db.ProductionCompany
	if err := json.Unmarshal(movie.ProductionCompanies, &productionCompanies); err != nil {
		return nil, err
	}

	return &models.MovieDetails{
		ID: movie.ID,
		Title: movie.Title,
		Overview: movie.Overview,
		BackdropPath: movie.BackdropPath,
		Budget: movie.Budget,
		Popularity: movie.Popularity,
		PosterPath: movie.PosterPath,
		ReleaseDate: movie.ReleaseDate,
		Revenue: movie.Revenue,
		Runtime: movie.Runtime,
		VoteAverage: movie.VoteAverage,
		VoteCount: movie.VoteCount,
		Status: movie.Status,
		CollectionName: movie.CollectionName,
		CollectionPosterPath: movie.CollectionPosterPath,
		CastMembers: castMembers,
		CrewMembers: crewMembers,
		Genres: genres,
		Countries: countries,
		Languages: languages,
		ProductionCompanies: productionCompanies,
	}, nil
}

func (s *MovieService) GetSimilarMovies(ctx context.Context, embedding *pgvector.Vector) ([]db.Movie, error) {
	movie, err := s.r.GetSimilarMovies(ctx, embedding)
	if err != nil {
		return nil, err
	}
	return movie, nil
}