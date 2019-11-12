package main

import (
	"awesomeProject/base/console"
	"fmt"
)

func main() {
	var first, second, result int
	var operator string

	console.Writeln("Enter the first number:")
	fmt.Scan(&first)
	console.Writeln("Enter the operator:")
	fmt.Scan(&operator)
	console.Writeln("Enter the second number:")
	fmt.Scan(&second)

	if operator == "+" {
		result = first + second
		println("The sum is: ", first, operator, second, "=", result)
	} else if operator == "-" {
		result = first - second
		println("The difference is: ", first, operator, second, "=", result)
	} else if operator == "*" {
		result = first * second
		println("The result of multiplication is: ", first, operator, second, "=", result)
	} else if operator == "/" {
		result = first / second
		println("The quotient is: ", first, operator, second, "=", result)
	} else {
		println("The entered data are not correct.")
	}
}
