package main

import "fmt"

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
	return a / b
}

func main() {
	var a, b float64
	var operator string

	fmt.Println("Enter the first number:")
	fmt.Scanln(&a)

	fmt.Println("Enter Operator(+, -, *, /)")
	fmt.Scanln(&operator)

	fmt.Println("Enter the Second number:")
	fmt.Scanln(&b)

	var result float64

	switch operator {
	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		result = divide(a, b)
	default:
		fmt.Println("Unknown operator")
		//result = ("Unknown Operator")

	}

	fmt.Println("Answer=", result)

}
