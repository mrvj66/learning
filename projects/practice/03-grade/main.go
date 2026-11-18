package main

import "fmt"

func main() {
	var score int

	fmt.Print("Введите баллы (0-100): ")
	fmt.Scan(&score)

	switch {
	case score < 0 || score > 100:
		fmt.Println("Таких баллов не бывает")
	case score >= 90:
		fmt.Println("Отлично")
	case score >= 70:
		fmt.Println("Хорошо")
	case score >= 50:
		fmt.Println("Удволетворительно")
	case score >= 0:
		fmt.Println("Попробуй еще раз!")
	}
}
