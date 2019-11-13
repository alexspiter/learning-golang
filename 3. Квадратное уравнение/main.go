package main

import (
	"../base/console"
	"fmt"
	"math"
)

func main() {

	var a, b, c, D, x, x1, x2 float64

	a = console.ReadFloat("Введите a: ")
	console.Writeln("Введите b:")
	fmt.Scan(&b)
	console.Writeln("Введите c:")
	fmt.Scan(&c)

	D = b*b - 4*a*c

	if (D == 0) {
		x = -b / (2 * a)
		println("x = ", x)
	} else if (D < 0) {
		println("Данное уравнение не имеет корней")
	} else {
		x1 = (-b + math.Sqrt(D)) / 2 * a
		x2 = (-b - math.Sqrt(D)) / 2 * a
		console.Writeln("x1 =", x1)
		console.Writeln("x2 =", x2)
	}
}
