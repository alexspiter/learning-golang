package main

import "../base/console"

func main() {
	var first, second int

	console.Writeln("Введите произвольные четное и нечетное числа")
	first = console.ReadInt("Введите первое число: ")
	second = console.ReadInt("Введите второе число: ")

	var isOdd = first|1 == first
	if isOdd {
		console.Writeln(first, "is an odd number.")
	} else {
		console.Writeln(second, "is an odd number.")
	}
}
