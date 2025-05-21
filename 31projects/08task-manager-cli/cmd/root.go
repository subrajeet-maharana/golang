package cmd

import (
	"flag"

	"github.com/subrajeet-maharana/golang/31services/08task-manager-cli/storage"
	"github.com/subrajeet-maharana/golang/31services/08task-manager-cli/tasks"
	"github.com/subrajeet-maharana/golang/31services/08task-manager-cli/util"
)

func Execute() {
	var title = flag.String("add", "Go to Gym", "Enter the Title for the task")
	var deleteId = flag.Int64("delete", 0, "Enter the ID of the task to delete: ")
	var markId = flag.Int64("mark", 0, "Enter the Task ID to Mark: ")

	flag.Parse()

	switch {
	case *title != "":
		var newTask tasks.Task = util.CreateTask(*title)
		storage.MasterTasks.AddTask(newTask)
		storage.MasterTasks.ListTasks()
	case *deleteId >= 0:
		storage.MasterTasks.DeleteTask(int(*deleteId))
		storage.MasterTasks.ListTasks()
	case *markId >= 0:
		storage.MasterTasks.MarkAsDone(int(*markId))
		storage.MasterTasks.ListTasks()
	default:
		storage.MasterTasks.ListTasks()
	}
}
