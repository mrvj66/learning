package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Сколько тебе лет? ")
	fmt.Scan(&age)

	fmt.Print("Какое твое имя? ")
	fmt.Scan(&name)

	futureAge := age + 10

	fmt.Printf("Привет, %s! Через 10 лет тебе будет %d\n", name, futureAge)
}
