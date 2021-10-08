package db

import ()

func (db *DB) GetUserNextRole(userID string) (Role, error) {
	var role Role
	err := db.DB.Get(&role,
		`SELECT role.id, role.name, role.min_points 
		FROM role WHERE role.id NOT IN 
		(SELECT role.id FROM role JOIN rank ON role.id=rank.role_id
		WHERE rank.user_id=$1) ORDER BY role.min_points ASC LIMIT 1`, userID)
	if err != nil {
		return role, err
	}
	return role, nil
}
