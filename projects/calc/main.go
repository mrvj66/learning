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

	switch op {
	case "+":
		fmt.Printf("%d + %d = %d\n", a, b, a+b)
	case "-":
		fmt.Printf("%d - %d = %d\n", a, b, a-b)
	case "*":
		fmt.Printf("%d * %d = %d\n", a, b, a*b)
	case "/":
		if b == 0 {
			fmt.Println("Деление на ноль невозможно")
		} else {
			fmt.Printf("%d / %d = %.2f\n", a, b, float64(a)/float64(b))
		}
	case "%":
		if b == 0 {
			fmt.Println("Деление на ноль невозможно")
		} else {
			fmt.Printf("%d %% %d = %d\n", a, b, a%b)
		}
	default:
		fmt.Println("Неизвестное действие")
	}
}
