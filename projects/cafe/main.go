package main

import "fmt"

const coffePrice = 249.00
const waterPrice = 149.00
const croissantPrice = 124.00

func main() {
	total := 0.0

	fmt.Println("=========== Меню ===========")
	fmt.Printf("1.Кофе            %.2f\n", coffePrice)
	fmt.Printf("2.Вода            %.2f\n", waterPrice)
	fmt.Printf("3.Круассан        %.2f\n", croissantPrice)

	for {
		var choice int
		var quantity int
		var name string
		var price float64

		fmt.Print("Выбери товар (0 - закончить): ")
		fmt.Scan(&choice)

		if choice == 0 {
			break
		}

		switch choice {
		case 1:
			name = "Кофе"
			price = coffePrice
		case 2:
			name = "Вода"
			price = waterPrice
		case 3:
			name = "Круассан"
			price = croissantPrice
		default:
			fmt.Println("Такого товара нет")
			continue
		}

		fmt.Print("Введите количество товара: ")
		fmt.Scan(&quantity)

		if quantity <= 0 {
			fmt.Println("Неправильное количество товара")
			continue
		}

		cost := float64(quantity) * price
		total += cost
		fmt.Printf("%s: %d шт. * %.2f = %.2f руб.\n", name, quantity, price, cost)
	}
	fmt.Println("---------------------------")
	fmt.Printf("Итого:     %.2f\n", total)
}
