package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	exit := make(chan struct{}, 1)
	input := make(chan int)
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second*3) // Контекст с таймаутом
	defer cancel1()
	ctx2, cancel2 := context.WithCancel(context.Background()) // Контекст с cancel()

	// Остановка горутины по таймауту
	go func(ctx context.Context) {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine 1 is done")
			break
		}
	}(ctx1)

	// Остановка горутины по каналу для выхода
	go func() {
		defer wg.Done()
		select {
		case <-exit:
			fmt.Println("Goroutine 2 is done")
			break
		}
	}()

	// Остановка в случае закрытия канала для передачи данных
	go func() {
		defer wg.Done()
		for {
			el, ok := <-input
			if !ok {
				fmt.Println("Goroutine 3 is done")
				break
			} else {
				fmt.Println(el)
			}
		}
	}()

	// Остановка горутины по cancel() функции
	go func(ctx context.Context) {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine 4 is done")
			break
		}
	}(ctx2)

	time.Sleep(time.Second)

	cancel2()
	exit <- struct{}{}
	close(exit)
	close(input)

	wg.Wait()
}
