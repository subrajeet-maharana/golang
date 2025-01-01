package main

import (
	"bufio"
	"fmt"
	"os"
)

type Todo struct {
  ID int
  Task string
  Completed bool
}

var todos []Todo

func showAllTasks(){
  fmt.Println("All Tasks")
  for _, todo := range todos {
    var completionStatus string = "Completed"
    if(!todo.Completed) {
      completionStatus ="In Progress"
    }
    fmt.Printf("ID: %d | Task: %s | Status: %s\n",todo.ID, todo.Task, completionStatus)
  }
}

func addNewTask() {
  fmt.Printf("Add task: ")
  reader:=bufio.NewReader(os.Stdin)
  task, _ :=reader.ReadString('\n')
  newTodo := Todo{
    ID: len(todos),
    Task: task,
    Completed: false,
  }
  todos = append(todos, newTodo)
  fmt.Println("Todo added successfully...")
}

func markAsCompleted() {
  var id int
  fmt.Println("Enter the ID of the task: ")
  fmt.Scanln(&id)
  todos[id].Completed=true
  fmt.Printf("Todo with ID: %d marked as completed..", id)
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
    }
  }
}
