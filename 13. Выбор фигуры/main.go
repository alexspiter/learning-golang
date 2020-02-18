package main

import "../base/console"

func main() {
	var x int
	console.Writeln("Введите номер фигуры: 1 - квадрат, 2 - прямоугольник, 3 - треугольник, 4 - круг")
	x = console.ReadInt()
	switch x {
	case 1:
		console.Writeln("квадрат")
	case 2:
		console.Writeln("прямоугольник")
	case 3:
		console.Writeln("треугольник")
	case 4:
		console.Writeln("круг")
	default:
		console.Writeln("Значение переменной не определено.")
	}

}
