package db

import (
	pgx "github.com/jackc/pgx/v4/pgxpool"
)

// `Role` represents a role the user has in the server.
type Role struct {
	// ID is the ID of the user.
	ID string `db:"id"`

	// Name is the name of the user.
	Name string `db:"name"`

	// MinPoints is the minimum amount of points for entering the role.
	MinPoints int `db:"min_points"`
}

// `User` represents a user in the server.
type User struct {
	// ID is the ID of the user.
	ID string `db:"id"`

	// Username is the username of the user.
	Username string `db:"username"`

	// Points is the amount of points the user has.
	Points int `db:"points"`

	// Bans is the amount of bans the user has.
	Bans int `db:"ban_alerts"`
}

// `DB` is the database.
type DB struct {
	DB *pgx.Pool
}

// `New` creates a new database from a connection pool.
func New(dbPool *pgx.Pool) (*DB, error) {
	return &DB{
		DB: dbPool,
	}, nil
}
