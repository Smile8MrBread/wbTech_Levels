package main

import "fmt"

const (
	i      = 62    // номер бита
	isZero = false // установка на 0
)

func main() {
	var n int64 = 1234
	fmt.Println(n)
	fmt.Printf("%064b\n", n)

	// Гугл мне в помощь xD
	if isZero {
		n &^= 1 << i // Искл и
	} else {
		n |= 1 << i // Искл или
	}

	fmt.Println(n)
	fmt.Printf("%064b\n", n) // Вывод 64 нулей - битов, которые берут значение n
}
