package db

import ()

func (db *DB) AddUser(userID, username string) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO user(id, username) VALUES (?, ?)", userID, username)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
