package db

import ()

func (db *DB) GetUserNextRole(userID string) (Role, error) {
	var role Role
	rows, err := db.DB.Query(
		`SELECT role.id, role.name, role.min_points 
	FROM role JOIN rank ON role.id=rank.role_id
	WHERE role.id NOT IN rank.role_id 
	AND user_id=?
	ORDER BY role.min_points ASC LIMIT 1`, userID)
	if err != nil {
		return role, err
	}
	if err = rows.Scan(&role.ID, role.Name, role.MinPoints); err != nil {
		return role, err
	}
	return role, nil
}
