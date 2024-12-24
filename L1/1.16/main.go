package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var l, r []int
	p := rand.Intn(len(arr)) // рандомный "пивот"
	for i := range arr {
		if i == p { // Позже будет добавлен вручную
			continue
		}
		// Разбиваем на 2 массива: в правом все больше
		// и то что равно писоту, а влевом все что меньше
		if arr[i] >= arr[p] {
			r = append(r, arr[i])
		} else {
			l = append(l, arr[i])
		}
	}

	// Рекуривно прогоняем так же левую и правую части
	left := quickSort(l)
	right := quickSort(r)

	// Все объединяем и возвращаем
	left = append(left, arr[p])
	left = append(left, right...)
	return left
}

func main() {
	arr := []int{3, 5, 1, 5, 2, 0}

	fmt.Println(quickSort(arr))
}
