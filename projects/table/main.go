package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введи число: ")
	fmt.Scan(&n)

	for i := 1; i <= 20; i++ {
		fmt.Printf("%d * %d = %d\n", i, n, i*n)
	}
}
