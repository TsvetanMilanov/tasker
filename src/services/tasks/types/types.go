package types

import "github.com/TsvetanMilanov/tasker/src/services/tasks/declarations"

// CreateHandler ...
type CreateHandler struct {
	Tasks declarations.ITasksService `di:""`
}
