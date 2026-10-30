package main

import "fmt"

const deliveryFee = 99.0

func main() {

	var name string
	var price float64
	var quantity int

	fmt.Print("Название товара: ")
	fmt.Scan(&name)
	fmt.Print("Цена: ")
	fmt.Scan(&price)
	fmt.Print("Количество товара: ")
	fmt.Scan(&quantity)

	cost := price * float64(quantity)
	total := cost + deliveryFee

	fmt.Printf("%s: %d шт. * %.2f = %.2f руб\n", name, quantity, price, cost)
	fmt.Printf("Доставка: %.2f руб.\n", deliveryFee)
	fmt.Println("-----------------------")
	fmt.Printf("Итог: %.2f руб.\n", total)
}
