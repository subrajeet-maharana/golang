package tasks

import (
	"fmt"
	"slices"
	"time"
)

type Task struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Done    bool      `json:"done"`
	Created time.Time `json:"created"`
}

type TaskList struct {
	Tasks   []Task
	TaskInd int
}

func (t *TaskList) AddTask(task Task) {
	t.Tasks = append(t.Tasks, task)
	t.TaskInd++
}

func (t *TaskList) DeleteTask(id int) {
	for i := range t.Tasks {
		if t.Tasks[i].ID == id {
			t.Tasks = slices.Delete(t.Tasks, i, i+1)
			return
		}
	}
}

func (t *TaskList) MarkAsDone(id int) {
	for i := range t.Tasks {
		if t.Tasks[i].ID == id {
			t.Tasks[i].Done = true
			return
		}
	}
}

func (t *TaskList) ListTasks() {
	fmt.Println("\n===== LIST ALL THE TASKS CURRENTLY IN USE =====")
	for _, task := range t.Tasks {
		fmt.Printf("ID: %d\nTitle: %s\nDone: %t\nCreated: %s\n\n", task.ID, task.Title, task.Done, task.Created)
	}
}
