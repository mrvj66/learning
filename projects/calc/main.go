package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Укажи первое число: ")
	fmt.Scan(&a)

	fmt.Print("Укажи второе число: ")
	fmt.Scan(&b)

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)
	fmt.Printf("%d / %d = %.2f\n", a, b, float64(a)/float64(b))
}
