// main.go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Simple Calculator")
	fmt.Println("------------------")

	num1 := getUserInput("Enter the first number:")
	operator := getUserInput("Enter the operator (+, -, *, /):")
	num2 := getUserInput("Enter the second number:")

	num1Float, err1 := strconv.ParseFloat(num1, 64)
	num2Float, err2 := strconv.ParseFloat(num2, 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input. Please enter valid numbers.")
		os.Exit(1)
	}

	result := calculate(num1Float, operator, num2Float)
	fmt.Printf("%.2f %s %.2f = %.2f\n", num1Float, operator, num2Float, result)
}

func getUserInput(prompt string) string {
	var input string
	fmt.Print(prompt + " ")
	fmt.Scanln(&input)
	return input
}

// ... (rest of the calculate function and other operations as before)

// main.go (continued)

func calculate(num1 float64, operator string, num2 float64) float64 {
	switch operator {
	case "+":
		return add(num1, num2)
	case "-":
		return subtract(num1, num2)
	case "*":
		return multiply(num1, num2)
	case "/":
		return divide(num1, num2)
	default:
		fmt.Println("Invalid operator. Please use +, -, *, /")
		os.Exit(1)
		return 0
	}
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Cannot divide by zero.")
		os.Exit(1)
	}
	return a / b
}
