package main

import "fmt"

func main() {
	var a int
	var op string
	var b int

	fmt.Print("Укажи первое число: ")
	fmt.Scan(&a)

	fmt.Print("Укажи действие: +, -, *, /, %: ")
	fmt.Scan(&op)

	fmt.Print("Укажи второе число: ")
	fmt.Scan(&b)

	if op == "+" {
		fmt.Printf("%d + %d = %d\n", a, b, a+b)
	} else if op == "-" {
		fmt.Printf("%d - %d = %d\n", a, b, a-b)
	} else if op == "*" {
		fmt.Printf("%d * %d = %d\n", a, b, a*b)
	} else if op == "/" {
		if b == 0 {
			fmt.Println("Нельзя делить на ноль")
		} else {
			fmt.Printf("%d / %d = %.2f\n", a, b, float64(a)/float64(b))
		}
	} else if op == "%" {
		if b == 0 {
			fmt.Println("Деление на ноль нельзя")
		} else {
			fmt.Printf("%d %% %d = %d\n", a, b, a%b)
		}
	} else {
		fmt.Println("Неизвестное действие")
	}

}
