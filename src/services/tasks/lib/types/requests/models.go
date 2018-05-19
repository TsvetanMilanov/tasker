package requests

import "time"

// CreateTask request model.
type CreateTask struct {
	Name              string           `json:"name" validate:"required"`
	DurationInMinutes uint64           `json:"durationInMinutes" validate:"required,gt=0"`
	Notification      TaskNotification `json:"notification" validate:"dive"`
}

// TaskNotification request model.
type TaskNotification struct {
	Time time.Time `json:"time" validate:"required"`
}
