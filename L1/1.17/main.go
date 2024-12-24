package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	need := 1 // Искомое число
	l, r := 0, len(arr)-1

	for r-l != 1 { // Ищем до тех пор, пока 2 указателя не будут рядом
		m := (r + l) / 2 // Берем середину
		// Сдвигаем границы
		if arr[m] >= need {
			r = m
		} else {
			l = m
		}
	}

	if arr[l] == need {
		fmt.Println(l)
	} else {
		fmt.Println(r)
	}
}
