package models

import (
	"database/sql"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

func (u *User) Save(db *sql.DB) error {
	insertQuery := "INSERT INTO users (name, age) VALUES (?, ?)"
	result, err := db.Exec(insertQuery, u.Name, u.Age)
	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (u *User) UpdateUser(db *sql.DB) error {
	updateQuery := "UPDATE users SET name = ?, age = ? WHERE id = ?"
	_, err := db.Exec(updateQuery, u.Name, u.Age, u.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByID(db *sql.DB, userID int64) (*User, error) {
	query := "SELECT id, name, age FROM users WHERE id = ?"
	row := db.QueryRow(query, userID)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUserByID(db *sql.DB, userID int64) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, userID)
	if err != nil {
		return err
	}

	return nil
}
