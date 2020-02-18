package main

import (
	"../base/console"
	"strings"
)

func main() {
	var string string

	console.Writeln("Введите текст:")
	string = console.ReadString()
	console.Writeln(len(string))
	strings.Split(string, "")
	console.Writeln(string)
}
