package main

import "../base/console"

func main() {

	var number, ones, tens, hundreds int

	number = console.ReadInt("Введите любое трехзначное число:")

	ones = number % 10
	tens = number % 100 / 10
	hundreds = number % 1000 / 100

}
