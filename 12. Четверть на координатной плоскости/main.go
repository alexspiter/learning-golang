package main

import "../base/console"

func main() {
	var x, y int

	console.Writeln("Введите координаты точки")
	x = console.ReadInt("Введите x: ")
	y = console.ReadInt("Введите y: ")

	if x > 0 && y > 0 {
		console.Writeln("Точка с заданными координатами находится в I четверти.")
		return
	}

	if x > 0 && y < 0 {
		console.Writeln("Точка с заданными координатами находится во II четверти.")
		return
	}

	if x < 0 && y < 0 {
		console.Writeln("Точка с заданными координатами находится в III четверти.")
		return
	}

	if x < 0 && y > 0 {
		console.Writeln("Точка с заданными координатами находится в VI четверти.")
		return
	}

	if x == 0 && y != 0 || y == 0 && x != 0 {
		console.Writeln("Точка с заданными координатами находится на оси координат.")
		return
	}

	console.Writeln("Точка с координатами [0, 0] находится в центре координатной плоскости.")
}
