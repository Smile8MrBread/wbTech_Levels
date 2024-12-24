package main

import "fmt"

func Type(a interface{}) {
	// Хватаю тип переменной и иду с ним по куйсам
	switch a.(type) {
	case int:
		fmt.Println(a, "is int")
	case string:
		fmt.Println(a, "is string")
	case bool:
		fmt.Println(a, "is bool")
	case chan any: // Не отработал как хотелось
		fmt.Println(a, "is chan")
	default:
		fmt.Println(a, "is unknown")
	}
}

func main() {
	var a int
	var b string
	var c bool
	var d chan interface{}
	var e chan bool

	Type(a)
	Type(b)
	Type(c)
	Type(d)
	Type(e)
}
