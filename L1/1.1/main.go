package main

import "fmt"

// Родительская структура
type Human struct {
	Name string
	Age  int
}

// Дочерняя структура
type Action struct {
	Human // Встраивание
	Sport string
}

func (h *Human) SayHello() {
	fmt.Printf("Hi, i'm %s, %d year old\n", h.Name, h.Age)
}

func (a *Action) SaySport() {
	fmt.Printf("My sport is %s\n", a.Sport)
}

func main() {
	human := Human{
		Name: "Ilya",
		Age:  24,
	}

	action := Action{
		Human: human,
		Sport: "hobie-horsing",
	}

	// Из дочерней можно использовать методы родительской
	action.SayHello()
	action.SaySport()
}
