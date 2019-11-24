package main

import "../base/console"

func main() {

	var year int

	year = console.ReadInt("Введите год: ")

	if year%4 == 0 {
		console.Writeln(year, " - високосный год.")
	} else if year%400 == 0 {
		console.Writeln(year, " - високосный год.")
	} else {
		console.Writeln(year, " - невисокосный год.")
	}
}
