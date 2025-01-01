package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
  ID int
  Task string
  Completed bool
}

var todos []Todo

func showAllTasks(){
  fmt.Println("All Tasks: ")
  fmt.Println("--------------------------")
  for _, todo := range todos {
    var completionStatus string = "Completed"
    if(!todo.Completed) {
      completionStatus ="In Progress"
    }
    fmt.Printf("ID: %d | Task: %s | Status: %s\n",todo.ID, todo.Task, completionStatus)
  }
  fmt.Println("--------------------------")
}

func addNewTask() {
  fmt.Printf("Enter task description: ")
  reader := bufio.NewReader(os.Stdin)
  task, _ := reader.ReadString('\n')
  task = strings.TrimSpace(task)
  newTodo := Todo{
    ID: len(todos),
    Task: task,
    Completed: false,
  }
  todos = append(todos, newTodo)
  fmt.Printf("Task '%s' added successfully!", task)
}

func markAsCompleted() {
  var id int
  fmt.Println("Enter the Todo ID to mark as completed: ")
  fmt.Scanln(&id)
  todos[id].Completed=true
  fmt.Printf("Todo with ID: %d ('%s') marked as completed!", id, todos[id].Task)
}

func deleteTask(){
  var id int
  fmt.Println("Enter the Todo ID to delete: ")
  fmt.Scanln(&id)
  if id>=len(todos) || id<0 {
    fmt.Println("Invalid ID")
    return
  }
  if id==len(todos)-1 {
    todos=todos[:id]
  } else {
    todos=append(todos[:id], todos[id:]...)
  }
  // todos=slices.Delete(todos,id,id+1)
  fmt.Printf("Todo with ID: %d ('%s') has been deleted successfully!", id, todos[id].Task)
}

func main() { 
  fmt.Println("Welcome to the TO-DO CLI App")
  for {
    var option int
    fmt.Printf("\nPlease choose an option:\n1. View All Tasks\n2. Add a new Task\n3. Mark a task as Completed\n4. Delete a Task\n5. Exit\n")
    fmt.Scanln(&option)
    switch option {
    case 1:
      showAllTasks()
    case 2:
      addNewTask()
    case 3:
      markAsCompleted()
    case 4:
      deleteTask()
    case 5:
      return
    }
  }
}
