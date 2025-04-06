package models

import (
	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string		`binding:"required"`
	Location    string		`binding:"required"`
	Description string		`binding:"required"`
	DateTime   	string 		`binding:"required"`
	UserID      int	
}

// var events []Event = []Event{}

// Save method to save an event
// This method appends the event to the events slice
func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id) 
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id

	return nil
}


func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, datetime, user_id FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}



// Note:
// Query is used when you want to get data from the database
// Exec is used when you want to execute a statement that has variables