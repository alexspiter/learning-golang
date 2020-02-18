package main

import (
	"../base/console"
	"math"
)

func main() {
	var x int
	var r, side1, side2, sSquare, sTriangle, sCircle float64

	x = console.ReadInt("Введите номер фигуры: 1 - квадрат, 2 - прямоугольник, 3 - треугольник, 4 - круг\n")

	switch x {
	case 1:
		side1 = console.ReadFloat("Введите длину ")
		sSquare = side1 * side1
		console.Writeln("Площадь квадрата равна ", sSquare)
	case 2:
		side1 = console.ReadFloat("Введите длину ")
		side2 = console.ReadFloat("Введите ширину ")
		sSquare = side1 * side2
		console.Writeln("Площадь прямоугольника равна ", sSquare)
	case 3:
		side1 = console.ReadFloat("Введите высоту ")
		side2 = console.ReadFloat("Введите длину основания ")
		sTriangle = 0.5 * side1 * side2
		console.Writeln("Площадь треугольника равна", sTriangle)
	case 4:
		r = console.ReadFloat("Введите радиус круга ")
		sCircle = math.Pi * (r * r)
		console.Writeln("Площадь круга равна", sCircle)
	default:
		console.Writeln("Значение переменной не определено.")
	}
}
