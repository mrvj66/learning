package main

import "fmt"

func main() {

	for {
		var number int

		fmt.Print("Введите число(или 0 для выхода): ")
		fmt.Scan(&number)

		if number == 0 {
			fmt.Println("Выход из программы")
			break
		}

		if number%2 == 0 {
			fmt.Printf("%d - число четное\n", number)
		} else {
			fmt.Printf("%d - нечетное число\n", number)
		}
	}
}
