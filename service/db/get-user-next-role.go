package db

import (
	"context"
)

// `GetUserNextRole` returns the next role the user will reach in the server.
// A context `ctx` is used for the database connection.
func (db *DB) GetUserNextRole(userID string, ctx context.Context) (Role, error) {
	var role Role
	row := db.DB.QueryRow(ctx,
		`SELECT role.id, role.name, role.min_points 
		FROM role WHERE role.id NOT IN 
		(SELECT role.id FROM role JOIN rank ON role.id=rank.role_id
		WHERE rank.user_id=$1) ORDER BY role.min_points ASC LIMIT 1`, userID)
	err := row.Scan(&role.ID, &role.Name, &role.MinPoints)
	if err != nil {
		return role, err
	}
	return role, nil
}
