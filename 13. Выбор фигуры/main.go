package main

import (
	"../base/console"
)

const Square = 1
const Rectangle = 2
const Triangle = 3
const Circle = 4

func main() {
	console.Writeln("Введите номер фигуры (1 - квадрат, 2 - прямоугольник, 3 - треугольник, 4 - круг")
	x := console.ReadInt()

	switch x {
	case Square:
		console.Writeln("квадрат")
	case Rectangle:
		console.Writeln("прямоугольник")
	case Triangle:
		console.Writeln("треугольник")
	case Circle:
		console.Writeln("круг")
	default:
		console.Writeln("Значение переменной не определено.")
	}

}
