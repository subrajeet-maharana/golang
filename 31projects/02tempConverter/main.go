package main

import "fmt"

func celsiusToFahrenheit(temp float32) float32 {
	return (9 * temp / 5) + 32
}

func fahrenheitToCelsius(temp float32) float32 {
	return ((temp - 32) * 5) / 9
}

func celsiusToKelvin(temp float32) float32 {
	return temp + 273.15
}

func kelvinToCelsius(temp float32) float32 {
	return temp - 273.15
}

func main() {
	fmt.Println("Welcome to Clarity Temperature Converter")
	var option, subOption, temp float32
	fmt.Printf("What is your input temperature format?\n1.Celsius\n2.Fahrenheit\n3.Kelvin\nChoose your option: ")
	fmt.Scanln(&option)
	fmt.Println("Enter the temperature: ")
	fmt.Scanln(&temp)
	switch option {
	case 1:
		fmt.Println("\nConvert to?\n1.Fahrenheit\n2.Kelvin\nChoose your option: ")
		fmt.Scanln(&subOption)
		switch subOption {
		case 1:
			fmt.Println("In Fahrenheit: ", celsiusToFahrenheit(temp))
		case 2:
			fmt.Println("In Kelvin: ", celsiusToKelvin(temp))
		}
	case 2:
		fmt.Println("\nConvert to?\n1.Celsius\n2.Kelvin\nChoose your option: ")
		fmt.Scanln(&subOption)
		var inCelsius float32 = fahrenheitToCelsius(temp)
		switch subOption {
		case 1:
			fmt.Println("In Celsius: ", inCelsius)
		case 2:
			fmt.Println("In Kelvin: ", celsiusToKelvin(inCelsius))
		}
	case 3:
		fmt.Println("\nConvert to?\n1.Celsius\n2.Fahrenheit\nChoose your option: ")
		fmt.Scanln(&subOption)
		var inCelsius float32 = kelvinToCelsius(temp)
		switch subOption {
		case 1:
			fmt.Println("In Celsius: ", inCelsius)
		case 2:
			fmt.Println("In Fahrenheit: ", celsiusToFahrenheit(inCelsius))
		}
	}
}
