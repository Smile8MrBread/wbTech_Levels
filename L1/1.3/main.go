package main

import (
	"fmt"
	"sync"
)

const wpCount = 3 // Количество воркеров

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	resChan := make(chan int)
	doChan := make(chan int)
	go func() {
		for i := range arr {
			doChan <- arr[i]
		}
		close(doChan)
	}()

	for i := 0; i < wpCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for a := range doChan {
				resChan <- a * a
			}
		}()
	}

	// Ждем всех воркеров и после закрываем канал, чтобы начать чтение
	go func() {
		wg.Wait()
		close(resChan)
	}()

	res := 0
	for a := range resChan {
		res += a
	}

	fmt.Println(res)
}
