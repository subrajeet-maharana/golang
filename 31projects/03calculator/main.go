package main

import "fmt"

func main(){
  fmt.Println("Welcome to the Integrity Calculator")
  fmt.Printf("1.Addition\n2.Subtraction\n3.Multiplication\n4.Division\n");
  for {
      fmt.Println("Choose option for calcuation.")
      var option, num1, num2 float32
      fmt.Scanln(&option)
      fmt.Println("Enter the numbers: ")
      fmt.Scanln(&num1, &num2) 
      fmt.Printf("The result is: ")
      switch option {
      case 1:
        fmt.Printf("%f\n", num1+num2) 
      case 2:
        fmt.Printf("%f\n", num1-num2)
      case 3:
        fmt.Printf("%f\n", num1*num2) 
      case 4:
        fmt.Printf("%f\n", num1/num2)
      }
  }
}
