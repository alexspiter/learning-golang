package main

import "../base/console"

func main() {
	var first, second, third int
	first = console.ReadInt("Введите первое число: ")
	second = console.ReadInt("Введите второе число: ")
	third = console.ReadInt("Введите третье число: ")

	if isBetween(first, second, third) {
		console.Writeln(first, " - среднее число")
		return
	}

	if isBetween(second, first, third) {
		console.Writeln(second, " - среднее число")
		return
	}

	console.Writeln(third, " - среднее число")
}

func isBetween(num, first, second int) bool {
	return first < num && num < second || second < num && num < first
}
