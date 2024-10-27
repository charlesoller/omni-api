// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	GetAllMovies(ctx context.Context, arg GetAllMoviesParams) ([]Movie, error)
	GetMovie(ctx context.Context, id int32) (Movie, error)
	UpdateMovieEmbedding(ctx context.Context, arg UpdateMovieEmbeddingParams) error
	UpsertCastMember(ctx context.Context, arg UpsertCastMemberParams) (CastMember, error)
	UpsertCollection(ctx context.Context, arg UpsertCollectionParams) (Collection, error)
	UpsertCountry(ctx context.Context, arg UpsertCountryParams) (Country, error)
	UpsertCountryISO(ctx context.Context, iso31661 string) (Country, error)
	UpsertCredit(ctx context.Context, id int32) error
	UpsertCreditCastMember(ctx context.Context, arg UpsertCreditCastMemberParams) (CreditsCastMember, error)
	UpsertCreditCrewMember(ctx context.Context, arg UpsertCreditCrewMemberParams) (CreditsCrewMember, error)
	UpsertCrewMember(ctx context.Context, arg UpsertCrewMemberParams) (int32, error)
	UpsertGenre(ctx context.Context, arg UpsertGenreParams) (Genre, error)
	UpsertLanguage(ctx context.Context, arg UpsertLanguageParams) (Language, error)
	UpsertMovie(ctx context.Context, arg UpsertMovieParams) (Movie, error)
	UpsertMovieCountry(ctx context.Context, arg UpsertMovieCountryParams) (MovieCountry, error)
	UpsertMovieGenre(ctx context.Context, arg UpsertMovieGenreParams) (MovieGenre, error)
	UpsertMovieLanguage(ctx context.Context, arg UpsertMovieLanguageParams) (MovieLanguage, error)
	UpsertMovieProductionCompany(ctx context.Context, arg UpsertMovieProductionCompanyParams) (MovieProductionCompany, error)
	UpsertProductionCompany(ctx context.Context, arg UpsertProductionCompanyParams) (ProductionCompany, error)
}

var _ Querier = (*Queries)(nil)
