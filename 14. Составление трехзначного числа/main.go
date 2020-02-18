package main

import "../base/console"

func main() {

	var number, ones, tens, hundreds int

	number = console.ReadInt("Введите любое трехзначное число:")

	ones = number % 10
	tens = number % 100 / 10
	hundreds = number % 1000 / 100

	printNumber(ones, tens, hundreds)
	printNumber(ones, hundreds, tens)
	printNumber(tens, ones, hundreds)
	printNumber(tens, hundreds, ones)
	printNumber(hundreds, tens, ones)
	printNumber(hundreds, ones, tens)
}

func printNumber(a, b, c int) {
	console.Write(a)
	console.Write(b)
	console.Write(c)
	console.Writeln()
}
