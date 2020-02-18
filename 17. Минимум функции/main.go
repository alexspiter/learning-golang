package main

import "../base/console"

func main() {
	var a, b, doubleA, x float64

	a = console.ReadFloat("Введите значение a: ")
	b = console.ReadFloat("Введите значение b: ")
	console.ReadInt("Введите значение с: ")

	doubleA = 2 * a
	x = -b / doubleA

	if a < 0 {
		console.Writeln("При a<0, ветви функции направлены вниз, т.е. x - максимум.")
		return
	}
	if a == 0 {
		console.Writeln("При а = 0 решение невозможно")
		return
	}
	if a > 0 {
		console.Writeln("Минимум функции равен: ", x)
	}

}
