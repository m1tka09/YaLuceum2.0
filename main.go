package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) { // Удаляем все пробелы из выражения expression = strings.ReplaceAll(expression, " ", "")

	// Проверяем правильность выражения
	if !isValid(expression) {
		return 0, fmt.Errorf("Invalid expression")
	}

	// Преобразуем выражение в обратную польскую запись
	postfix := infixToPostfix(expression)

	// Вычисляем результат выражения
	result, err := evaluatePostfix(postfix)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func isValid(expression string) bool { // Проверяем наличие недопустимых символов validChars := "0123456789+-*/()" for _, char := range expression { if !strings.Contains(validChars, string(char)) { return false } }

	// Проверяем правильность расстановки скобок
	openCount := strings.Count(expression, "(")
	closeCount := strings.Count(expression, ")")
	if openCount != closeCount {
		return false
	}

	return true
}

func infixToPostfix(expression string) []string {
	var postfix []string
	var stack []string

	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	tokens := strings.Split(expression, "")

	for _, token := range tokens {
		if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for stack[len(stack)-1] != "(" {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		} else if _, isOperator := precedence[token]; isOperator {
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else {
			postfix = append(postfix, token)
		}
	}

	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix
}

func evaluatePostfix(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("Invalid expression")
			}

			num2 := stack[len(stack)-1]
			num1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("Invalid expression")
	}

	return stack[0], nil
}

// func main() {
// 	result, err := Calc("3 + 4 * 2 / (1 - 5)^2")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Result:", result)
// 	}
// }
