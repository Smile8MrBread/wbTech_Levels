package main

import "fmt"

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]struct{}) // Стркуктура ничего не весит

	for i := range str {
		m[str[i]] = struct{}{}
	}

	fmt.Println(m)
}
