package console

import "fmt"

func ReadString(a ...interface{}) string {
	if len(a) > 0 {
		Write(a...)
	}

	var result string
	fmt.Scan(&result)

	return result
}

func ReadChar(a ...interface{}) rune {
	if len(a) > 0 {
		Write(a...)
	}

	var result rune
	fmt.Scan(&result)

	return result
}

func ReadInt(a ...interface{}) int {
	if len(a) > 0 {
		Write(a...)
	}

	var result int
	fmt.Scan(&result)

	return result
}

func ReadFloat(a ...interface{}) float64 {
	if len(a) > 0 {
		Write(a...)
	}

	var result float64
	fmt.Scan(&result)

	return result
}

func Write(a ...interface{}) {
	fmt.Print(a...)
}

func Writeln(a ...interface{}) {
	fmt.Println(a...)
}
