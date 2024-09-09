package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int       `json:"userID"`
}


var events = []Event{}

func (e Event) Save(){
	//later : add it to a databse
	events = append(events, e)
}

func GetAllEvents () []Event{
	return events
}