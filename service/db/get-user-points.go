package db

import ()

func (db *DB) GetUserPoints(userID string) (int, error) {
	var points int
	rows := db.DB.QueryRow("SELECT points FROM user WHERE id=?", userID)
	err := rows.Scan(&points)
	if err != nil {
		return -1, err
	}
	return points, nil
}
