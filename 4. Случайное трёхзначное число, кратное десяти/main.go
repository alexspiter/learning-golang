package main

import "../base/console"
"../base/rand"

func main() {
	var randNumber int

	randNumber = (rand.RandomInt(10, 99)) * 10
	console.Writeln("Случайное трёхзначное число, кратное 10 =", randNumber)
}
