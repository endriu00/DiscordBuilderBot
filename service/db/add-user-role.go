package db

import ()

func (db *DB) AddUserRole(userID, roleID string) error {
	_, err := db.DB.Exec(`INSERT INTO rank(user_id, role_id) VALUES (?, ?)`, userID, roleID)
	if err != nil {
		return err
	}
	return nil
}
