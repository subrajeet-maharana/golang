package main

import "fmt"

func celsiusToFahrenheit(temp float32) float32 {
  return (9*temp/5)+32
}
func main(){
  fmt.Println("Welcome to our Temperature Converter")
  fmt.Println("Temp in Fahrenheit is: ", celsiusToFahrenheit(86))
  fmt.Println("Temp in Fahrenheit is: ", celsiusToFahrenheit(-40))
}
