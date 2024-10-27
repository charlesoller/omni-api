package models

import (
	"time"

	"github.com/charlesoller/omni-api/internal/db"
)

type MovieDetails struct {
	ID                   int32                  `json:"id"`
	Title                string                 `json:"title"`
	Overview             string                 `json:"overview"`
	BackdropPath         string                 `json:"backdrop_path"`
	Budget               int64                  `json:"budget"`
	Popularity           float64                `json:"popularity"`
	PosterPath           string                 `json:"poster_path"`
	ReleaseDate          time.Time              `json:"release_date"`
	Revenue              int64                  `json:"revenue"`
	Runtime              int32                  `json:"runtime"`
	VoteAverage          float64                `json:"vote_average"`
	VoteCount            int32                  `json:"vote_count"`
	Status               string                 `json:"status"`
	CollectionName       string                 `json:"collection_name"`
	CollectionPosterPath string                 `json:"collection_poster_path"`
	CastMembers          []db.CastMember        `json:"cast_members"`
	CrewMembers          []db.CrewMember        `json:"crew_members"`
	Genres               []db.Genre             `json:"genres"`
	Countries            []db.Country           `json:"countries"`
	Languages            []db.Language          `json:"languages"`
	ProductionCompanies  []db.ProductionCompany `json:"production_companies"`
}

