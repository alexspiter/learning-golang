package main

import "fmt"

func main() {
	var first, second, division1 int
	fmt.Println("Enter the first number:")
	fmt.Scan(&first)
	fmt.Println("Enter the second number:")
	fmt.Scan(&second)
	division1 = first % second
	fmt.Println("The remainder of division is ", division1)
}
