package main

import (
	"awesomeProject/base/console"
	"awesomeProject/base/rand"
	"fmt"
)

func main() {
	console.Writeln(rand.RandomInt(100, 999))

	var generatedNumber, ones, tens, hundreds int
	var sum, multiplication int

	generatedNumber = rand.RandomInt(100, 999)
	ones = generatedNumber % 10
	tens = generatedNumber % 100 / 10
	hundreds = generatedNumber % 1000 / 100
	fmt.Println(generatedNumber, ones, tens, hundreds)

	sum = ones + tens + hundreds
	multiplication = ones * tens * hundreds

	fmt.Println("The random generated number is: ", generatedNumber)
	fmt.Println("The sum of numerals is: ", sum)
	fmt.Println("The multiplication is: ", multiplication)
}
