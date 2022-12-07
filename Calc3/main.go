package main

import (
	"fmt"
	"strings"
)

// Structure of operands and the operator
type Expression struct {
	X, Y      int
	Operation Operate
}

// Filling Expression structure
func (exp *Expression) FillingExpression(stringarr []string) (*Expression, []string) {

	for _, elem := range stringarr {
		_, ok := singledigits[elem]

		if ok {
			exp.X = singledigits[stringarr[0]]
			exp.Y = singledigits[stringarr[2]]
		} else {
			exp.Operation = operators[stringarr[1]]
		}

	}
	exp.X = exp.Operation(exp.X, exp.Y)

	stringarr = stringarr[3:]

	return exp, stringarr
}

// Operation function
type Operate func(int, int) int

// Map of operators "+" "-" and funcs
var operators = map[string]Operate{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
}

// Map of single digits
var singledigits = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

// Preparing input condition with trim spaces
func PreparingInputSequence(condition string) []string {
	stringArr := []string{}
	conditionArr := strings.Split(condition, "")

	for _, str := range conditionArr {
		if str != " " {
			stringArr = append(stringArr, str)
		}
	}
	return stringArr
}

//  Operations with filled structures
func Operations(exp *Expression, stringarr []string) int {

	for _, elem := range stringarr {
		if len(stringarr) >= 2 {

			_, ok := singledigits[elem]
			if ok {
				exp.Y = singledigits[stringarr[1]]

				exp.X = exp.Operation(exp.X, exp.Y)
				stringarr = stringarr[2:]

			} else {
				exp.Operation = operators[stringarr[0]]
			}

		} else {
			fmt.Println("The sequence is empty")
			break
		}

	}

	return exp.X
}

func main() {

	// First input
	firstInput := "1+1"
	preparedSequence := PreparingInputSequence(firstInput)
	fmt.Println("Prepared first sequence: ", PreparingInputSequence(firstInput)) // [1 + 1]

	firstExpression := Expression{}

	completeFirstExpression, firstSeq := firstExpression.FillingExpression(preparedSequence)

	resultOfFirstInput := Operations(completeFirstExpression, firstSeq)

	fmt.Println("Result of first input: ", resultOfFirstInput) // Output 2

}
