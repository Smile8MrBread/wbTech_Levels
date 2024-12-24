package main

import (
	"fmt"
	"strings"
)

func main() {
	words := "snow dog sun"

	// По аналогии с 1.19
	str := strings.Split(words, " ")
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]
	}

	fmt.Println(strings.Join(str, " "))
}
