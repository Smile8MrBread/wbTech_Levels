package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[int][]float64)
	t := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sort.Float64s(t) // Сортирую для удобства

	temp := 0
	for i := range t {
		// Проверяю, что температура в шаге 10
		if i == 0 || t[i] >= float64(temp+10) {
			temp = int(t[i]) - int(t[i])%10 // Убераю единицы
			m[temp] = append(m[temp], t[i])
		} else {
			m[temp] = append(m[temp], t[i])
		}
	}

	for key, val := range m {
		fmt.Println(key, "-", val)
	}
}
