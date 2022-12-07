package main

import (
	"fmt"
	"log"
	"strings"
)

var arabianDigits = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

var rimDigits = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "7": 7, "8": 8, "9": 9,
}

var operators = map[string]operate{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
	"/": func(x, y int) int { return x / y },
	"*": func(x, y int) int { return x * y },
}

type operate func(int, int) int

type Virag struct {
	X, Y      int
	Operation operate
}

func (v *Virag) FillVirag(sarr []string) (*Virag, []string) {
	for _, el := range sarr {
		_, ok := arabianDigits[el]
		//_, ok2 := rimDigits[el]
		if ok {
			v.X = arabianDigits[sarr[0]]
			v.Y = arabianDigits[sarr[2]]
		} else {
			v.Operation = operators[sarr[1]]
		} /*else if ok2 && !ok{
		  	v.X = rimDigits[sarr[0]]
		  	v.Y = rimDigits[sarr[2]]
		  }else if ok && ok2{
		  	log.Fatal("Can't operate Arbian and Roman digits")
		  	break
		  }*/
	}
	v.X = v.Operation(v.X, v.Y)
	sarr = sarr[3:]
	return v, sarr
}

func ParsingSequence(condition string) []string {
	stringArr := []string{}
	conditionArr := strings.Split(condition, "")

	for _, str := range conditionArr {
		if str != " " {
			stringArr = append(stringArr, str)
		}
	}
	return stringArr
}

func Operations(v *Virag, sarr []string) int {

	for _, el := range sarr {
		if len(sarr) >= 2 {

			_, ok := arabianDigits[el]
			_, ok2 := rimDigits[el]
			if ok && !ok2 {
				v.Y = arabianDigits[sarr[1]]
				v.X = v.Operation(v.X, v.Y)
				sarr = sarr[2:]

			} else {
				v.Operation = operators[sarr[0]]
			} /*else if ok2 && !ok {
			  	v.Y = rimDigits[sarr[1]]
			  	v.X = v.Operation(v.X, v.Y)
			      sarr = sarr[2:]
			  } else if ok && ok2{
			  	log.Fatal("Can't operate Arbian and Roman digits")
			  	break
			  }*/

		} else {
			log.Fatal("The sequence is empty")
			break
		}

	}

	return v.X
}

func main() {
	var str_input string
	fmt.Scanf("%s", &str_input)
	parsedSeq := ParsingSequence(str_input)
	viragenie := Virag{}
	fmt.Println(parsedSeq)
	completeVirag, seq := viragenie.FillVirag(parsedSeq)
	fmt.Println("Seq: ", seq)
	res := Operations(completeVirag, seq)
	fmt.Println("Result: ", res)

}
