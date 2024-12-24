package main

import "fmt"

func main() {
	// Возможно не понял задание - реализовал типа set
	var res []int
	a := []int{0, 1, 2, 2, 3, 4}
	b := []int{1, 5, 5, 6, 1, 0}
	m := make(map[int]struct{}) // структура ничего не весит

	for i := range a {
		m[a[i]] = struct{}{}
	}

	for i := range b {
		m[b[i]] = struct{}{}
	}

	for i := range m {
		res = append(res, i)
	}

	fmt.Println(res)
}
