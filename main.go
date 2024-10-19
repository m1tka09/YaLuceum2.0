package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	for i := num; i < num+10; i++ {
		fmt.Println(fibonacci(i))
	}
}
