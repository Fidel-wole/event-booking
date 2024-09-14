package models

import (
	"fmt"
	"time"

	"github.com/Fidel-wolee/event-booking/db"
	
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int64      `json:"userID"`
} 
// Save method to insert an event into the database
func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES(?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("could not prepare query: %v", err)
	}
	defer stmt.Close()

	dateTimeStr := e.DateTime.Format("2006-01-02T15:04:05Z07:00")

	// Execute the statement with the event's values
	result, err := stmt.Exec(e.Name, e.Description, e.Location, dateTimeStr, e.UserID)
	if err != nil {
		return fmt.Errorf("could not execute query: %v", err)
	}

	// Get the last inserted ID and assign it to the event
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("could not retrieve last insert ID: %v", err)
	}
	e.ID = id

	return nil
}

// GetAllEvents returns all events stored in the database
func GetAllEvents() ([]Event, error) {
	rows, err := db.DB.Query("SELECT id, name, description, location, dateTime, user_id FROM events")
	if err != nil {
		return nil, fmt.Errorf("could not query events: %v", err)
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		var dateTimeStr string
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserID); err != nil {
			return nil, fmt.Errorf("could not scan event row: %v", err)
		}
		// Parse the dateTime string back to time.Time
		event.DateTime, err = time.Parse("2006-01-02T15:04:05Z07:00", dateTimeStr)
		if err != nil {
			return nil, fmt.Errorf("could not parse dateTime: %v", err)
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e *Event) Update() error {
	query := `
UPDATE events
SET name = ?, description =?, location =?, dateTime =?
WHERE id = ?
`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("could not prepare query: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e Event) Delete() error{
	query := "DELETE FROM events WHERE id=?"
	stmt, err := db.DB.Prepare(query)

	if err != nil{
		return err
	}

	defer stmt.Close()
    _, err = stmt.Exec(e.ID)
	return err
   
}