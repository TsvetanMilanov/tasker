package services

import (
	"github.com/TsvetanMilanov/tasker/src/services/tasks/lib/declarations"

	"github.com/TsvetanMilanov/tasker-common/common/cutils"
	"github.com/TsvetanMilanov/tasker/src/services/tasks/lib/data/models"
	"github.com/TsvetanMilanov/tasker/src/services/tasks/lib/types/requests"
)

// TasksService implements ITasksService.
type TasksService struct {
	DB declarations.ITasksDB `di:""`
}

// Create creates new task item in the database from the
// provided request model for the provided user.
func (s *TasksService) Create(req requests.CreateTask, userID string) (taskID string, err error) {
	id, err := cutils.CreateUUIDString()
	if err != nil {
		return "", err
	}

	task := models.Task{
		TaskID:            id,
		Name:              req.Name,
		DurationInMinutes: req.DurationInMinutes,
		UserID:            userID,
		Notification: models.Notification{
			Time: req.Notification.Time,
		},
	}

	err = s.DB.CreateTask(task)
	if err != nil {
		return "", err
	}

	return id, nil
}
