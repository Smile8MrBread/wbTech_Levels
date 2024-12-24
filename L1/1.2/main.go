package main

import (
	"fmt"
	"sync"
)

// Worker pool
func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{} // для отслеживания завершения всех горутин

	resChan := make(chan int) // для удобства вывода
	doChan := make(chan int)  // числа для воркеров
	go func() {
		for i := range arr {
			doChan <- arr[i]
		}
		close(doChan)
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for a := range doChan {
				resChan <- a * a
			}
		}()
	}

	// Ждем всех воркеров и закрываем канал,
	//чтобы начать с него чтение
	go func() {
		wg.Wait()
		close(resChan)
	}()

	for a := range resChan {
		fmt.Println(a)
	}
}
