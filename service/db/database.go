package db

import (
	"github.com/jmoiron/sqlx"
)

// `Role` represents a role in the server.
type Role struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	MinPoints int    `db:"min_points"`
}

// `User` represents a user in the server.
type User struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	Points   int    `db:"points"`
	Bans     int    `db:"ban_alerts"`
}

// `DB` is the database.
type DB struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) (*DB, error) {
	return &DB{
		DB: db,
	}, nil
}
