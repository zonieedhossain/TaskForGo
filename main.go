package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
getOperation:
	var operation string
	fmt.Print("Please select an operation: +, -, *, / : ")
	fmt.Scanln(&operation)

	var num1 string
	fmt.Print("Please input the first number: ")
	fmt.Scanln(&num1)

	var num2 string
	fmt.Print("Please input the second number: ")
	fmt.Scanln(&num2)

	switch operation {
	case "+":
		var firstNumber int = stringToInt(num1)
		var secondNumber int = stringToInt(num2)
		fmt.Print("Result: ")
		fmt.Println(calc.Add(firstNumber, secondNumber))
	case "-":
		var firstNumber int = stringToInt(num1)
		var secondNumber int = stringToInt(num2)
		fmt.Print("Result: ")
		fmt.Println(calc.Subtract(firstNumber, secondNumber))

	case "*":
		var firstNumber float64 = stringToFloat64(num1)
		var secondNumber float64 = stringToFloat64(num2)
		fmt.Print("Result: ")
		fmt.Println(calc.Multiply(firstNumber, secondNumber))

	case "/":
		var firstNumber float64 = stringToFloat64(num1)
		var secondNumber float64 = stringToFloat64(num2)
		fmt.Print("Result: ")
		fmt.Println(calc.Divide(firstNumber, secondNumber))

	default:
		fmt.Println("Invalid operation selected. Please try again!")
		goto getOperation
	}
}

func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}

func stringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return f
}
