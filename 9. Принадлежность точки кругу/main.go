package main

import "../base/console"

func main() {

	var x0, y0, x1, y1, r, x, y, formula, squaredR int

	console.Writeln("Введите координаты центра окружности: ")
	x0 = console.ReadInt("Введите x: ")
	y0 = console.ReadInt("Введите y: ")
	r = console.ReadInt("Введите r: ")

	console.Writeln("Введите координаты точки: ")
	x1 = console.ReadInt("Введите x: ")
	y1 = console.ReadInt("Введите y: ")

	// формула: (x1-x0)^2 + (y1-y0)^2 = r^2
	x = x1 - x0
	y = y1 - y0
	console.Writeln(x, y)
	formula = (x * x) + (y * y)
	console.Writeln(formula)
	squaredR = r * r
	console.Writeln(squaredR)

	if formula == squaredR {
		console.Writeln("Точка находится на окружности")
		return
	}
	if formula < squaredR {
		console.Writeln("Точка находится внутри круга")
	} else {
		console.Writeln("Точка находится вне круга")
	}
}
