package db

import (
	"context"
)

// `AddUser` adds the user with ID `userID` and username `username` to the database.
// A context `ctx` is used for the database connection.
func (db *DB) AddUser(userID, username string, ctx context.Context) error {
	_, err := db.DB.Exec(ctx, "INSERT INTO discord_user(id, username, points, ban_alerts) VALUES ($1, $2, 0, 0)", userID, username)
	if err != nil {
		return err
	}
	return nil
}
