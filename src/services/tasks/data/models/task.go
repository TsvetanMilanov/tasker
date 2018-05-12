package models

import "time"

// Task represents task in the database.
type Task struct {
	TaskID            string       `validate:"required"`
	UserID            string       `validate:"required"`
	Name              string       `validate:"required"`
	DurationInMinutes uint64       `validate:"required"`
	Notification      Notification `validate:"dive,required"`
}

// Notification represents task notification in the database.
type Notification struct {
	Time time.Time `validate:"required"`
}
