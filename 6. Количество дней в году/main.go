package main

import "../base/console"

func main() {
	var year int
	var dividedBy400, dividedBy100, dividedBy4 bool

	year = console.ReadInt("Введите год: ")
	dividedBy400 = year%400 == 0
	dividedBy100 = year%100 == 0
	dividedBy4 = year%4 == 0

	if dividedBy400 || (dividedBy4 && !dividedBy100) {
		console.Writeln(year, " - високосный год, 366 дней.")
		return
	}

	console.Writeln(year, " - невисокосный год, 365 дней.")
}
