package models

import "go-course/api/db"

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Create the save method that should return error
func (u User) Save() error {
	/// write the query
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	/// prepare the statement
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	/// defer statement close
	defer stmt.Close()
	/// execute the statement

	result, err := stmt.Exec(query, u.Email, u.Password)
	if err != nil {
		return err
	}
	/// get the last insert id
	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	/// assign the id to the user id
	u.ID = lastID

	/// return nil
	return nil
}

/// teste
