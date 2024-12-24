package main

import (
	"fmt"
	"sync"
)

func main() {
	input := make(chan int)
	output := make(chan int)
	wg := sync.WaitGroup{}
	arr := []int{2, 4, 9, 11, 12, 33, 34, 45, 46, 69}

	wg.Add(3)
	// Данные входят
	go func() {
		defer wg.Done()
		for i := range arr {
			input <- arr[i]
		}
		close(input)
	}()

	// Данные считаются
	go func() {
		defer wg.Done()
		for a := range input {
			output <- a * 2
		}
		close(output)
	}()

	// Данные выходят
	go func() {
		defer wg.Done()
		for a := range output {
			fmt.Println(a)
		}
	}()

	wg.Wait()
}
