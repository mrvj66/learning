package main

import "fmt"

func main() {
	a := 17
	b := 5

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)
	fmt.Printf("%d / %d = %.2f\n", a, b, float64(a)/float64(b))
}
