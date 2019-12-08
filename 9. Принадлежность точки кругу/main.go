package main

import "../base/console"

func main() {
	var circleX, circleY, pointX, pointY, circleRadius, x, y, formula, squaredR int

	console.Writeln("Введите координаты центра окружности: ")
	circleX = console.ReadInt("Введите x: ")
	circleY = console.ReadInt("Введите y: ")
	circleRadius = console.ReadInt("Введите circleRadius: ")

	console.Writeln("Введите координаты точки: ")
	pointX = console.ReadInt("Введите x: ")
	pointY = console.ReadInt("Введите y: ")

	// формула: (x1 - x2)^2 + (y1 - y2)^2 = r^2
	x = pointX - circleX
	y = pointY - circleY
	formula = (x * x) + (y * y)
	squaredR = circleRadius * circleRadius

	if formula == squaredR {
		console.Writeln("Точка находится на окружности")
		return
	}

	if formula < squaredR {
		console.Writeln("Точка находится внутри круга")
		return
	}

	console.Writeln("Точка находится вне круга")
}
