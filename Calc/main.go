package main

import "fmt"

func main() {

	var number1, number2 float64
	var operator string

	fmt.Print("Введите число1 :")
	fmt.Scanln(&number1)

	fmt.Print("Введите число2 :")
	fmt.Scanln(&number2)

	fmt.Print("Введите оператор :")
	fmt.Scanln(&operator)

	switch operator {

	case "+":
		fmt.Printf("%.0f %s %.0f = %.0f", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%.0f %s %.0f = %.0f", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%.0f %s %.0f = %.0f", number1, operator, number2, number1*number2)

	case "/":
		if number2 == 0.0 {
			fmt.Println("Деление на 0")
		} else {
			fmt.Printf("%.0f %s %.0f = %.0f", number1, operator, number2, number1/number2)
		}
	default:
		fmt.Println("оператор неопознан")

	}
}
