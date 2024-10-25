package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/charlesoller/omni-api/internal/db"
)

type Store struct {
	Q  *db.Queries
	db *sql.DB
}

func NewStore(db *sql.DB, queries *db.Queries) *Store {
	return &Store{
		Q:  queries,
		db: db,
	}
}

func (s *Store) ExecTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
