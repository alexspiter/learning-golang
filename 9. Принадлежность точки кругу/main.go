package main

import "../base/console"

func main() {

	var x, y, x1, y1, r, dotX, dotY int

	console.Writeln("Введите координаты центра окружности: ")
	x = console.ReadInt("Введите x: ")
	y = console.ReadInt("Введите y: ")
	r = console.ReadInt("Введите r: ")

	console.Writeln("Введите координаты точки: ")
	dotX = console.ReadInt("Введите x: ")
	dotY = console.ReadInt("Введите y: ")

	x1 = x + r
	y1 = y + r
	(x * x) + (y * -y) == r * r // уравнение окружности

	console.Writeln(squaredR)

	if

}
