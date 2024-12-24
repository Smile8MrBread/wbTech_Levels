package main

import "fmt"

// Немного алгосов ;)
func main() {
	s := "главрыбаa"
	runes := []rune(s)

	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}

	fmt.Println(string(runes))
}
