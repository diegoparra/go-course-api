package models

import (
	"fmt"
	"time"

	"go-course/api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"        binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location"    binding:"required"`
	DateTime    time.Time `json:"date_time"   binding:"required"`
	UserId      int       `json:"user_id"`
}

var events = []Event{}

func (e Event) Save() error {
	query := `INSERT INTO events (name, description, location, dateTime, user_id) VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id

	events = append(events, e)

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)
	if err != nil {
		return []Event{}, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserId,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events where id = ?`

	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserId,
	)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func UpdateEvent(e Event, id int64) error {
	query := `UPDATE events set name = ?, description = ?, location = ?, dateTime = ? where id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func DeleteEvent(id int64) error {
	e, err := GetEventByID(id)
	if err != nil {
		return err
	}

	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	if err != nil {
		return err
	}

	return nil
}
