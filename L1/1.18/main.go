package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Counter struct {
	Count int
}

// Для наглядности
func (c *Counter) PlusOne() {
	c.Count++
}

func main() {
	counter := Counter{Count: 0}
	mu := sync.RWMutex{} // Вроде нужен, но не уверен
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	for i := 0; i < 3; i++ {
		go func() {
			for {
				select {
				case <-exit:
					break
				default:
					mu.RLock() //Блок на чтение
					counter.PlusOne()
					fmt.Println(counter.Count)
					mu.RUnlock()
					time.Sleep(time.Second)
				}
			}
		}()
	}

	<-exit
}
