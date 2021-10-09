package db

import (
	"context"
)

// `UpdateUserPoints` updates the user points for the user with ID `userID`
// by an amount of `points` points.
// A context `ctx` is used for the database connection.
func (db *DB) UpdateUserPoints(userID string, points int, ctx context.Context) error {
	_, err := db.DB.Exec(ctx, "UPDATE discord_user SET points=points+$1 WHERE id=$2", points, userID)
	if err != nil {
		return err
	}
	return nil
}
