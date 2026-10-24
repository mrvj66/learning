package main

import "fmt"

func main() {

	coffe := "Кофе"
	coffePrice := 249.90
	coffeQuantity := 3
	coffeCost := coffePrice * float64(coffeQuantity)

	curassan := "Курасан"
	curassanPrice := 120.50
	curassanQuantity := 2
	curassanCost := curassanPrice * float64(curassanQuantity)

	water := "Вода"
	waterPrice := 59.00
	waterQuantity := 1
	waterCost := waterPrice * float64(waterQuantity)

	fmt.Printf("%s: %d шт. * %.2f = %.2f руб\n", coffe, coffeQuantity, coffePrice, coffeCost)
	fmt.Printf("%s: %d шт. * %.2f = %.2f руб\n", curassan, curassanQuantity, curassanPrice, curassanCost)
	fmt.Printf("%s: %d шт. * %.2f = %.2f руб\n", water, waterQuantity, waterPrice, waterCost)

	total := coffeCost + curassanCost + waterCost
	fmt.Println("-----------------------")
	fmt.Printf("Итог: %.2f руб.\n", total)

}
