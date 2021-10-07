package db

import ()

func (db DB) UpdateUserPoints(userID string, points int) error {
	_, err := db.DB.Exec("UPDATE discord_user SET points=points+? WHERE id=?", points, userID)
	if err != nil {
		return err
	}
	return nil
}
