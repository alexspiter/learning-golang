package main

import "../base/console"
import "../base/rand"

func main() {
	var randNumber int
	randNumber = (rand.RandomInt(10, 99)) * 10
	console.Writeln("Случайное трёхзначное число, кратное десяти =", randNumber)
}
