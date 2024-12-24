package main

var justString string

// Как понимаю - строка v будет все время существовать, т.к justString ссылается на v
// Чтобы создать новую строку, надо текущий срез преобразовать в массив byte, чтобы создался отдельный
// массив byte и уже его преобразовать в строку
func someFunc() {
	v := "dsadasdsadsadasdasdasdsadsadasdasdsaddsadasdsadsadasdasdasdsadsadasdasdsaddsadasdsadsadasdasdasdsadsadasdasdsaddsadasdsadsadasdasdasdsadsadasdasdsaddsadasdsadsadasdasdasdsadsadasdasdsaddsadasdsadsadasdasdasdsadsadasdasdsaddsadasdsadsadasdasdasdsadsadasdasdsaddsadasdsadsadasdasdasdsadsadasdasdsad"
	justString = string([]byte(v[:100]))
}
func main() {
	someFunc()
}
