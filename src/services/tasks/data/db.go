package data

import "github.com/google/uuid"

// DB wraps the database service.
type DB struct {
}

// CreateTask saves the new task in the database and returns its ID.
func (db *DB) CreateTask(name, user string) (taskID string, err error) {
	guid, err := uuid.NewUUID()
	if err != nil {
		return taskID, err
	}

	taskID = guid.String()
	return taskID, err
}
