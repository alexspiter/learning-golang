package main

import (
	"awesomeProject/base/console"
	"awesomeProject/base/rand"
)

func main() {
	console.Writeln(rand.RandomInt(100, 999))

	var generatedNumber, ones, tens, hundreds int
	var sum, multiplication int

	generatedNumber = rand.RandomInt(100, 999)
	ones = generatedNumber % 10
	tens = generatedNumber % 100 / 10
	hundreds = generatedNumber % 1000 / 100

	sum = ones + tens + hundreds
	multiplication = ones * tens * hundreds

	console.Writeln("The random generated number is: ", generatedNumber)
	console.Writeln("The sum of numerals is: ", sum)
	console.Writeln("The multiplication is: ", multiplication)
}
