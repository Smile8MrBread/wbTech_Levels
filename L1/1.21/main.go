package main

import "fmt"

// Очень много пустого кода
// Резюмируя сделанное:
// Как я понял ,мы делаем общий интерфейс,
// который будут реализовывать адаптеры каждой из структур,
// а методы этих адаптеров будут вызывать соответствующие
// методы соответствующих структур
type Dog struct{}
type Cat struct{}

type DogAdapter struct {
	*Dog
}
type CatAdapter struct {
	*Cat
}

func (d *Dog) DogWof() {
	fmt.Println("Woof, woof")
}
func (c *Cat) CatMeow() {
	fmt.Println("Meow, meow")
}
func (c *Cat) CatName() {
	fmt.Println("My name is Barsik")
}

type AnimalsAdapter interface {
	Reaction()
}

func (d *DogAdapter) Reaction() {
	d.DogWof()
}

func (c *CatAdapter) Reaction() {
	c.CatMeow()
	c.CatName()
}

func NewDogAdapter(dog *Dog) *DogAdapter {
	return &DogAdapter{dog}
}

func NewCatAdapter(cat *Cat) *CatAdapter {
	return &CatAdapter{cat}
}

func main() {
	animals := []AnimalsAdapter{NewDogAdapter(&Dog{}), NewCatAdapter(&Cat{})}

	for _, el := range animals {
		el.Reaction()
	}
}
