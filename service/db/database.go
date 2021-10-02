package db

import (
	"github.com/jmoiron/sqlx"
)

type Stats struct {
	Points    int
	BanAlerts int
}

type User struct {
	ID       string
	Username string
}

type DB struct {
	DB *sqlx.DB
}
