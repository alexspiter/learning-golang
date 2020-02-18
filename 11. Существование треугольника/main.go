package main

import "../base/console"

func main() {
	var sideA, sideB, sideC int

	sideA = console.ReadInt("Введите длину первого отрезка: ")
	sideB = console.ReadInt("Введите длину второго отрезка: ")
	sideC = console.ReadInt("Введите длину третьего отрезка: ")

	if sumIsBigger(sideA, sideB, sideC) && sumIsBigger(sideA, sideC, sideB) && sumIsBigger(sideB, sideC, sideA) {
		console.Writeln("Существование треугольника возможно.")
	} else {
		console.Writeln("Существование треугольника невозможно.")
	}
}

func sumIsBigger(a, b, c int) bool {
	return a+b >= c
}
