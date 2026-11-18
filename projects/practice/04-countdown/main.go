package main

import "fmt"

func main() {
	var number int

	fmt.Print("Укажи число отсчета: ")
	fmt.Scan(&number)

	for i := number; i >= 1; i-- {
		fmt.Println(i)
	}
	fmt.Println("Старт!")
}
