package declarations

import (
	"github.com/TsvetanMilanov/tasker/src/services/tasks/lib/data/models"
	"github.com/TsvetanMilanov/tasker/src/services/tasks/lib/types/config"
	"github.com/TsvetanMilanov/tasker/src/services/tasks/lib/types/requests"
	"github.com/aws/aws-sdk-go/aws"
)

// IDBClient describes methods for working with AWS Dynamo DB.
type IDBClient interface {
	Create(item interface{}, table string, awsCfg *aws.Config) error
}

// ITasksService describes methods for working with tasks.
type ITasksService interface {
	Create(req requests.CreateTask, userID string) (taskID string, err error)
}

// ITasksConfigService describes methods for working with the tasks configuration.
type ITasksConfigService interface {
	GetDynamoDBConfig() config.DynamoDBConfig
}

// ITasksDB describes methods for working with the tasks database.
type ITasksDB interface {
	CreateTask(task models.Task) error
}
