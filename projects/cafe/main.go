package main

import "fmt"

const coffePrice = 249.00
const curassantPrice = 124.00
const waterPrice = 130.00

func main() {
	var name string
	var price float64
	var quantity int
	var choice int

	fmt.Println("======== Меню ========")
	fmt.Printf("1. Кофе         %.2f\n", coffePrice)
	fmt.Printf("2. Курассан     %.2f\n", curassantPrice)
	fmt.Printf("3. Вода         %.2f\n", waterPrice)

	fmt.Print("Выберите номер товара: ")
	fmt.Scan(&choice)

	fmt.Print("Выберите количество товара: ")
	fmt.Scan(&quantity)

	if choice == 1 {
		name = "Кофе"
		price = coffePrice
	} else if choice == 2 {
		name = "Курассан"
		price = curassantPrice
	} else if choice == 3 {
		name = "Вода"
		price = waterPrice
	}

	if name == "" {
		fmt.Println("Такого товара нету")
	} else if quantity <= 0 {
		fmt.Println("Нужно положительное количество")
	} else {
		cost := price * float64(quantity)
		fmt.Printf("%s: %d шт. * %.2f = %.2f руб.\n", name, quantity, price, cost)
	}
}
