package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Хотел изначально использовать float64, но не нравится экспонента
	a := big.Int{}
	b := big.Int{}
	a.SetString("1234567899", 10)
	b.SetString("123456789", 10)

	sum := new(big.Int).Add(&a, &b)
	dif := new(big.Int).Sub(&a, &b)
	mult := new(big.Int).Mul(&a, &b)
	div := new(big.Int).Div(&a, &b)

	fmt.Println("Sum is", sum)
	fmt.Println("Sub is", dif)
	fmt.Println("Mult is", mult)
	fmt.Println("Div is", div)
}
