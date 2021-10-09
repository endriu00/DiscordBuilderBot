package db

import (
	"context"
)

// `GetUserPoints` returns the points of the user with ID `roleID`.
// A context `ctx` is used for the database connection.
func (db *DB) GetUserPoints(userID string, ctx context.Context) (int, error) {
	var points int
	rows := db.DB.QueryRow(ctx, "SELECT points FROM discord_user WHERE id=$1", userID)
	err := rows.Scan(&points)
	if err != nil {
		return -1, err
	}
	return points, nil
}
