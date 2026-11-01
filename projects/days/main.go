package main

import "fmt"

func main() {

	var name string
	var kind string
	var number int

	fmt.Print("Введите номер дня(1-7): ")
	fmt.Scan(&number)

	switch number {
	case 1:
		name = "Понедельник"
	case 2:
		name = "Вторник"
	case 3:
		name = "Среда"
	case 4:
		name = "Четверг"
	case 5:
		name = "Пятница"
	case 6:
		name = "Суббота"
	case 7:
		name = "Воскресенье"
	}

	switch number {
	case 1, 2, 3, 4, 5:
		kind = "рабочий"
	case 6, 7:
		kind = "выходной"
	}

	if name == "" {
		fmt.Println("Такого дня нет")
	} else {
		fmt.Printf("%s - %s\n", name, kind)
	}
}
