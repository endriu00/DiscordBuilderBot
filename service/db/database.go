package db

import (
	"github.com/jmoiron/sqlx"
)

// `Role` represents a role in the server.
type Role struct {
	ID        string
	Name      string
	MinPoints int
}

// `User` represents a user in the server.
type User struct {
	ID       string
	Username string
	Points   int
	Bans     int
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
