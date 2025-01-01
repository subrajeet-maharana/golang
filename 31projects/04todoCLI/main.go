package main

import "fmt"

type Todo struct {
  ID int
  Task string
  Completed bool
}

func showAllTasks(todos []Todo){
  fmt.Println("All Tasks")
  for _, todo := range todos {
    fmt.Printf("ID: %d | Task: %s | ", todo.ID, todo.Task)
    var completionStatus string = "Completed"
    if(!todo.Completed) {
      completionStatus ="In Progress"
    }
    fmt.Printf("Status: %s\n", completionStatus)
  }
}

func addNewTask(todos []Todo) {
  var task string
  fmt.Scanln("Add task: ", &task)
  newTodo := Todo{
    ID: len(todos),
    Task: task,
    Completed: false,
  }
  todos = append(todos, newTodo)
  fmt.Println("Todo added successfully...")
}

func main() { 
  fmt.Println("Welcome to the TO-DO CLI App")
  var todos []Todo
  for {
    var option int
    fmt.Printf("\nPlease choose an option:\n1. View All Tasks\n2. Add a new Task\n3. Mark a task as Completed\n4. Delete a Task\n5. Exit\n")
    fmt.Scanln(option)
    switch option {
    case 1:
      showAllTasks(todos)
    case 2:
      addNewTask(todos)
    }
  }
}
