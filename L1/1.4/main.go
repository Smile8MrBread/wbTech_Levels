package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const wgCount = 3 // количество воркеров

func main() {
	exitChan := make(chan os.Signal, 1)                    // Канал для завершения работы, буфферезированыый чтобы избежать возможных последсвий
	signal.Notify(exitChan, os.Interrupt, syscall.SIGTERM) // Слушаем CTRL+C и обычный сигнал завершения
	taskChan := make(chan int, 100)
	wg := sync.WaitGroup{}

	// Поток данных
	go func() {
		for {
			taskChan <- rand.Int()
			//time.Sleep(time.Second / 10)
		}
	}()

	for i := 0; i < wgCount; i++ {
		wg.Add(1)
		go func(n int) {
			// Жду сигнал о выходе из программы, а иначе читаю
			// данные из канала потока и вывожу их в терминал
			select {
			case <-exitChan:
				break
			default:
				for el := range taskChan {
					fmt.Printf("%d: %d\n", n, el)
				}
			}
		}(i + 1) // передаю i для id воркера
	}

	// wg.Wait в горутине, чтобы он не мешал exitChan
	// и мог штатно отработать и закрыть каналы
	go func() {
		wg.Wait()
		close(taskChan)
		close(exitChan)
	}()

	<-exitChan
}
