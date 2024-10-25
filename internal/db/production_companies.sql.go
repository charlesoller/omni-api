// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: production_companies.sql

package db

import (
	"context"
)

const upsertProductionCompany = `-- name: UpsertProductionCompany :one
INSERT INTO production_companies (
  id, name, logo_path, origin_country
) VALUES (
  $1, $2, $3, $4
)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  logo_path = EXCLUDED.logo_path,
  origin_country = EXCLUDED.origin_country
RETURNING id, name, logo_path, origin_country
`

type UpsertProductionCompanyParams struct {
	ID            int32  `json:"id"`
	Name          string `json:"name"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

func (q *Queries) UpsertProductionCompany(ctx context.Context, arg UpsertProductionCompanyParams) (ProductionCompany, error) {
	row := q.queryRow(ctx, q.upsertProductionCompanyStmt, upsertProductionCompany,
		arg.ID,
		arg.Name,
		arg.LogoPath,
		arg.OriginCountry,
	)
	var i ProductionCompany
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LogoPath,
		&i.OriginCountry,
	)
	return i, err
}