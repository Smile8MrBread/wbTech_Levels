package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	arr[2], arr[0] = arr[0], arr[2]

	fmt.Println(arr)
}
