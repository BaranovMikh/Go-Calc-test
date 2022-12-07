package main

import "fmt"

func main() {

	functions := map[string]func(int, int) int{
		"divide": divide,
	}

	var currentNumber int
	fmt.Println("Введите 1 число")
	fmt.Scan(&currentNumber)

	for true {
		var functioName string
		var number int
		fmt.Println("Оператор?")
		fmt.Scan(&functioName)
		if functioName == "done" {
			break
		}
		fmt.Println("введите 2 число")
		fmt.Scan(&number)

		currentNumber = functions[functioName](currentNumber, number)
	}

	fmt.Println("твоё число")
	fmt.Println(currentNumber)
}
func divide(x, y int) int {
	return x / y
}
