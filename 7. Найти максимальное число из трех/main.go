package main

import "../base/console"

func main() {
	var first, second, third int

	first = console.ReadInt("Введите первое число: ")
	second = console.ReadInt("Введите второе число: ")
	third = console.ReadInt("Введите третье число: ")

	if first > second && first > third {
		console.Writeln(first, " is max")
		return
	}

	if second > third {
		console.Writeln(second, "is max")
		return
	}

	console.Writeln(third, "is max")
}
