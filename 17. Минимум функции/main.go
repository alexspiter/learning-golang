package main

import "../base/console"

func main() {
	var a, b, x float64

	a = console.ReadFloat("Введите значение a: ")
	b = console.ReadFloat("Введите значение b: ")
	console.ReadInt("Введите значение с: ")

	x = -b / 2 * a

	if a < 0 {
		console.Writeln("При a<0, ветви функции направлены вниз, т.е. x - максимум.")
		return
	}
	if a == 0 {
		console.Writeln("При а = 0 решение невозможно")
		return
	}
	console.Writeln("Минимум функции равен: ", x)

}
