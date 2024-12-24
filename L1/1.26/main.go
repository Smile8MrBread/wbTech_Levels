package main

import (
	"fmt"
	"strings"
)

func CheckUniq(str string) bool {
	str = strings.ToLower(str) // Перевожу все в нижний регистр
	// Структура ничего не весит
	m := make(map[string]struct{})

	for i := range str {
		if _, in := m[string(str[i])]; in {
			return false
		}

		m[string(str[i])] = struct{}{}
	}

	return true
}

func main() {
	str := "abcd"

	fmt.Println(CheckUniq(str))
}
