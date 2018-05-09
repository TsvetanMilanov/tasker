package types

import "github.com/TsvetanMilanov/tasker/src/services/tasks/declarations"

// CreateHandler ...
type CreateHandler struct {
	DB declarations.IDB `di:""`
}
