package util

import (
	"time"

	"github.com/subrajeet-maharana/golang/31services/08task-manager-cli/storage"
	"github.com/subrajeet-maharana/golang/31services/08task-manager-cli/tasks"
)

func CreateTask(title string) tasks.Task {
	newTask := tasks.Task{
		ID:      storage.MasterTasks.TaskInd,
		Title:   title,
		Done:    false,
		Created: time.Now(),
	}
	return newTask
}
