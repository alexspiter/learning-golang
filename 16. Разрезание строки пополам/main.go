package main

import (
	"../base/console"
)

func main() {
	var text string

	console.Writeln("Введите текст:")
	text = console.ReadString()
	half := len(text) / 2

	console.Writeln(text[:half])
	console.Writeln(text[half:])
}
