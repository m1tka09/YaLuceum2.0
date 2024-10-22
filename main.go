package main

import "fmt"

func Calc(expression string) (float64, error) {
	// pa := expression[0] //получение первого символа
	// pb := expression[len(expression)-1]
	// az := expression[len(expression)-2]
	// if az != "*" || az != "/" || az != "+" || az != "-" { // обработка случая, если а-двузначное
	// 	pa = expression[1] // обработка случая, если а-двузначное
	// }
	// return

	for i, c := 0, len(expression); i < c; i++ {
		x := expression[i]
		if x == "+" {

		}
	}
}

func main() {
	var expression float64
	fmt.Scanln(&expression)
}
