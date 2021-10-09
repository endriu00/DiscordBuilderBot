package db

import (
	"context"
)

// `AddUserRole` adds the role with ID `roleID` to the user `userID`.
// A context `ctx` is used for the database connection.
func (db *DB) AddUserRole(userID, roleID string, ctx context.Context) error {
	_, err := db.DB.Exec(ctx, `INSERT INTO rank(user_id, role_id) VALUES ($1, $2)`, userID, roleID)
	if err != nil {
		return err
	}
	return nil
}
