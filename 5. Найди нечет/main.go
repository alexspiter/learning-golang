package main

import "../base/console"

func main() {
	var first, second int

	console.Writeln("Введите произвольные четное и нечетное числа")
	first = console.ReadInt("Введите первое число: ")
	second = console.ReadInt("Введите второе число: ")

	if first%2 != 0 {
		console.Writeln(first, "is an odd number.")
		return
	} else {
		console.Writeln(second, "is an odd number.")
	}
}
