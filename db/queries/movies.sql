-- name: UpsertMovie :one
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
RETURNING *;

-- name: UpdateMovieEmbedding :exec
UPDATE movies 
SET embedding = $2
WHERE id = $1;

-- name: GetAllMovies :many
SELECT * FROM movies
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: GetMoviesByName :many
SELECT *
FROM movies
WHERE title ILIKE '%' || $1 || '%'
ORDER BY popularity DESC
LIMIT $2 OFFSET $3;

-- name: GetMovie :one
SELECT * FROM movies
WHERE id = $1;

-- name: GetRandomMovie :one
SELECT *
FROM movies
ORDER BY RANDOM()
LIMIT 1;

-- name: FindSimilarMovies :many
SELECT * FROM movies
ORDER BY embedding <=> $1
LIMIT 30; 

-- name: GetMovieDetails :one
WITH cast_members_agg AS (
  SELECT 
    credits_cast_member.credit_id,
    json_agg(
      json_build_object(
        'id', cast_members.id,
        'name', cast_members.name,
        'profile_path', cast_members.profile_path
      )
    ) AS cast_members
  FROM credits_cast_member
  JOIN cast_members ON credits_cast_member.member_id = cast_members.id
  GROUP BY credits_cast_member.credit_id
),
crew_members_agg AS (
  SELECT 
    credits_crew_member.credit_id,
    json_agg(
      json_build_object(
        'id', crew_members.id,
        'name', crew_members.name,
        'profile_path', crew_members.profile_path
      )
    ) AS crew_members
  FROM credits_crew_member
  JOIN crew_members ON credits_crew_member.member_id = crew_members.id
  GROUP BY credits_crew_member.credit_id
),
genres_agg AS (
  SELECT 
    movie_id,
    json_agg(
      json_build_object(
        'id', genres.id,
        'name', genres.name
      )
    ) AS genres
  FROM movie_genres
  JOIN genres ON movie_genres.genre_id = genres.id
  GROUP BY movie_id
),
countries_agg AS (
  SELECT 
    movie_id,
    json_agg(
      json_build_object(
        'name', countries.name
      )
    ) AS countries
  FROM movie_countries
  JOIN countries ON movie_countries.country_id = countries.iso_3166_1
  GROUP BY movie_id
),
languages_agg AS (
  SELECT 
    movie_id,
    json_agg(
      json_build_object(
        'name', languages.name,
        'english_name', languages.english_name
      )
    ) AS languages
  FROM movie_languages
  JOIN languages ON movie_languages.language_id = languages.iso_639_1
  GROUP BY movie_id
),
production_companies_agg AS (
  SELECT 
    movie_id,
    json_agg(
      json_build_object(
        'name', production_companies.name,
        'logo_path', production_companies.logo_path
      )
    ) AS production_companies
  FROM movie_production_companies
  JOIN production_companies ON movie_production_companies.company_id = production_companies.id
  GROUP BY movie_id
)
SELECT 
  movies.id,
  movies.title,
  movies.overview,
  movies.backdrop_path,
  movies.budget,
  movies.popularity,
  movies.poster_path,
  movies.release_date,
  movies.revenue,
  movies.runtime,
  movies.vote_average,
  movies.vote_count,
  movies.status,
  COALESCE(collections.name, '') AS collection_name,
  COALESCE(collections.poster_path, '') AS collection_poster_path,
  COALESCE(cast_members_agg.cast_members, '[]'::json) AS cast_members,
  COALESCE(crew_members_agg.crew_members, '[]'::json) AS crew_members,
  COALESCE(genres_agg.genres, '[]'::json) AS genres,
  COALESCE(countries_agg.countries, '[]'::json) AS countries,
  COALESCE(languages_agg.languages, '[]'::json) AS languages,
  COALESCE(production_companies_agg.production_companies, '[]'::json) AS production_companies
FROM 
  movies
LEFT JOIN collections ON movies.collection_id = collections.id
LEFT JOIN cast_members_agg ON cast_members_agg.credit_id = movies.id
LEFT JOIN crew_members_agg ON crew_members_agg.credit_id = movies.id
LEFT JOIN genres_agg ON genres_agg.movie_id = movies.id
LEFT JOIN countries_agg ON countries_agg.movie_id = movies.id
LEFT JOIN languages_agg ON languages_agg.movie_id = movies.id
LEFT JOIN production_companies_agg ON production_companies_agg.movie_id = movies.id
WHERE
  movies.id = $1;