package main

import "fmt"

const i = 2

func main() {
	arr := []int{1, 2, 3, 4, 5}
	var res []int

	// Вырезание кусочка
	res = append(res, arr[:i]...)
	res = append(res, arr[i+1:]...)
	fmt.Println(res)
}
