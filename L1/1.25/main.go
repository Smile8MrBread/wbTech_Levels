package main

import (
	"fmt"
	"time"
)

// Сомнительно, но окей
func Sleep(millisec int) {
	select {
	// Жду пока не пройдет время
	case <-time.After(time.Millisecond * time.Duration(millisec)):
		return
	}
}

func main() {
	fmt.Println(time.Now())
	Sleep(3000)
	fmt.Println(time.Now())
}
