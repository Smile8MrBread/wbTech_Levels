package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 5 // Количество секунд

func main() {
	inputChan := make(chan int, 100)
	// Пустая структура ничего не весит
	exitChan := make(chan struct{}, 1)

	go func() {
		select {
		case <-exitChan:
			break
		default:
			for {
				inputChan <- rand.Int()
				time.Sleep(time.Second / 3)
			}
		}
	}()

	go func() {
		select {
		case <-exitChan:
			break
		default:
			for el := range inputChan {
				fmt.Println("Received:", el)
			}
		}
	}()

	// Ждем n секунд и заканчиваем программу
	time.Sleep(n * time.Second)
	exitChan <- struct{}{}
	close(inputChan)
	close(exitChan)
}
