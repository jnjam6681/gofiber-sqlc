package postgres

import "database/sql"

type Repo struct {
	*Queries
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db:      db,
		Queries: New(db),
	}
}
