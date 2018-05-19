package services

import (
	"os"

	"github.com/TsvetanMilanov/tasker/src/services/tasks/lib/types/config"
)

// TasksConfigService implements ITasksConfigService
type TasksConfigService struct {
}

// GetDynamoDBConfig returns the AWS Dynamo DB config gathered from env vars.
func (s *TasksConfigService) GetDynamoDBConfig() config.DynamoDBConfig {
	return config.DynamoDBConfig{
		Region:         os.Getenv("AWS_REGION"),
		TasksTableName: os.Getenv("TASKS_TABLE_NAME"),
	}
}
