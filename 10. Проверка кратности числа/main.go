package main

import "../base/console"

func main() {
	var first, second int

	console.Writeln("Введите два числа.")
	first = console.ReadInt("Введите первое число: ")
	second = console.ReadInt("Введите второе число: ")

	if first%second == 0 {
		console.Writeln("Первое число делится на второе без остатка")
	} else {
		console.Writeln("Первое число не делится на второе без остатка")
	}
}
