package main

import "../base/console"

func main() {
	var first, second, third int
	first = console.ReadInt("Введите первое число: ")
	second = console.ReadInt("Введите второе число: ")
	third = console.ReadInt("Введите третье число: ")

	if first > second && second > third || first < second && second < third {
		console.Writeln(second, " - среднее число")
		return
	}

	if first < second && first > third || first > second && first < third {
		console.Writeln(first, " - среднее число")
	} else {
		console.Writeln(third, " - среднее число")
	}
}
