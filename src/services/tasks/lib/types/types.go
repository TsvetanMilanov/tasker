package types

import "github.com/TsvetanMilanov/tasker/src/services/tasks/lib/declarations"

// CreateHandler ...
type CreateHandler struct {
	Tasks declarations.ITasksService `di:""`
}
