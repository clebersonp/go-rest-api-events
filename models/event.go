package models

import "time"

// Event - struct tags for json deserialized string into structs and
// binding feature from gin package to required missing fields
type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"user_id"`
}

var events []Event

func (e *Event) Save() {
	// later: add it to a database
	events = append(events, *e)
}

func GetAllEvents() []Event {
	if events == nil {
		return []Event{}
	}
	return events
}
