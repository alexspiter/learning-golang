package main

import "../base/console"

func main() {

	var length1, length2, length3 int

	length1 = console.ReadInt("Введите длину первого отрезка: ")
	length2 = console.ReadInt("Введите длину второго отрезка: ")
	length3 = console.ReadInt("Введите длину третьего отрезка: ")

	if length1+length2 >= length3 && length2+length3 >= length1 && length1+length3 >= length2 {
		console.Writeln("Существование треугольника возможно.")
	} else {
		console.Writeln("Существование треугольника невозможно.")
	}
}
