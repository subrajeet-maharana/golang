package main

import "fmt"

type Todo struct {
  ID int
  Task string
  Completed bool
}

func main() { 
  fmt.Println("Welcome to the TO-DO CLI App")
  var todos []Todo
  for {
    fmt.Printf("\nPlease choose an option:\n1. View All Tasks\n2. Add a new Task\n3. Mark a task as Completed\n4. Delete a Task\n5. Exit\n")
  }
}
