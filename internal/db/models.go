// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"time"

	pgvector_go "github.com/pgvector/pgvector-go"
)

type CastMember struct {
	ID                 int32   `json:"id"`
	CastID             int32   `json:"cast_id"`
	CreditID           string  `json:"credit_id"`
	Gender             int16   `json:"gender"`
	Adult              bool    `json:"adult"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
}

type Collection struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
}

type Country struct {
	Iso31661 string `json:"iso_3166_1"`
	Name     string `json:"name"`
}

type Credit struct {
	ID int32 `json:"id"`
}

type CreditsCastMember struct {
	CreditID  int32  `json:"credit_id"`
	MemberID  int32  `json:"member_id"`
	Character string `json:"character"`
	Order     int32  `json:"order"`
}

type CreditsCrewMember struct {
	CreditID   int32  `json:"credit_id"`
	MemberID   int32  `json:"member_id"`
	Department string `json:"department"`
	Job        string `json:"job"`
}

type CrewMember struct {
	ID                 int32   `json:"id"`
	CreditID           string  `json:"credit_id"`
	Gender             int16   `json:"gender"`
	Adult              bool    `json:"adult"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
}

type Genre struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Language struct {
	EnglishName string `json:"english_name"`
	Iso6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
}

type Movie struct {
	ID               int32              `json:"id"`
	Title            string             `json:"title"`
	OriginalTitle    string             `json:"original_title"`
	Overview         string             `json:"overview"`
	ReleaseDate      time.Time          `json:"release_date"`
	Runtime          int32              `json:"runtime"`
	Budget           int64              `json:"budget"`
	Revenue          int64              `json:"revenue"`
	Popularity       float64            `json:"popularity"`
	VoteAverage      float64            `json:"vote_average"`
	VoteCount        int32              `json:"vote_count"`
	Status           string             `json:"status"`
	Tagline          string             `json:"tagline"`
	Homepage         string             `json:"homepage"`
	OriginalLanguage string             `json:"original_language"`
	Adult            bool               `json:"adult"`
	BackdropPath     string             `json:"backdrop_path"`
	PosterPath       string             `json:"poster_path"`
	CollectionID     int32              `json:"collection_id"`
	Embedding        pgvector_go.Vector `json:"embedding"`
}

type MovieCountry struct {
	MovieID   int32  `json:"movie_id"`
	CountryID string `json:"country_id"`
}

type MovieGenre struct {
	MovieID int32 `json:"movie_id"`
	GenreID int32 `json:"genre_id"`
}

type MovieLanguage struct {
	MovieID    int32  `json:"movie_id"`
	LanguageID string `json:"language_id"`
}

type MovieProductionCompany struct {
	MovieID   int32 `json:"movie_id"`
	CompanyID int32 `json:"company_id"`
}

type ProductionCompany struct {
	ID            int32  `json:"id"`
	Name          string `json:"name"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}
