package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	m := make(map[int]int)
	mu := sync.Mutex{} // Для конкурентной записи
	wg := sync.WaitGroup{}
	input := make(chan int, 100)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			input <- rand.Intn(10)
			time.Sleep(time.Second / 2)
		}
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-exit:
				break
			default:
				for a := range input {
					// Перед тем чтобы изменить данные - лочим мапу для
					// остальных горутинок на чтение и запись, т.к во время
					// изменения кто-то или мы можем взять уже неактуальные данные
					mu.Lock()
					m[a]++
					fmt.Printf("%d: %d\n", a, m[a])
					mu.Unlock()
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(input)
		close(exit)
	}()

	<-exit
}
