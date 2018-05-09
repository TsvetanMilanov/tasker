package declarations

// IDB describes methods for working with the tasks db.
type IDB interface {
	CreateTask(name, user string) (taskID string, err error)
}
