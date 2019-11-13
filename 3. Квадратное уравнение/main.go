package main

import (
	"../base/console"
	"math"
)

func main() {
	var a, b, c, discriminant, scrtDiscriminant, doubleA, x1, x2 float64

	a = console.ReadFloat("Введите a: ")
	b = console.ReadFloat("Введите b: ")
	c = console.ReadFloat("Введите c: ")

	discriminant = b*b - 4*a*c

	if discriminant == 0 {
		x1 = -b / (2 * a)
		console.Writeln("x = ", x1)
		return
	}

	if discriminant < 0 {
		console.Writeln("Данное уравнение не имеет корней.")
		return
	}

	scrtDiscriminant = math.Sqrt(discriminant)
	doubleA = 2 * a
	x1 = (-b + scrtDiscriminant) / doubleA
	x2 = (-b - scrtDiscriminant) / doubleA
	console.Writeln("x1 =", x1)
	console.Writeln("x2 =", x2)
}
