package main

import (
	"../base/console"
	"fmt"
)

func main() {

	deposit := console.ReadFloat("Введите сумму вклада: ")
	depositRate := console.ReadFloat("Введите размер процентной ставки: ")
	depositTerm := console.ReadFloat("Введите срок вклада (в месяцах): ")

	income := (deposit * depositRate * depositTerm * 30) / (365 * 100)

	console.Writeln("Сумма вклада по истечению срока составит ")
	fmt.Printf("%.2f", income+deposit)
	console.Writeln("Прибыль")
	fmt.Printf("%.2f", income)
}
