package main

import (
	"../base/console"
	"math"
)

const Square = 1
const Rectangle = 2
const Triangle = 3
const Circle = 4

func main() {

	var figureNumber int
	figureNumber = console.ReadInt("Введите номер фигуры: 1 - квадрат, 2 - прямоугольник, 3 - треугольник, 4 - круг\n")

	switch figureNumber {
	case Square:
		side1 := console.ReadFloat("Введите длину ")
		console.Writeln("Площадь квадрата равна ", side1*side1)
	case Rectangle:
		side1 := console.ReadFloat("Введите длину ")
		side2 := console.ReadFloat("Введите ширину ")
		console.Writeln("Площадь прямоугольника равна ", side1*side2)
	case Triangle:
		side1 := console.ReadFloat("Введите высоту ")
		side2 := console.ReadFloat("Введите длину основания ")
		console.Writeln("Площадь треугольника равна", 0.5*side1*side2)
	case Circle:
		r := console.ReadFloat("Введите радиус круга ")
		console.Writeln("Площадь круга равна", math.Pi*(r*r))
	default:
		console.Writeln("Значение переменной не определено.")
	}
}
