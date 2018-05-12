package data

import (
	"github.com/TsvetanMilanov/tasker/src/services/tasks/data/models"
	"github.com/TsvetanMilanov/tasker/src/services/tasks/declarations"
	"github.com/aws/aws-sdk-go/aws"
)

// TasksDB implements ITasksDB.
type TasksDB struct {
	DB          declarations.IDBClient           `di:""`
	TasksConfig declarations.ITasksConfigService `di:""`
}

// CreateTask saves the new task in the database.
func (db *TasksDB) CreateTask(task models.Task) error {
	cfg := db.TasksConfig.GetDynamoDBConfig()
	awsCfg := aws.NewConfig().WithRegion(cfg.Region)
	return db.DB.Create(task, cfg.TasksTableName, awsCfg)
}
