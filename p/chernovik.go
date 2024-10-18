//  package main

// import "fmt"

// func main() {
// 	var i int
// 	fmt.Println("Введите число:")
//     fmt.Scanln(&i)
// 	switch {
// 	case i > 0:
// 		fmt.Println("Число положительное")
// 	case i < 0:
// 		fmt.Println("Число отрицательное")
// 	default:
// 		fmt.Println("Введён ноль")
// 	}}

// var i, a int
// 	fmt.Scanln(&i, &a)

// 	if i > a {

// 		fmt.Println("Первое число больше второго")
// 	}
// 	if i < a {
// 		fmt.Println("Второе число больше первого")
// 	}
// 	if i == a {
// 		fmt.Println("Числа равны")
// 	}

//var i int
	// fmt.Scanln(&i)
	// if i > 0 && i%2 == 0 {
	// 	fmt.Println("Число положительное и четное")
	// }

	// if i > 0 && i%2 != 0 {
	// 	fmt.Println("Число положительное и нечетное")
	// }

	// if i < 0 && i%2 == 0 {
	// 	fmt.Println("Число отрицательное и четное")
	// }

	// if i < 0 && i%2 != 0 {
	// 	fmt.Println("Число отрицательное и нечетное")
	// }

	// if i == 0 {
	// 	fmt.Println("Введен ноль")
	// }

	//var i, a int
	// fmt.Scanln(&i, &a)
	// if i > 0 && a > 0 {
	// 	fmt.Println("Оба числа положительные")
	// }

	// if i < 0 && a < 0 {
	// 	fmt.Println("Оба числа отрицательные")
	// }

	// if i > 0 && a < 0 || i < 0 && a > 0 {
	// 	fmt.Println("Одно число положительное, а другое отрицательное")
	// }

	// if i == 0 && a != 0 || i != 0 && a == 0 {
	// 	fmt.Println("Одно из чисел равно нулю")
	// }
	

// 	package main

// import "fmt"

// func main() {
// 	var i, a, c int
// 	fmt.Scanln(&i, &a, &c)
// 	if i == a && i == c && a == c {
// 		fmt.Println("Все числа равны")
// 	}

// 	if i == a && !(a == c) && !(i == c) || i == c && !(a == c) && !(i == a) || a == c && !(a == i) && !(i == c) {
// 		fmt.Println("Два числа равны")
// 	}

// 	if i != a && i != c && a != c {
// 		fmt.Println("Все числа разные")
// 		// !(i == a) || !(i == c) || !(a == c)
// 	}
// 	if !((i == a && i == c && a == c) && (i == a && !(a == c) && !(i == c) || i == c && !(a == c) && !(i == a) || a == c && !(a == i) && !(i == c)) && (i != a && i != c && a != c)) {
// 		fmt.Println("Некорректный ввод")
// 	}
// }

// if i < 10 && i >= 0 {
// 	fmt.Println("Число меньше 10")
// }

// if i < 100 && i >= 10 && i > 0 {
// 	fmt.Println("Число меньше 100")
// }

// if i < 1000 && i > 10 && i >= 100 && i > 0 {
// 	fmt.Println("Число меньше 1000")
// }

// if i >= 1000 && i > 0 {
// 	fmt.Println("Число больше или равно 1000")
// }

// if i < 0 {
// 	fmt.Println("Некорректный ввод")
// }
//if (i > 0 && a > 0 && c > 0) || (i < 0 && a > 0 && c > 0) || (i > 0 && a < 0 && c > 0) || (i > 0 && a > 0 && c < 0) || (i < 0 && a < 0 && c < 0) || (i < 0 && a < 0 && c > 0) {
	// 	if i == a && i == c && a == c {
	// 		fmt.Println("Все числа равны")
	// 	}

	// 	if i == a && !(a == c) && !(i == c) || i == c && !(a == c) && !(i == a) || a == c && !(a == i) && !(i == c) {
	// 		fmt.Println("Два числа равны")

	// 	}

	// 	if i != a && i != c && a != c {
	// 		fmt.Println("Все числа разные")
	// 		// !(i == a) || !(i == c) || !(a == c)
	// 	}
	// } else {
	// 	fmt.Println("Некорректный ввод")
	// }

// 	package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var a, b, c int
// 	fmt.Println("ax**2+bx+c")
// 	fmt.Println("a=")
// 	fmt.Scanln(&a)
// 	fmt.Println("b=")
// 	fmt.Scanln(&b)
// 	fmt.Println("c=")
// 	fmt.Scanln(&c)
// 	var n float64
// 	x := math.Pow(n, 2)
// 	a*x + b*x + c = 0

// }

// for i, letter := range "Hello, world!" {
	// 	fmt.Println(i, string(letter))
	// }


	// var a int
	// fmt.Scanln(&a)
	// for i := 1; i < 11; i++ {
	// 	res := a * i
	// 	fmt.Println(a, "*", i, "=", res)
	// }

// 	package main

// // fmt package provides the function to print anything
// import "fmt"

// func main() {

// 	// declaring the integer number using the var keyword
// 	// whose factorial we have to find
// 	var number, iterator int64 // initializing the variable whose factorial we want to find  number = 9
// fmt.Scanln(&number)
// 	// declaring the answer variable of int64 type and initializing with 1
// 	var answer int64 = 1

// 	// Running the for loop to find the factorial
// 	for iterator = 1; iterator <= number; iterator++ {
// 		answer = answer * iterator
// 	}

// 	// Printing the factorial of the respective number
// 	fmt.Println(answer)
// }

// var sum, i int
	// if i < 0 {
	// 		fmt.Println("Некорректный ввод")
	// 	}
	// for i := 20; i <= 530; i++ {
	// 	if i%6 == 0 && i%8 != 0 {
	// 		sum += i
	// 	}
	// }
	// fmt.Println(sum)

	// var a int
	// fmt.Scanln(&a)
	
	// var sum int
	// for i := 0; i <= a; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	sum += i
	// }
	// if a < 0 {
	// 	fmt.Println("Некорректный ввод")
	// }
	// if a > 0 {
	// 	fmt.Println(sum)
	// }

	// var  a int
	// fmt.Scanln(&a)
	// for i := 1; i <= a; i++ {
	// 	if i%3 == 0 {
	// 		// for {
	// 			fmt.Println(i)
	// 		// }
	// 	}
	// }

	// package main

// import "fmt"

// func main() {
// 	var sum, a int
// 	fmt.Scanln(&a)
// for i := 1; i <= a; i++ {
// 	if i % 3 != 0 && i % 5 != 0 {
// 		sum += i
// 	}
// }
// fmt.Println(sum)
// }

// func fibonacciIterative(n int) int {
// 	if n <= 1 {
// 	  return n
// 	}
// 	var n2, n1 = 0, 1
// 	for i := 2; i <= n; i++ {
// 	  n2, n1 = n1, n1 + n2
// 	}
// 	return n1
//   }
  
//   func main() {
// 	fmt.Println(fibonacciIterative(9))
  
  // package main 2

// import (
// 	"fmt"
// 	"math"
// )

// func SqRoots () {
// 	var a, b, c float64

// }

// func main() {
// 	if a < 0 {
// 		fmt.Println("Некорректный ввод")
// 	}
// }

// // package main

// // import (
// // 	"fmt"
// // 	"math"
// // )

// // // Main function
// // func main() {

// // 	var a, b, c float64
// // 	if a<0 {
// // 		fmt.Println("Некоррректный ввод")
// // 	}
// // 	if a>=0 {
// // 		fmt.Scanln(&a, &b, &c)
// // 	sqa := math.Sqrt(a)
// // 	sqb := math.Sqrt(b)
// // 	sqc := math.Sqrt(c)

// // 	fmt.Println(sqa, sqb, sqc)
// // 	}

// // }

// package main 

// import (
// 	"fmt"
// 	// "math"
// )

// func fib(n int) int {
//     // крайний случай, так как при f(1) и так далее мы не сможем посчитать обе следующих итерации
//     if n < 2 {
//        return n
//     }

//     return fib(n - 1) + fib(n - 2)
// }

// func main() {
//     fmt.Println(fib(6)) // Вывод: 8
// }
// package main

// import (
// 	"fmt"
// 	"math"
// 	// "os"
// )

// func findDiscriminant(a, b, c float64) float64 {
// 	discriminant := math.Pow(b, 2) - 4.0*a*c

// 	// возвращаем значения вычисления из findDiscriminant
// 	return discriminant
// }

// func SqRoot() {
// 	var a, b, c float64
// 	fmt.Scanln(&a, &b, &c)
// 	D := findDiscriminant(a, b, c)
// 	RootD := math.Sqrt(D)
// 	var x1 float64 = (-b - RootD) / 2 * a
// 	var x2 float64 = (-b + RootD) / 2 * a
// 	if a < 0 {
// 		fmt.Println("Некорректный ввод")
// 	} else {
// 		if x1 == x2 {
// 			fmt.Println(x1)
// 		}
// 		if x1 != x1 && x2 != x2 {
// 			fmt.Println(0, 0)

// 		} else {
// 			fmt.Println(x1, x2)
// 		}
// 	}
// }

// func main() {
// 	SqRoot()
// }

// package main

// import (
// 	"fmt"
// 	// "math"
// )

// func Add(a, b float64) float64 {
// 	c := a + b
// 	fmt.Println(c)
// 	return c
// }

// func Multiply(a, b float64) float64 {
// 	c := a * b
// 	fmt.Println(c)
// 	return c
// }

// func PrintNumbersAscending(n int) {
// 	i := 1
// 	for i <= n {
// 		fmt.Println(i)
// 		i++
// 	}
// }

// func main() {
// 	var a, b float64
// 	var n int
// 	fmt.Scanln(&a, &b, &n)
// 	Add(a, b)
// 	Multiply(a, b)
// 	PrintNumbersAscending(n)
// }
package main
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var n int
// 	fmt.Scanln(&n)
// 	var i int
// 	for i = 1; i < factorial(n); i++ {
// fmt.Println(i)
// 	}
// 	fmt.Println(factorial(n))
// }

// func factorial(n int) int {
// 	if n == 1 {
// 		return 1
// 	}

// 	return n * factorial(n-1)
// }

// // package main

// // import (
// // 	"fmt"
// // 	"math"
// // )

// // // func findDiscriminant(a, b, c float64) float64 {
// // // 	discriminant := math.Pow(b, 2) - 4.0*a*c

// // // 	// возвращаем значения вычисления из findDiscriminant
// // // 	return discriminant
// // // }

// // func SqRoot() {
// // 	var a, b, c float64
// // 	fmt.Scanln(&a, &b, &c)
// // 	D := math.Pow(b, 2) - 4.0*a*c
// // 	// D := findDiscriminant(a, b, c)
// // 	RootD := math.Sqrt(D)
// // 	var x1 float64 = (-b - RootD) / 2 * a
// // 	var x2 float64 = (-b + RootD) / 2 * a

// // 	if a < 0 {
// // 		fmt.Println("Некорректный ввод")
// // 	} else {
// // 		if x1 == x2 {
// // 			fmt.Println(x1)
// // 		}
// // 		if x1 != x1 && x2 != 2 {
// // 			fmt.Println(0, 0)

// // 		} else {
// // 			fmt.Println(x1, x2)
// // 		}
// // 	}
// // }

// // func main() {
// // 	SqRoot()
// // }

// package main

// import (
//     "fmt"
// 	"os"
//     "strconv"
//     //more1
// )

// func main() {
// 	if len(os.Args) != 3 {
//     	fmt.Println("Two integer parameters expected")
//     	os.Exit(1)
// 	}
// 	a, err := strconv.Atoi(os.Args[1])
// 	if err != nil {
//     	fmt.Println("invalid first argument:", err)
//     	os.Exit(1)
// 	}
// 	b, err := strconv.Atoi(os.Args[2])
// 	if err != nil {
//     	fmt.Println("invalid second argument:", err)
//     	os.Exit(1)
// 	}
// 	func adder(a, b int) (int, error) {
// 		x := a + b
// 		if (x > a) != (b > 0) {
// 			return x, errors.New("addition out of bounds")
// 		}
// 		return x, nil
// 	}
// 	} else {
//     	fmt.Println(c)
// 	}
// 	d, err := subtractor(a, b)
// 	if err != nil {
//     	fmt.Println("subtracting failed:", err)
// 	} else {
//     	fmt.Println(d)
// 	}
// 	e, err := multiplier(a, b)
// 	if err != nil {
//     	fmt.Println("multiplying failed:", err)
// 	} else {
//     	fmt.Println(e)
// 	}
// 	f, err := divider(a, b)
// 	if err != nil {
//     	fmt.Println("dividing failed:", err)
// 	} else {
//     	fmt.Println(f)
// 	}

// }





























// // func Add(a, b float64) float64 {var "a+b" string


// // func Add () {}

// // func Multiply () {}

// // func PrintNumbersAscending () {}

// // func main () {

// // }


// // 	c := a + b
// // 	fmt.Println(c)
// // 	return c
// // }

// // func Multiply(a, b float64) float64 {
// // 	c := a * b
// // 	fmt.Println(c)
// // 	return c
// // }

// // func PrintNumbersAscending(n int) {
// // 	i := 1
// // 	for i <= n {
// // 		fmt.Println(i)
// // 		i++
// // 	}
// // }

// // func main() {
// // 	var a, b float64
// // 	var n int
// // 	fmt.Scanln(&a, &b, &n)
// // 	Add(a, b)
// // 	Multiply(a, b)
// // 	PrintNumbersAscending(n)
// // }

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//


// // // // // // package main

// // // // // // import (
// // // // // // 	"slices"
// // // // // // 	"testing"
// // // // // // )

// // // // // // func TestUnderLimit(t *testing.T) {
// // // // // // 	type test struct {
// // // // // // 		nums      []int
// // // // // // 		n         int
// // // // // // 		limit     int
// // // // // // 		expected  []int
// // // // // // 		wantError bool
// // // // // // 	}

// // // // // // 	tests := []test{
// // // // // // 		{
// // // // // // 			nums:      []int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12},
// // // // // // 			n:         5,
// // // // // // 			limit:     3,
// // // // // // 			expected:  []int{2, 0, 2},
// // // // // // 			wantError: false,
// // // // // // 		},
// // // // // // 		{
// // // // // // 			nums:      nil,
// // // // // // 			wantError: true,
// // // // // // 		},
// // // // // // 		{
// // // // // // 			nums:      []int{},
// // // // // // 			n:         5,
// // // // // // 			limit:     3,
// // // // // // 			expected:  []int{},
// // // // // // 			wantError: false,
// // // // // // 		},
// // // // // // 		{
// // // // // // 			nums:      []int{3, 5, 6},
// // // // // // 			n:         5,
// // // // // // 			limit:     10,
// // // // // // 			expected:  []int{3, 5, 6},
// // // // // // 			wantError: false,
// // // // // // 		},
// // // // // // 		{
// // // // // // 			nums:      []int{-13, 0, 6},
// // // // // // 			n:         1,
// // // // // // 			limit:     -5,
// // // // // // 			expected:  []int{-13},
// // // // // // 			wantError: false,
// // // // // // 		},
// // // // // // 		{
// // // // // // 			nums:      []int{},
// // // // // // 			n:         -1,
// // // // // // 			limit:     5,
// // // // // // 			wantError: true,
// // // // // // 		},
// // // // // // 	}

// // // // // // 	for _, tc := range tests {
// // // // // // 		result, err := UnderLimit(tc.nums, tc.limit, tc.n)
// // // // // // 		if tc.wantError {
// // // // // // 			if err == nil {
// // // // // // 				t.Fatalf("expec... File is too long to be displayed fully

// // // // // // // package main

// // // // // // // import (
// // // // // // // 	"fmt"
// // // // // // // 	"math"
// // // // // // // 	// "os"
// // // // // // // )

// // // // // // // func findDiscriminant(a, b, c float64) float64 {
// // // // // // // 	discriminant := math.Pow(b, 2) - 4.0*a*c

// // // // // // // 	// возвращаем значения вычисления из findDiscriminant
// // // // // // // 	return discriminant
// // // // // // // }

// // // // // // // func SqRoot() {
// // // // // // // 	var a, b, c float64
// // // // // // // 	fmt.Scanln(&a, &b, &c)
// // // // // // // 	D := findDiscriminant(a, b, c)
// // // // // // // 	RootD := math.Sqrt(D)
// // // // // // // 	var x1 float64 = (-b - RootD) / 2 * a
// // // // // // // 	var x2 float64 = (-b + RootD) / 2 * a
// // // // // // // 	if a < 0 {
// // // // // // // 		fmt.Println("Некорректный ввод")
// // // // // // // 	} else {
// // // // // // // 		if x1 == x2 {
// // // // // // // 			fmt.Println(x1)
// // // // // // // 		}
// // // // // // // 		if x1 != x1 && x2 != 2 {
// // // // // // // 			fmt.Println(0, 0)

// // // // // // // 		} else {
// // // // // // // 			fmt.Println(x1, x2)
// // // // // // // 		}
// // // // // // // 	}
// // // // // // // }

// // // // // // // func main() {
// // // // // // // 	SqRoot()
// // // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	// "math"
// // // // // // )

// // // // // // func Add(a, b float64) float64 {
// // // // // // 	c := a + b
// // // // // // 	fmt.Println(c)
// // // // // // 	return c
// // // // // // }

// // // // // // func Multiply(a, b float64) float64 {
// // // // // // 	c := a * b
// // // // // // 	fmt.Println(c)
// // // // // // 	return c
// // // // // // }

// // // // // // func PrintNumbersAscending(n int) {
// // // // // // 	i := 1
// // // // // // 	for i <= n {
// // // // // // 		fmt.Print(i, " ")
// // // // // // 		i++
// // // // // // 	}
// // // // // // }

// // // // // // func main() {
// // // // // // 	var a, b float64
// // // // // // 	var n int
// // // // // // 	fmt.Scanln(&a, &b, &n)
// // // // // // 	Add(a, b)
// // // // // // 	Multiply(a, b)
// // // // // // 	PrintNumbersAscending(n)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var min int
// // // // // //     fmt.Scan(&min)
// // // // // // 	h := min/60
// // // // // // 	m:= min%60
// // // // // // 	fmt.Println(min, "- это", h, "час", m, "минут.")
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // //     a := -3
// // // // // //     b := 2
// // // // // //     R := f(a)
// // // // // //     for t := a; t <= b; t++ {
// // // // // //         if f(t) > R {
// // // // // //             R = f(t)
// // // // // //         }
// // // // // //     }
// // // // // //     fmt.Println(R)
// // // // // // }

// // // // // // func f(x int) int {
// // // // // //     return (x + 5) * (x + 3)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	i := 12
// // // // // // 	var k int
// // // // // // 	fmt.Scan(&k)
// // // // // // 	for i > 0 && f(i) >= k {
// // // // // // 		i--
// // // // // // 	}
// // // // // // 	fmt.Println(i)
// // // // // // }

// // // // // // func f(n int) int {
// // // // // // 	return n*n + 20
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // //     var n, k int
// // // // // //     fmt.Scan(&n, &k)

// // // // // //     factN := factorial(n)
// // // // // //     factK := factorial(k)
// // // // // //     nSubtractK := n - k
// // // // // //     factNSubtractK := factorial(nSubtractK) // можно factorial(n - k), и тогда строчка выше не нужна

// // // // // //     fmt.Println(factN / (factK * factNSubtractK))
// // // // // // }

// // // // // // func factorial(number int) int {
// // // // // //     fact := 1
// // // // // //     for i := 1; i <= number; i++ {
// // // // // //         fact = fact * i
// // // // // //     }
// // // // // //     return fact
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a, b, c, d int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	fmt.Scan(&b)
// // // // // // 	fmt.Scan(&c)
// // // // // // 	fmt.Scan(&d)
// // // // // // 	sum := a + b + c + d
// // // // // // 	sumn := sum * 3
// // // // // // 	fmt.Println(sumn)
// // // // // // }

// // // // // // package main

// // // // // // func main () {
// // // // // // 	nraz()
// // // // // // }

// // // // // // func nraz (a int, b int) {
// // // // // //   for i := 0; a<1000000000; i++ {
// // // // // // 	a1 := a/10
// // // // // // 	a2 := (a/10)/10
// // // // // // 	b1 := b/10
// // // // // //   }
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scanln(&a)
// // // // // // 	b := a + 1
// // // // // // 	c := a - 1
// // // // // // 	fmt.Println("Следующее за числом", a, "число:", b)
// // // // // // 	fmt.Println("Для числа", a, "предыдущее число:", c)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	x := 432
// // // // // // 	y := x / 100
// // // // // // 	x = (x % 100) * 10
// // // // // // 	x = x + y
// // // // // // 	fmt.Println(x)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scanln(&a)
// // // // // // 	i := a
// // // // // // 	c:=i%10
// // // // // // 	for c < 10 {
// // // // // // 		b := i / 10

// // // // // // 	fmt.Println(b)
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"errors"
// // // // // // 	"fmt"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var a, b, c int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	b = a % 100
// // // // // // 	c = b / 10
// // // // // // 	if a <= 10000 {
// // // // // // 		fmt.Println(c)
// // // // // // 	} else {
// // // // // // 		errors.New("Error")
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	b := a / 100
// // // // // // 	c := (a % 100) / 10
// // // // // // 	d := a % 10
// // // // // // 	e := b+c+d
// // // // // // 	fmt.Println(e)
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"strconv"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var str string = "5"
// // // // // // 	num, _ := strconv.Atoi(str)
// // // // // // 	fmt.Println(num * 2) // вывод 10
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	b := a / 100
// // // // // // 	c := (a % 100) / 10
// // // // // // 	d := a % 10
// // // // // // 	fmt.Println(d, c, b)
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // )

// // // // // // var integer int64 = 32500
// // // // // // var floatNum float64 = 22000.456

// // // // // // func main() {

// // // // // // 	// Обычный способ вывода десятичного числа
// // // // // // 	fmt.Printf("%d \n", integer)

// // // // // // 	// Всегда показывает знак
// // // // // // 	fmt.Printf("%+d \n", integer)

// // // // // // 	// Вывод с другим основанием x -16, o-8, b -2, d - 10
// // // // // // 	fmt.Printf("%x \n", integer)
// // // // // // 	fmt.Printf("%#x \n", integer)

// // // // // // 	// Отступ перед нулями
// // // // // // 	fmt.Printf("%010d \n", integer)

// // // // // // 	// Оставляет отступ с пробелами
// // // // // // 	fmt.Printf("% 10d \n", integer)

// // // // // // 	// Отступ с правой стороны
// // // // // // 	fmt.Printf("% -10d \n", integer)

// // // // // // 	// Вывод вещественного значения
// // // // // // 	// с плавающей запятой
// // // // // // 	fmt.Printf("%f \n", floatNum)

// // // // // // 	// Вещественное число
// // // // // // 	// с ограниченной точностью = 5 (после запятой)
// // // // // // 	fmt.Printf("%.5f \n", floatNum)

// // // // // // 	// Вещественное число
// // // // // // 	// в научной заметке
// // // // // // 	fmt.Printf("%e \n", floatNum)

// // // // // // 	// Вещественное число
// // // // // // 	// %e для крупной экспоненты
// // // // // // 	// или %f в противном случае
// // // // // // 	fmt.Printf("%g \n", floatNum)

// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"strconv"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	b := a / 100
// // // // // // 	c := (a % 100) / 10
// // // // // // 	d := a % 10

// // // // // // 	var numb int = b
// // // // // //    strb := strconv.Itoa(numb)
// // // // // // //    fmt.Println(strb) // вывод 55
// // // // // //    var numc int = c
// // // // // //    strc := strconv.Itoa(numc)
// // // // // // //    fmt.Println(strc) // вывод 55
// // // // // //    var numd int = d
// // // // // //    strd := strconv.Itoa(numd)
// // // // // // //    fmt.Println(strd) // вывод 55
// // // // // // fmt.Println(strd+strc+strb)
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"strconv"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	b := a / 100
// // // // // // 	c := (a % 100) / 10
// // // // // // 	d := a % 10
// // // // // // 	fmt.Println(conv(d) + conv(c) + conv(b))
// // // // // // }

// // // // // // func conv(num int) string {
// // // // // // 	str := strconv.Itoa(num)
// // // // // // 	return str
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a, b float32
// // // // // // 	fmt.Scanln(&a)
// // // // // // 	fmt.Scanln(&b)
// // // // // // 	s := (a * b) / 2
// // // // // // 	fmt.Println(s)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var F float32
// // // // // // 	fmt.Scan(&F)
// // // // // // 	C := (F - 32.0) * (5.0 / 9.0)
// // // // // // 	fmt.Println(C)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a, b float64
// // // // // // 	fmt.Scan(&a, &b)
// // // // // // 	var c = (a + b) / 2
// // // // // // 	var d = float64(c)
// // // // // // 	fmt.Println(d)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a float32
// // // // // // 	fmt.Scan(&a)
// // // // // // 	b := float32(int(a))
// // // // // // 	c := a-b
// // // // // // 	fmt.Println(c)
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"math"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var b float64
// // // // // // 	fmt.Scan(&b)
// // // // // // 	kb := b / (math.Pow(2, 13))
// // // // // // 	fmt.Print(kb)
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"math"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var a, b float64
// // // // // // 	fmt.Scan(&a, &b)
// // // // // // 	c := math.Sqrt((math.Pow(a, 2) + math.Pow(b, 2)))
// // // // // // 	P := a + b + c
// // // // // // 	fmt.Println(P)
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"math"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var x1, y1, x2, y2 float64
// // // // // // 	fmt.Scanln(&x1)
// // // // // // 	fmt.Scanln(&y1)
// // // // // // 	fmt.Scanln(&x2)
// // // // // // 	fmt.Scanln(&y2)
// // // // // // 	s := math.Sqrt((math.Pow((x1 - x2), 2)) + (math.Pow((y1 - y2), 2)))
// // // // // // 	fmt.Println(s)
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	// "math"
// // // // // // )

// // // // // // func main() {
// // // // // // 	// var x1, y1, x2, y2 float64
// // // // // // 	// fmt.Scanln(&x1)
// // // // // // 	// fmt.Scanln(&y1)
// // // // // // 	// fmt.Scanln(&x2)
// // // // // // 	// fmt.Scanln(&y2)
// // // // // // 	// s := math.Abs(x1-x2) + math.Abs(y1-y2)
// // // // // // 	// fmt.Println(s)
// // // // // // 	m := 67
// // // // // // 	m = m + 13
// // // // // // 	var n int = m/4 - m/2
// // // // // // 	var c = m - n
// // // // // // 	fmt.Println(c)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a, b, c float64
// // // // // // 	fmt.Scan(&a)
// // // // // // 	fmt.Scan(&b)
// // // // // // 	fmt.Scan(&c)
// // // // // // 	d := a+b+c
// // // // // // 	e := a*b*c
// // // // // // 	fmt.Println(d, e)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	b := a / 100
// // // // // // 	if b < 10 {
// // // // // // 		fmt.Println(b)
// // // // // // 	} else {
// // // // // // 		c := b % 10
// // // // // // 		fmt.Println(c)
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	// "strconv"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	num := a % 1000
// // // // // // 	b := num / 100
// // // // // // 	c := (num % 100) / 10
// // // // // // 	d := num % 10
// // // // // // 	e := b+c+d
// // // // // // 	fmt.Println(e)
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"math"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var a, b float64
// // // // // // 	fmt.Scan(&a, &b)
// // // // // // 	c := math.Sqrt((math.Pow(a, 2) + math.Pow(b, 2)))
// // // // // // 	fmt.Println(c)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main(){
// // // // // //     var x int
// // // // // //     fmt.Scan(&x)

// // // // // //     fmt.Print(1)

// // // // // //     if x != 5 {
// // // // // //         fmt.Print(2)
// // // // // //     }

// // // // // //     fmt.Print(3)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main(){
// // // // // //     var x int
// // // // // //     fmt.Scan(&x)

// // // // // //     fmt.Print(1)

// // // // // //     if x != 5 {
// // // // // //         fmt.Print(2)
// // // // // //     }
// // // // // //     fmt.Print(3)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main(){
// // // // // //     var x int
// // // // // //     fmt.Scan(&x)

// // // // // //     fmt.Print(1)

// // // // // //     if x % 10 <= 7{
// // // // // //         fmt.Print(2)
// // // // // //     } else {
// // // // // //         fmt.Print(3)
// // // // // //     }
// // // // // //     fmt.Print(4)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main(){
// // // // // //     var x1, x2 int
// // // // // //     fmt.Scan(&x1, &x2)

// // // // // //     if x1 <= 2 * x2 {
// // // // // //         fmt.Print("ДА")
// // // // // //     } else {
// // // // // //         fmt.Print("НЕТ")
// // // // // //     }
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main(){
// // // // // //     var x int
// // // // // //     fmt.Scan(&x)

// // // // // //     if x % 2 == 0 {
// // // // // //         a := 5
// // // // // //     } else {
// // // // // //         b := 10
// // // // // //     }

// // // // // //     fmt.Print(a + b)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a, b int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	fmt.Scan(&b)
// // // // // // 	if a==b {
// // // // // // 		fmt.Print(a)
// // // // // // 	}

// // // // // // 	if a<b {
// // // // // // 		fmt.Print(b)
// // // // // // 	}

// // // // // // 	if a>b {
// // // // // // 		fmt.Print(a)
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a, b string
// // // // // // 	fmt.Scan(&a)
// // // // // // 	fmt.Scan(&b)
// // // // // // 	if a == b {
// // // // // // 		fmt.Println("Пароль принят")
// // // // // // 	} else {
// // // // // // 		fmt.Println("Пароль не принят")

// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	if a>=12 {
// // // // // // 		fmt.Println("Доступ разрешен")
// // // // // // 	} else  {
// // // // // // 		fmt.Println("Доступ запрещён")

// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	if a%2==0 {
// // // // // // 		fmt.Println("YES")
// // // // // // 	} else {
// // // // // // 		fmt.Println("NO")

// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a, b int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	fmt.Scan(&b)
// // // // // // 	if a%b == 0 {
// // // // // // 		fmt.Println("YES")
// // // // // // 	} else {
// // // // // // 		fmt.Println("NO")

// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var x1, x2 int
// // // // // // 	fmt.Scan(&x1, &x2)

// // // // // // 	if x1 < x2 {
// // // // // // 		fmt.Print(1)
// // // // // // 	}
// // // // // // 	if x1 > x2 {
// // // // // // 		fmt.Print(2)
// // // // // // 	}
// // // // // // 	if x1 == x2 {
// // // // // // 		fmt.Print(3)
// // // // // // 	}
// // // // // // 	if x1 != x2 {
// // // // // // 		fmt.Print(4)
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	if a < 0 {
// // // // // // 		fmt.Println(-1)
// // // // // // 	}

// // // // // // 	if a > 0 {
// // // // // // 		fmt.Println(1)
// // // // // // 	}

// // // // // // 	if a == 0 {
// // // // // // 		fmt.Println(0)
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	b := a / 100
// // // // // // 	c := a / 1000
// // // // // // 	// c := (a % 100) / 10
// // // // // // 	// d := a % 10
// // // // // // 	if b > 0 && c == 0 {
// // // // // // 		fmt.Println("YES")
// // // // // // 	} else {
// // // // // // 		fmt.Println("NO")
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	if -1 < a && a < 17 {
// // // // // // 		fmt.Println("Принадлежит")
// // // // // // 	} else {
// // // // // // 		fmt.Println("Не принадлежит")
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	if a <= -3 || a >= 7 {
// // // // // // 		fmt.Println("Принадлежит")
// // // // // // 	} else {
// // // // // // 		fmt.Println("Не принадлежит")
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	b := a / 100
// // // // // // 	c := (a % 100) / 10
// // // // // // 	d := a % 10
// // // // // // 	// !(b == c && d == c && b == d) && !((!(b==c && d==c) || !(b==c && b==d) || !(d==c && b==d)))
// // // // // // 	if b != c && d != c && b != d {
// // // // // // 		fmt.Println("YES")
// // // // // // 	} else {
// // // // // // 		fmt.Println("NO")
// // // // // // 	}
// // // // // // 	// if !(b == c && d == c) || !(b == c && b == d) || !(d == c && b == d) {
// // // // // // 	// 	fmt.Println("NO")
// // // // // // 	// }
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	b := (a / 100000) % 10
// // // // // // 	c := (a / 10000) % 10
// // // // // // 	d := (a / 1000) % 10
// // // // // // 	e := (a / 100) % 10
// // // // // // 	f := (a % 100) / 10
// // // // // // 	g := a % 10
// // // // // // 	sum1 := b+c+d
// // // // // // 	sum2 := e+f+g
// // // // // // 	if sum1 == sum2 {
// // // // // // 		fmt.Println("YES")
// // // // // // 	} else {
// // // // // // 		fmt.Println("NO")
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var y int
// // // // // // 	fmt.Scan(&y)
// // // // // // 	if (y%4 == 0 && !(y%100 == 0)) || y%400 == 0 {
// // // // // // 		fmt.Println("YES")
// // // // // // 	} else {
// // // // // // 		fmt.Println("NO")
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var y int
// // // // // // 	fmt.Scan(&y)
// // // // // // 	if y<=13 {
// // // // // // 		fmt.Println("детство")
// // // // // // 	}

// // // // // // 	if y>=14 && y<=24 {
// // // // // // 		fmt.Println("молодость")
// // // // // // 	}

// // // // // // 	if y>=25 && y<=59 {
// // // // // // 		fmt.Println("зрелость")
// // // // // // 	}

// // // // // // 	if y>=60 {
// // // // // // 		fmt.Println("старость")
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var nl, sl, n, s int
// // // // // // 	fmt.Scan(&nl)
// // // // // // 	fmt.Scan(&sl)
// // // // // // 	fmt.Scan(&n)
// // // // // // 	fmt.Scan(&s)

// // // // // // 	if nl == n || sl == s {
// // // // // // 		fmt.Println("YES")
// // // // // // 	} else {
// // // // // // 		fmt.Println("NO")
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"math"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var xl, yl, x, y float64
// // // // // // 	fmt.Scan(&xl)
// // // // // // 	fmt.Scan(&yl)
// // // // // // 	fmt.Scan(&x)
// // // // // // 	fmt.Scan(&y)

// // // // // // 	if math.Abs(x-xl)==math.Abs(y-yl) {
// // // // // // 		fmt.Println("YES")
// // // // // // 	} else {
// // // // // // 		fmt.Println("NO")
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var x int
// // // // // // 	fmt.Scan(&x)
// // // // // // 	if x <= 10000 {
// // // // // // 		if x >= -3 && x <= 1 || x >= 5 && x <= 10 {
// // // // // // 			fmt.Println("YES")
// // // // // // 		} else {
// // // // // // 			fmt.Println("NO")
// // // // // // 		}
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main(){
// // // // // //     var x int
// // // // // //     fmt.Scan(&x)
// // // // // //     if x >= -3 && x < 2 || x > 4 && x < 10 {
// // // // // //         fmt.Println("YES")
// // // // // //     } else {
// // // // // //         fmt.Println("NO")
// // // // // //     }
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var w int
// // // // // // 	fmt.Scan(&w)
// // // // // // 	if w < 60 {
// // // // // // 		fmt.Println("Легкий вес")
// // // // // // 	}

// // // // // // 	if w > 60 && w < 64 {
// // // // // // 		fmt.Println("Первый полусредний вес")
// // // // // // 	}

// // // // // // 	if w > 64 && w < 69 {
// // // // // // 		fmt.Println("Полусредний вес")
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // //     var a int
// // // // // //     fmt.Scan(&a)
// // // // // //     if a <= 59 {
// // // // // //     fmt.Println("Легкий вес")
// // // // // //     } else {
// // // // // //         if a < 64 {
// // // // // //             fmt.Println("Первый полусредний вес")
// // // // // //         } else {
// // // // // //             if a < 69 {
// // // // // //                 fmt.Println("Полусредний вес")}
// // // // // //         }
// // // // // //     }
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a, b, c int
// // // // // // 	fmt.Scanln(&a)
// // // // // // 	fmt.Scanln(&b)
// // // // // // 	fmt.Scanln(&c)
// // // // // // 	if a == b && b == c && a == c {
// // // // // // 		if a != b && b != c && a != c {
// // // // // // 			fmt.Print(0)
// // // // // // 		}
// // // // // // 		fmt.Print(3)

// // // // // // 	} else {
// // // // // // 		fmt.Print(2)
// // // // // // 	}

// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"math"
// // // // // // )

// // // // // // func SqRoot(a float64, b float64, c float64) int {
// // // // // // 	// var a, b, c float64
// // // // // // 	// fmt.Scanln(&a, &b, &c)
// // // // // // 	D := math.Pow(b, 2) - 4.0*a*c

// // // // // // 	RootD := math.Sqrt(D)
// // // // // // 	x1 := int((-b - RootD) / 2 * a)
// // // // // // 	x2 := int((-b + RootD) / 2 * a)

// // // // // // 	return x1
// // // // // // 	return x2

// // // // // // }

// // // // // // func main() {
// // // // // // 	var a, b, c float64
// // // // // // 	fmt.Scanln(&a, &b, &c)
// // // // // // 	SqRoot(a, b, c)

// // // // // // 	if a < 0 {
// // // // // // 		fmt.Println("Некорректный ввод")
// // // // // // 	} else {
// // // // // // 		if x1 == x2 {
// // // // // // 			fmt.Println(x1)
// // // // // // 		}
// // // // // // 		if x1 != x1 && x2 != 2 {
// // // // // // 			fmt.Println(0, 0)

// // // // // // 		} else {
// // // // // // 			fmt.Println(x1, x2)
// // // // // // 		}
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"errors"
// // // // // // 	"fmt"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var a, b int
// // // // // // 	var act string
// // // // // // 	fmt.Scan(&a)
// // // // // // 	fmt.Scan(&b)
// // // // // // 	fmt.Scan(&act)

// // // // // // 	if act == "-" || act == "+" || act == "/" || act == "*" {
// // // // // // 		if act == "+" {
// // // // // // 			res := a + b
// // // // // // 			fmt.Println(res)
// // // // // // 		}

// // // // // // 		if act == "-" {
// // // // // // 			res := a - b
// // // // // // 			fmt.Println(res)
// // // // // // 		}

// // // // // // 		if act == "*" {
// // // // // // 			res := a * b
// // // // // // 			fmt.Println(res)
// // // // // // 		}

// // // // // // 		if act == "/" {
// // // // // // 			if b == 0 {
// // // // // // 				var divError = errors.New("На ноль делить нельзя!")
// // // // // //                                 fmt.Println(divError)
// // // // // // 			} else {
// // // // // // 				res := float64(a) / float64(b)
// // // // // // 				fmt.Println(res)
// // // // // // 			}

// // // // // // 		}

// // // // // // 	} else {
// // // // // // 		fmt.Println("Неверная операция")
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"math"
// // // // // // )

// // // // // // func SqRoot(a float64, b float64, c float64) {
// // // // // // 	// var a, b, c float64
// // // // // // 	// fmt.Scanln(&a, &b, &c)
// // // // // // 	D := math.Pow(b, 2) - 4.0*a*c

// // // // // // 	RootD := math.Sqrt(D)
// // // // // // 	var x1 float64 = (-b - RootD) / 2 * a
// // // // // // 	var x2 float64 = (-b + RootD) / 2 * a

// // // // // // 	if a < 0 {
// // // // // // 		fmt.Println("Некорректный ввод")
// // // // // // 	} else {
// // // // // // 		if x1 == x2 {
// // // // // // 			fmt.Println(x1)
// // // // // // 		}
// // // // // // 		if x1 != x1 && x2 != 2 {
// // // // // // 			fmt.Println(0, 0)

// // // // // // 		} else {
// // // // // // 			fmt.Println(x1, x2)
// // // // // // 		}
// // // // // // 	}

// // // // // // }

// // // // // // func main() {
// // // // // // 	var a, b, c float64
// // // // // // 	fmt.Scanln(&a, &b, &c)
// // // // // // 	SqRoot(a, b, c)
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"math"
// // // // // // 	// "os"
// // // // // // )

// // // // // // func findDiscriminant(a, b, c float64) float64 {
// // // // // // 	discriminant := math.Pow(b, 2) - 4.0*a*c

// // // // // // 	// возвращаем значения вычисления из findDiscriminant
// // // // // // 	return discriminant
// // // // // // }

// // // // // // func SqRoot() {
// // // // // // 	var a, b, c float64
// // // // // // 	fmt.Scanln(&a)
// // // // // // 	fmt.Scanln(&b)
// // // // // // 	fmt.Scanln(&c)
// // // // // // 	D := findDiscriminant(a, b, c)
// // // // // // 	RootD := float64(math.Sqrt(D))
// // // // // // 	var x1 = float64((-b - RootD) / 2 * a)
// // // // // // 	var x2 = float64((-b + RootD) / 2 * a)
// // // // // // 	if a >= 0 {
// // // // // // 		if x1 == x2 {
// // // // // // 			x1 = float64(x1)
// // // // // // 			fmt.Println(x1)
// // // // // // 		}
// // // // // // 		if x1 == 0 && x2 == 0 {
// // // // // // 			x1 = float64(x1)
// // // // // // 			x2 = float64(x2)
// // // // // // 			fmt.Println(x1)
// // // // // // 			fmt.Println(x2)
// // // // // // 		}
// // // // // // 	} else {
// // // // // // 		fmt.Println("Некорректный ввод")
// // // // // // 	}
// // // // // // }

// // // // // // func main() {
// // // // // // 	SqRoot()
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"math"
// // // // // // )

// // // // // // func main() {
// // // // // // 	var a, b, c float64
// // // // // // 	fmt.Scan(&a)
// // // // // // 	fmt.Scan(&b)
// // // // // // 	fmt.Scan(&c)
// // // // // // 	D := math.Pow(b, 2) - 4 * a * c
// // // // // // 	RootD := math.Sqrt(D)
// // // // // // 	if D == 0 {
// // // // // // 		x := (-b / (2 * a))
// // // // // // 		fmt.Println(x)
// // // // // // 	}

// // // // // // 	if D < 0 {
// // // // // // 		fmt.Println()
// // // // // // 	}

// // // // // // 	if D > 0 {
// // // // // // 		x1 := ((-b - RootD) / (2 * a))
// // // // // // 		x2 := ((-b + RootD) / (2 * a))
// // // // // // 		fmt.Println(x1)
// // // // // // 		fmt.Println(x2)
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Println(&a)
// // // // // // 	m := map[int]int{
// // // // // // 		1:  31,
// // // // // // 		2:  29,
// // // // // // 		3:  31,
// // // // // // 		4:  30,
// // // // // // 		5:  31,
// // // // // // 		6:  30,
// // // // // // 		7:  31,
// // // // // // 		8:  31,
// // // // // // 		9:  30,
// // // // // // 		10: 31,
// // // // // // 		11: 30,
// // // // // // 		12: 31,
// // // // // // 	}
// // // // // // 	fmt.Println(a)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	switch a {
// // // // // // 	case 1:
// // // // // // 		fmt.Println(31)
// // // // // // 	case 2:
// // // // // // 		fmt.Println(29)
// // // // // // 	case 3:
// // // // // // 		fmt.Println(31)
// // // // // // 	case 4:
// // // // // // 		fmt.Println(30)
// // // // // // 	case 5:
// // // // // // 		fmt.Println(31)
// // // // // // 	case 6:
// // // // // // 		fmt.Println(30)
// // // // // // 	case 7:
// // // // // // 		fmt.Println(31)
// // // // // // 	case 8:
// // // // // // 		fmt.Println(31)
// // // // // // 	case 9:
// // // // // // 		fmt.Println(30)
// // // // // // 	case 10:
// // // // // // 		fmt.Println(31)
// // // // // // 	case 11:
// // // // // // 		fmt.Println(30)
// // // // // // 	case 12:
// // // // // // 		fmt.Println(31)
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	switch a {
// // // // // // 	case 12, 1, 2:
// // // // // // 		fmt.Println("Зима")
// // // // // // 	case 3, 4, 5:
// // // // // // 		fmt.Println("Весна")
// // // // // // 	case 6, 7, 8:
// // // // // // 		fmt.Println("Лето")
// // // // // // 	case 9, 10, 11:
// // // // // // 		fmt.Println("Осень")
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	b:=a%10
// // // // // // 	if b%2==0 {
// // // // // // 		fmt.Println("YES")
// // // // // // 	} else {
// // // // // // 		fmt.Println("NO")
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var n int
// // // // // // 	fmt.Scan(&n)
// // // // // // 	if n%2==0 {
// // // // // // 		fmt.Println(n+2)
// // // // // // 	} else {
// // // // // // 		fmt.Println(n+1)
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // //     "fmt"
// // // // // //     "math"
// // // // // // )

// // // // // //     func SqRoot() {
// // // // // // 	var a, b, c float64

// // // // // //     fmt.Scanln(&a, &b, &c)

// // // // // //     d := b * b - 4 * a * c

// // // // // //     if d == 0 {
// // // // // //         fmt.Println(-b / (2.0 * a))
// // // // // //     } else {
// // // // // //         if d > 0 {
// // // // // //             x1 := (-b + math.Sqrt(d)) / (2 * a)
// // // // // //             x2 := (-b - math.Sqrt(d)) / (2 * a)
// // // // // //             if x1 < x2 {
// // // // // //                 fmt.Println(float32(x1), float32(x2))
// // // // // //             } else {
// // // // // //                 fmt.Println(float32(x2), float32(x1))
// // // // // //             }
// // // // // //         }
// // // // // //     }
// // // // // // 	if d<0 {
// // // // // // 		fmt.Println(0, 0)
// // // // // // 	}
// // // // // // }

// // // // // // func main() {
// // // // // // 	SqRoot()
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // //     "fmt"
// // // // // //     "math"
// // // // // // )
// // // // // //     func SqRoot() {
// // // // // // 	var a, b, c float64

// // // // // //     fmt.Scan(&a, &b, &c)

// // // // // //     d := b * b - 4 * a * c

// // // // // //     if d == 0 {
// // // // // //         fmt.Println(-b / (2.0 * a))
// // // // // //     } else {
// // // // // //         if d > 0 {
// // // // // //             x1 := (-b + math.Sqrt(d)) / (2 * a)
// // // // // //             x2 := (-b - math.Sqrt(d)) / (2 * a)
// // // // // //             if x1 < x2 {
// // // // // //                 fmt.Println(float32(x1), float32(x2))
// // // // // //             } else {
// // // // // //                 fmt.Println(float32(x2), float32(x1))
// // // // // //             }
// // // // // //         }

// // // // // //     }
// // // // // // 	if d<0 {
// // // // // // 		fmt.Println(0, 0)
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	// "math"
// // // // // // )

// // // // // // func Add(a, b float64) float64 {
// // // // // // 	c := a + b
// // // // // // 	fmt.Println(a+b)
// // // // // // 	return c
// // // // // // }

// // // // // // func Multiply(a, b float64) float64 {
// // // // // // 	c := a * b
// // // // // // 	fmt.Println(a*b)
// // // // // // 	return c
// // // // // // }

// // // // // // func PrintNumbersAscending(n int) {
// // // // // // 	i := 1
// // // // // // 	for i <= n {
// // // // // // 		fmt.Println(i)
// // // // // // 		i++
// // // // // // 	}
// // // // // // }

// // // // // // func main() {
// // // // // // 	var a, b float64
// // // // // // 	var n int
// // // // // // 	fmt.Scanln(&a, &b, &n)
// // // // // // 	Add(a, b)
// // // // // // 	Multiply(a, b)
// // // // // // 	PrintNumbersAscending(n)
// // // // // // }

// // // // // // l := a / 100000000
// // // // // // k := a / 10000000
// // // // // // i := a / 1000000
// // // // // // h := a / 100000
// // // // // // g := a / 10000
// // // // // // f := a / 1000
// // // // // // b := a / 100
// // // // // // c := (a % 100) / 10
// // // // // // d := a % 10
// // // // // // e := b + c + d + f + g + h + i + k + l
// // // // // // fmt.Println(e)
// // // // // // break

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // )

// // // // // // func SumDigitsRecursive(n int) int {
// // // // // // 	var i int
// // // // // // 	for i = 1; a > 0; i++ {
// // // // // // 		b := a / 100
// // // // // // 		c := (a % 100) / 10
// // // // // // 		d := a % 10
// // // // // // 		var n int
// // // // // // 		for e > 0 && e < 10; n == 1; {
// // // // // // 			n := n * 10
// // // // // // 			e := (a / n) % 10
// // // // // // 		}
// // // // // // 		f := b + c + d + e
// // // // // // 		fmt.Println(f)
// // // // // // 	}
// // // // // // }

// // // // // // func main() {
// // // // // // 	var n int
// // // // // // 	fmt.Scan(&n)
// // // // // // 	SumDigitsRecursive(n)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // //     var n, k int
// // // // // //     fmt.Scan(&n, &k)

// // // // // //     factN := factorial(n)
// // // // // //     factK := factorial(k)
// // // // // //     nSubtractK := n - k
// // // // // //     factNSubtractK := factorial(nSubtractK) // можно factorial(n - k), и тогда строчка выше не нужна

// // // // // //     fmt.Println(factN / (factK * factNSubtractK))
// // // // // // }

// // // // // // func factorial(number int) int {
// // // // // //     fact := 1
// // // // // //     for i := 1; i <= number; i++ {
// // // // // //         fact = fact * i
// // // // // //     }
// // // // // //     return fact
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // )

// // // // // // func SumDigitsRecursive(n int) int {
// // // // // // 	var i, v int
// // // // // // 	for i = 1; n > 0; i++ {
// // // // // // 		b := n / 100
// // // // // // 		c := (n % 100) / 10
// // // // // // 		d := n % 10
// // // // // // 		for e > 0; e < 10; v == 1;  {
// // // // // // 			v := v * 10

// // // // // // 		}
// // // // // // 		f := b + c + d + e
// // // // // // 		fmt.Println(f)
// // // // // // 	}
// // // // // // 	return n
// // // // // // }

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	SumDigitsRecursive(a)
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // )

// // // // // // func SumDigitsRecursive(n int) int {
// // // // // // 	var i, v int
// // // // // // 	for i = 1; n > 0; i++ {
// // // // // // 		b := n / 100
// // // // // // 		c := (n % 100) / 10
// // // // // // 		d := n % 10
// // // // // // 		for v == 1  {
// // // // // // 			v := v * 10
// // // // // // 			e := (n/v)%10
// // // // // // 			f := b + c + d + e
// // // // // // 			fmt.Println(f)
// // // // // // 		}
// // // // // // 	}
// // // // // // 	return n
// // // // // // }

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	SumDigitsRecursive(a)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func SumDigitsRecursive(n int) int {
// // // // // // // Базовый случай: если n состоит из одной цифры, возвращаем ее значение
// // // // // // if n < 10 {
// // // // // // return n
// // // // // // }

// // // // // // // Рекурсивно вызываем функцию для оставшихся цифр числа
// // // // // // return n%10 + SumDigitsRecursive(n/10)
// // // // // // }

// // // // // // func main() {
// // // // // // // Пример использования функции
// // // // // // var number int
// // // // // // fmt.Scan(&number)
// // // // // // sum := SumDigitsRecursive(number)
// // // // // // fmt.Println(sum)
// // // // // // return

// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var a int
// // // // // // 	fmt.Scan(&a)
// // // // // // 	b := a / 100 // 1
// // // // // // 	c := (a % 100) / 10 // 2
// // // // // // 	d := a % 10 // 3
// // // // // // 	if b<c && b<d && c<d {
// // // // // // 		fmt.Println("YES")
// // // // // // 	} else {
// // // // // // 		fmt.Println("NO")
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var k, m int
// // // // // // 	fmt.Scan(&k)
// // // // // // 	fmt.Scan(&m)
// // // // // // 	n:=k/m
// // // // // // 	d:=k%m
// // // // // // 	if d!=0 {
// // // // // // 		n = n+1
// // // // // // 	}
// // // // // // 	fmt.Println(n)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func SumDigitsRecursive(n int) int {
// // // // // // // Базовый случай: если n состоит из одной цифры, возвращаем ее значение
// // // // // // if n < 10 {
// // // // // // return n
// // // // // // }

// // // // // // // Рекурсивно вызываем функцию для оставшихся цифр числа

// // // // // // sum:= n%10 + SumDigitsRecursive(n/10)
// // // // // // return sum
// // // // // // }

// // // // // // func main() {
// // // // // // // Пример использования функции
// // // // // // var number int
// // // // // // fmt.Scan(&number)
// // // // // // sum := SumDigitsRecursive(number)
// // // // // // fmt.Println(sum)
// // // // // // return

// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	var bin int
// // // // // // 	fmt.Scan(bin)

// // // // // // }

// // // // // // func IsPowerOfTwoRecursive(N int) {

// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func IsPowerOfTwoRecursive(N int) {
// // // // // //     if N == 1 {
// // // // // //         fmt.Println("YES")
// // // // // //     } else if N < 1 || N%2 != 0 {
// // // // // //         fmt.Println("NO")
// // // // // //     } else {
// // // // // //         IsPowerOfTwoRecursive(N / 2)
// // // // // //     }
// // // // // // }

// // // // // // func main() {
// // // // // // 	var N int
// // // // // // 	fmt.Scan(N)
// // // // // // 	IsPowerOfTwoRecursive(N)
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func IsPowerOfTwoRecursive(N int) {
// // // // // //     if N == 1 {
// // // // // //         fmt.Println("YES")
// // // // // //     } else if N < 1 || N%2 != 0 {
// // // // // //         fmt.Println("NO")
// // // // // //     } else {
// // // // // //         IsPowerOfTwoRecursive(N / 2)
// // // // // //     }
// // // // // // }

// // // // // // func main() {
// // // // // // 	var N int
// // // // // // 	fmt.Scan(N)
// // // // // //     IsPowerOfTwoRecursive(N) // YES
// // // // // //     // IsPowerOfTwoRecursive(18) // NO
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"math"
// // // // // // )

// // // // // // func SqRoot(a float64, b float64, c float64) float64 {
// // // // // // 	D := math.Pow(b, 2) - 4*a*c
// // // // // // 	RootD := math.Sqrt(D)
// // // // // // 	x1 := float64((-b - RootD) / 2 * a)
// // // // // // 	x2 := float64((-b + RootD) / 2 * a)
// // // // // // 	if D > 0 {
// // // // // // 		if x1 > x2 {
// // // // // // 			fmt.Println(x2, x1)
// // // // // // 		} else {
// // // // // // 			fmt.Println(x1, x2)
// // // // // // 		}
// // // // // // 	}

// // // // // // 	if D == 0 {
// // // // // // 		fmt.Println(x1)
// // // // // // 		return x1
// // // // // // 	}

// // // // // // 	if D < 0 {
// // // // // // 		fmt.Println(0, 0)
// // // // // // 	}

// // // // // // 	return x2

// // // // // // }

// // // // // // func main() {
// // // // // // 	var a, b, c float64
// // // // // // 	fmt.Scan(&a, &b, &c)
// // // // // // 	SqRoot(a, b, c)
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"math"
// // // // // // )

// // // // // // func SqRoots() {
// // // // // // 	var a, b, c float64
// // // // // //     fmt.Scanln(&a, &b, &c)
// // // // // //     D := math.Pow(b, 2) - 4*a*c
// // // // // // 	RootD := math.Sqrt(D)
// // // // // // 	x1 := float64((-b - RootD) / 2 * a)
// // // // // // 	x2 := float64((-b + RootD) / 2 * a)
// // // // // // 	if D > 0 {
// // // // // // 		if x1 > x2 {
// // // // // // 			fmt.Println(x2, x1)
// // // // // // 		} else {
// // // // // // 			fmt.Println(x1, x2)
// // // // // // 		}
// // // // // // 	}

// // // // // // 	if D == 0 {
// // // // // // 		fmt.Println(x1)
// // // // // // 	}

// // // // // // 	if D < 0 {
// // // // // // 		fmt.Println(0, 0)
// // // // // // 	}

// // // // // // }

// // // // // // func main() {
// // // // // // 	var a, b, c float64
// // // // // // 	fmt.Scan(&a, &b, &c)
// // // // // // 	SqRoots()
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // //     "fmt"
// // // // // //     "math"
// // // // // // )

// // // // // // func SqRoots() {
// // // // // //     var a, b, c float64
// // // // // //     fmt.Scanln(&a, &b, &c)

// // // // // //     d := b * b - 4 * a * c
// // // // // //     if d < 0.0 {
// // // // // //         fmt.Println("0 0")
// // // // // //     } else if sd := math.Sqrt(d); a > 0.0 {
// // // // // //         fmt.Printf("%g %g\n", (-b - sd) / (2.0 * a), (-b + sd) / (2.0 * a))
// // // // // //     } else {
// // // // // //         fmt.Printf("%g %g\n", (-b + sd) / (2.0 * a), (-b - sd) / (2.0 * a))
// // // // // //     }
// // // // // // }

// // // // // // func main() {
// // // // // // 	var a, b, c float64
// // // // // // 	fmt.Scan(&a, &b, &c)
// // // // // // 	SqRoots()
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // //  "fmt"
// // // // // //  "math"
// // // // // // )

// // // // // // func SqRoots() {
// // // // // //  var a, b, c float64

// // // // // //  _, err := fmt.Scanln(&a, &b, &c)
// // // // // //  if err != nil {
// // // // // //   return
// // // // // //  }

// // // // // //  discriminant := b*b - 4*a*c

// // // // // //  if discriminant > 0 {
// // // // // //   x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
// // // // // //   x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
// // // // // //   fmt.Println(math.Min(x1, x2), math.Max(x1, x2))
// // // // // //  } else if discriminant == 0 {
// // // // // //   x := -b / (2 * a)
// // // // // //   fmt.Printf("%.2f\n", x)
// // // // // //  } else {
// // // // // //   fmt.Println("0 0")
// // // // // //  }
// // // // // // }

// // // // // // func main() {
// // // // // //  SqRoots()
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"math"
// // // // // // )

// // // // // // func SqRoots() {
// // // // // // 	var a, b, c float64
// // // // // //     fmt.Scanln(&a, &b, &c)
// // // // // //     D := math.Pow(b, 2) - 4*a*c
// // // // // // 	RootD := math.Sqrt(D)
// // // // // // 	x1 := float64((-b - RootD) / 2 * a)
// // // // // // 	x2 := float64((-b + RootD) / 2 * a)
// // // // // // 	if D > 0 {
// // // // // // 		if x1 > x2 {
// // // // // // 			fmt.Println(x2, x1)
// // // // // // 		} else {
// // // // // // 			fmt.Println(x1, x2)
// // // // // // 		}
// // // // // // 	}

// // // // // // 	if D == 0 {
// // // // // // 		fmt.Println(x1)
// // // // // // 	}

// // // // // // 	if D < 0 {
// // // // // // 		fmt.Println(0, 0)
// // // // // // 	}

// // // // // // }

// // // // // // func main() {
// // // // // // 	var a, b, c float64
// // // // // // 	fmt.Scan(&a, &b, &c)
// // // // // // 	SqRoots()
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	// FiveSteps(steps)
// // // // // // }

// // // // // // func FiveSteps(array [5]int) [5]int {
// // // // // // 	var steps [5]int = [5]int{1, 2, 3, 4, 5}
// // // // // //     var newsteps [5]int = [5]int{}

// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func FiveSteps() {
// // // // // //     s := []int{1, 2, 3, 4, 5}

// // // // // //     for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// // // // // //         s[i], s[j] = s[j], s[i]
// // // // // //     }

// // // // // //     fmt.Println(s)
// // // // // // }

// // // // // // func main() {
// // // // // // 	FiveSteps()
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func FiveSteps() {
// // // // // //     steps := []int{1, 2, 3, 4, 5}

// // // // // //     fmt.Println(steps[4], steps[3], steps[2], steps[1], steps[0])
// // // // // // }

// // // // // // func main() {
// // // // // // 	FiveSteps()
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // )

// // // // // // func FiveSteps() {

// // // // // // }

// // // // // // func main() {
// // // // // // 	input := [5]int{1, 2, 3, 4, 5}
// // // // // // 	output := FiveSteps(input)
// // // // // // 	for i := 0; i < len(output); i++ {
// // // // // // 		fmt.Printf("%d ", output[i])
// // // // // // 	}
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func ReversSlice(slice []int) []int {
// // // // // //     reversed := make([]int, len(slice))
// // // // // //  	for i:=0; i<len(slice); i=i+1 {
// // // // // //  	    reversed[i] = slice[len(slice)-1-i]
// // // // // //     }
// // // // // // return slice
// // // // // // }

// // // // // // // func main() {
// // // // // // // 	ReversSlice(slice []int)
// // // // // // // }

// // // // // // func main() {
// // // // // // 	var slice [5]int = [5]int{1, 2, 3, 4, 5}
// // // // // // 	fmt.Println(slice)
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // )

// // // // // // func fiveSteps(arr [5]int) [5]int {
// // // // // // 	var reversedArr [5]int
// // // // // // 	for i := 0; i < 5; i++ {
// // // // // // 		reversedArr[i] = arr[4-i]
// // // // // // 	}
// // // // // // 	return reversedArr
// // // // // // }

// // // // // // func main() {
// // // // // // 	arr := [5]int{1, 2, 3, 4, 5}
// // // // // // 	reversedArr := fiveSteps(arr)
// // // // // // 	fmt.Println(reversedArr)
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // )

// // // // // // func ThirdElementInArray(array [6]int) int {

// // // // // //     z:=array[2]
// // // // // //     return z
// // // // // // }

// // // // // // func main() {
// // // // // // 	input := [6]int{1, 1, 1, 1, 1, 1}
// // // // // // 	output := ThirdElementInArray(input)
// // // // // // 	fmt.Print(output)
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"sort"
// // // // // // )

// // // // // // // import "math"

// // // // // // func FindMinMaxInArray(array [10]int) (int, int) {
// // // // // // 	array = []
// // // // // // 	sort.Ints(array)
// // // // // // 	min := array[0]
// // // // // // 	max := array[len(array)-1]
// // // // // // 	return max, min
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	// "fmt"
// // // // // // )

// // // // // // func findMaxMin(arr [10]int) (int, int) {
// // // // // // 	// Initialize the variables to hold the maximum and minimum values to draw comparisons.
// // // // // // 	max := arr[0]
// // // // // // 	min := arr[0]
// // // // // // 	// Iterate over the array
// // // // // // 	for i := 1; i < len(arr); i++ {
// // // // // // 		// if the current element is greater than the present maximum
// // // // // // 		if arr[i] > max {
// // // // // // 			max = arr[i]
// // // // // // 		}
// // // // // // 		// if the current element is smaller than the present minimum
// // // // // // 		if arr[i] < min {
// // // // // // 			min = arr[i]
// // // // // // 		}
// // // // // // 	}

// // // // // // 	return max, min
// // // // // // }

// // // // // // // func main() {
// // // // // // // 	// Example
// // // // // // // 	array := []int{9, 5, 7, 2, 1, 6, 8, 3, 4}
// // // // // // // 	maximum, minimum := findMaxMin(array)
// // // // // // // 	fmt.Println("Minimum:", minimum)
// // // // // // // 	fmt.Println("Maximum:", maximum)
// // // // // // // }

// // // // // // package main

// // // // // // func SumOfArray(array [6]int) int {
// // // // // // 	a := array[0] + array[1] + array[2] + array[3] + array[4] + array[5]
// // // // // // 	return a
// // // // // // }

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func PrettyArrayOutput(array [9]string) {
// // // // // // 	for i := 1; i < len(array); i++ {
// // // // // // 		for a := 0; a != 8; a++ {
// // // // // // 			fmt.Println(i, "я уже сделал:", array[0])
// // // // // // 		}
// // // // // // 	}
// // // // // // }

// // // // // // func main() {
// // // // // //     PrettyArrayOutput()
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // )

// // // // // // func prettyArrayOutput(arr [9]string) {
// // // // // // 	for i := 0; i < 7; i++ {
// // // // // // 		fmt.Printf("%d Я уже сделал: %s\n", i+1, arr[i])
// // // // // // 	}

// // // // // // 	fmt.Printf("8 Не успел сделать: %s\n", arr[7])
// // // // // // 	fmt.Printf("9 Не успел сделать: %s\n", arr[8])
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"sort"
// // // // // // )

// // // // // // func UnderLimit(nums []int, limit int, n int) ([]int, error) {
// // // // // // 	sort.Ints(nums)
// // // // // //     for limit > nums[n] {

// // // // // //         fmt.Println(nums[0])
// // // // // //     }
// // // // // //     return
// // // // // // }

// // // // // // func main() {

// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // //     "errors"
// // // // // //     //"fmt"
// // // // // // )

// // // // // // func UnderLimit(numsint, limit int, n int) ( []int, error) {
// // // // // //     // Проверка на отрицательное значение n
// // // // // //     if n < 0 {
// // // // // //         return nil, errors.New ("n cannot be negative")
// // // // // //     }
// // // // // //     // Инициализация результирующего слайса
// // // // // //     result := []int{}
// // // // // //     // Перебор элементов в слайсе
// // // // // //     for _, num := range nums {
// // // // // //         if num < limit {
// // // // // //             result = append(result, num)
// // // // // //             if len(result) == n {
// // // // // //                 break
// // // // // //             }
// // // // // //         }
// // // // // //     }
// // // // // //     return result, nil
// // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // //     "errors"
// // // // // //     // "fmt"
// // // // // // )

// // // // // // func UnderLimit(nums int, limit int, n int) ( []int, error) {
// // // // // //     // Проверка на отрицательное значение n
// // // // // //     if n < 0 {
// // // // // //         return nil, errors.New ("n cannot be negative")
// // // // // //     }
// // // // // //     // Инициализация результирующего слайса
// // // // // //     result := []int{}
// // // // // //     // Перебор элементов в слайсе
// // // // // //     for _, num := range nums {
// // // // // //         if num < limit {
// // // // // //             result = append(result, num)
// // // // // //             if len(result) == n {
// // // // // //                 break
// // // // // //             }
// // // // // //         }
// // // // // //     }
// // // // // //     return result, nil
// // // // // // }

// // // // // package main

// // // // // import (
// // // // // 	"errors"
// // // // // 	"fmt"
// // // // // )

// // // // // func UnderLimit(nums []int, limit int, n int) ([]int, error) {
// // // // // 	// Проверка на отрицательное значение n
// // // // // 	if n < 0 {
// // // // // 		return nil, errors.New ("n cannot be negative")
// // // // // 	}

// // // // // 	// Инициализация результирующего слайса
// // // // // 	result := []int{}

// // // // // 	// Перебор элементов в слайсе
// // // // // 	for _, num := range nums {
// // // // // 		if num < limit {
// // // // // 			result = append(result, num)
// // // // // 			if len(result) == n {
// // // // // 				break
// // // // // 			}
// // // // // 		}
// // // // // 	}

// // // // // 	return result, nil
// // // // // }

// // // // // func main() {
// // // // // 	nums := []int{5, 3, 10, 7, 2, 8}
// // // // // 	limit := 6
// // // // // 	n := 3

// // // // // 	result, err := UnderLimit(nums, limit, n)
// // // // // 	if err != nil {
// // // // // 		fmt.Println("Error:", err)
// // // // // 	} else {
// // // // // 		fmt.Println("Result:", result)
// // // // // 	}
// // // // //     fmt.Print(limit)
// // // // // }

// // // // package main

// // // // func Clean(nums []int, x int) ([]int) {
// // // //     // cleaned := make([]int, len(nums))
// // // //     for i:=0; i<len(nums); i++ {
// // // //         if nums[x] == x {
// // // //             nums[x] = 0
// // // //         }
// // // //     }
// // // //     return nums
// // // // }

// // // // // func main() {

// // // // // }

// package main

// func SliceCopy(nums []int) []int {
// 	new := nums
// 	a := len(new)
// 	b := len(nums)
// 	c := cap(new)
// 	d := cap(nums)
// if a==b {
//     c = d
// }
// 	return
// }

// // // func main() {

// // // }

// // package main

// // import "fmt"

// // func Clean(nums []int, x int) []int {
// // 	i := 0
// // 	for _, num := range nums {
// // 		if num != x {
// // 			nums[i] = num
// // 			i++
// // 		}
// // 	}
// // 	return nums[:i]
// // }

// // func main() {
// // 	nums := []int{1, 2, 3, 4, 2, 5, 2, 6}
// // 	x := 2
// // 	result := Clean(nums, x)
// // 	fmt.Println(result) // [1 3 4 5 6]
// // }

// package main

// import (
// 	"errors"
// 	//"fmt"
// )

// func UnderLimit(nums []int, limit int, n int) ([]int, error) {
// 	if n < 0 {
// 		return nil, errors.New ("n cannot be negative")
// 	}
//     result := []int{}

// 	for _, num := range nums {
// 		if num < limit {
// 			result = append(result, num)
// 			if len(result) == n {
// 				break
// 			}
// 		}
// 	}

// 	return result, nil
// }

// //func main() {
// //	nums := []int{5, 3, 10, 7, 2, 8}
// //	limit := 6
// ////	n := 3

// //	//result, err := UnderLimit(nums, limit, n)
// 	//if err != nil {
// //		fmt.Println("Error:", err)
// //	} else {
// 	//	fmt.Println("Result:", result)
// //	}
//   //  fmt.Print(limit)
// //}

// package main

// import (
// 	"errors"
// 	//"fmt"
// )

// func UnderLimit(nums []int, limit int, n int) ([]int, error) {
// 	// Проверка на отрицательное значение n
// 	if n < 0 {
// 		return nil, errors.New("n cannot be negative")
// 	}
// 	// Инициализация результирующего слайса
// 	result := []int{}
// 	// Перебор элементов в слайсе
// 	for _, num := range nums {
// 		if num < limit {
// 			result = append(result, num)
// 			if len(result) == n {
// 				break
// 			}
// 		}
// 	}
// 	// a := "n cannot be negative"

// 	return a, nil
// }

// package main

// import (
//     "errors"
//     //"fmt"
// )

// func UnderLimit(nums []int, limit int, n int) ( []int, error) {
//     // Проверка на отрицательное значение n
//     if n < 0 {
//         return nil, errors.New ("n cannot be negative")
//     }
//     // Инициализация результирующего слайса
//     result := []int{}
//     // Перебор элементов в слайсе
//     for _, num := range nums {
//         if num < limit {
//             result = append(result, num)
//             if len(result) == n {
//                 break
//             }
//         }
//     }
//     return result, nil
// }

// package main

// import (
//     "errors"
//     "fmt"
// )

// func UnderLimit(numsint, limit int, n int) ([]int, error) {
//     // Проверка на отрицательное значение n
//     if n < 0 {
//         return nil, errors.New ("n cannot be negative")
//     }
//     // Инициализация результирующего слайса
//     result := []int{}
//     // Перебор элементов в слайсе
//     for _, num := range nums {
//         if num < limit {
//             result = append(result, num)
//             if len(result) == n {
//                 break
//             }
//         }
//     }
//     return result, nil
// }

// func main() {
//     nums := []int{5, 3, 10, 7, 2, 8}
//     limit := 6
//     n := 3
//     result, err := UnderLimit(nums, limit, n)
//     if err != nil {
//         fmt.Println("Error:", err)
//     } else {
//         fmt.Println("Result:", result)
//     }
// }

// package main

// import (
// 	"errors"
// )

// func UnderLimit(nums []int, limit int, n int) ([]int, error) {
// 	if n <= 0 {
// 		return nil, errors.New("n должно быть больше нуля")
// 	}

// 	result := []int{}
// 	count := 0

// 	for _, num := range nums {
// 		if num < limit {
// 			result = append(result, num)
// 			count++
// 		}

// 		if count == n {
// 			break
// 		}
// 	}

// 	return result, nil
// }

// package main

// import "fmt"

// func SliceCopy(nums []int) []int {
// 	copiedSlice := make([]int, len(nums))
// 	copy(copiedSlice, nums)
// 	return copiedSlice
// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5}
// 	copied := SliceCopy(nums)

// 	fmt.Println("Original slice:", nums)
// 	fmt.Println("Copied slice:", copied)
// }

// package main

// import "fmt"

// func Join(nums1, nums2 []int) []int {
// 	joined := make([]int, 0, len(nums1)+len(nums2))

// 	joined = append(joined, nums1...)
// 	joined = append(joined, nums2...)

// 	return joined
// }

// func main() {
// 	nums1 := []int{1, 2, 3}
// 	nums2 := []int{4, 5, 6}

// 	result := Join(nums1, nums2)

// 	fmt.Println(result) // Вывод: [1 2 3 4 5 6]
// }

// package main

// import "fmt"

// func Mix(nums []int) []int {
// 	result := make([]int, len(nums))
// 	n := len(nums) / 2

// 	for i := 0; i < n; i++ {
// 		result[2*i] = nums[i]
// 		result[2*i+1] = nums[n+i]
// 	}

// 	return result
// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5, 6}
// 	result := Mix(nums)
// 	fmt.Println(result)
// } // Ожидаемый результат: [1 4 2 5 3 6]

// package main

// func StringLength(input string) int {
// 	a := len(input)
// 	if input=="" {
//         return -0
//     }
//     return a
// }

// func main() {
//     var input string
//     fmt.Scan(&input)
//     StringLength(input)
//     fmt.Println(a)
// }

// package main

// func ConcatenateStrings(str1, str2 string) string {
//     str := str1 + "" + str2
//     return str
// }

// package main

// import (
// 	"fmt"
// 	"unicode"
// )

// func isLatin(input string) bool {
// 	for _, r := range input {
// 		if r == _ {

// 		}
// 	}
// 	return
// }

// func main() {

// }

// func isLatin(input string) bool {
// 	for _, char := range input {
// 		if char < 'A' || (char > 'Z' && char < 'a') || char > 'z' {
// 			return false
// 		}
// 	}
// 	return true
// }

// package main

// import (
// 	"fmt"
// 	"unicode"
// )

// func isLatin(input string) bool {
// 	for _, char := range input {
// 		if !unicode.Is(unicode.Latin, char) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	fmt.Println(isLatin("Hello"))
// } // true fmt.Println(isLatin("Привет")) // false fmt.Println(isLatin("123")) // false }

// package main

// import (
//     "fmt"
// )

// func countOccurrences(nums []int) map[int]int {
//     occurrences := make(map[int]int)

//     for _, num := range nums {
//         occurrences[num]++
//     }

//     return occurrences
// }

// func main() {
//     numbers := []int{1, 2, 3, 2, 1, 3, 4, 5, 4, 1}
//     result := countOccurrences(numbers)

//     fmt.Println("Результат подсчёта повторений:")
//     for num, count := range result {
//         fmt.Printf("%d встречается %d раз(а)\n", num, count)
//     }
// }

// package main

// import (
//     "fmt"
// )

// func findStudentByID(id int, students map[int]string) (string, error) {
//     name, found := students[id]
//     if !found {
//         return "", fmt.Errorf("студент с ID %d не найден", id)
//     }

//     return name, nil
// }

// func main() {
//     students := map[int]string{
//         1: "Иванов",
//         2: "Петров",
//         3: "Сидоров",
//     }

//     id := 2
//     name, err := findStudentByID(id, students)
//     if err != nil {
//         fmt.Println(err)
//     } else {
//         fmt.Printf("Студент с ID %d: %s\n", id, name)
//     }
// }

// package main

// func FindMaxKey(m map[int]int) int {
//     var mapa := make(int, len(map))

//     return
// }

// func main() {

// }

// package main

// func FindMaxKey(m map[int]int) int {
// 	maxKey := 0
// 	for key := range m {
// 		if key > maxKey {
// 			maxKey = key
// 		}
//         if key < 0 {
//             maxKey = key
//         }
// 	}
// 	return maxKey
// }

// func main() {

// }

// package main

// func SumOfValuesInMap(m map[int]int) int {
// 	sum := 0
// 	for _, value := range m {
// 		sum += value
// 	}
// 	return sum
// }

// func main() {

// }

// package main

// import "fmt"

// func SwapKeysAndValues(m map[string]string) map[string]string {
// 	result := make(map[string]string)
// 	for key, value := range m {
// 		result[value] = key
// 	}
// 	return result
// }

// func main() {
// 	m := map[string]string{"Яндекс": "+74957397000", "Музей Яндекса": "+74991101133"}

// 	swapped := SwapKeysAndValues(m)
// 	fmt.Println(swapped)
// }

// package main

// func CountingSort(contacts []string) map[string]int {
// 	counts := make(map[string]int)

// 	for _, contact := range contacts {
// 		counts[contact]++
// 	}

// 	return counts
// }

// package main

// func DeleteLongKeys(m map[string]int) map[string]int {
// 	result := make(map[string]int)

// 	for key, value := range m {
// 		if len(key) >= 6 {
// 			result[key] = value
// 		}
// 	}

// 	return result
// }


// package main

// import (
//     "fmt"
//     "lesson/students"
// )

// func main() {
//     student1 := students.Student{}
//     fmt.Println(student1)
// }


// import "fmt"

// type Employee struct {
// 	name     (string)
// 	position (string)
// 	salary   (float64)
// 	bonus    (float64)
// }

// func (s Employee) CalculateTotalSalary() {
// 	fmt.Println("Employee:", s.name)
// 	fmt.Println("Position:", s.position)
// 	fmt.Print("Total Salary:")
// 	fmt.Printf("%.2f", s.bonus+s.salary)
// }

// func main() {
// 	man := Employee{name: "Anton", position: "product manager", salary: 100.0, bonus: 100.0}
// 	man.CalculateTotalSalary()
// }

// package main

// // import "fmt"

// // type Student struct {
// // 	name            (string)
// // 	solvedProblems  (int)
// // 	scoreForOneTask (float64)
// // 	passingScore    (float64)
// // }

// // func (s Student) IsExcellentStudent() bool {
// // 	if float64(s.solvedProblems)*s.passingScore < s.passingScore {
// // 		return true
// // 	} else {
// // 		return false
// // 	}
// // }

// // func main() {
// // 	student := Student{
// // 		name: "Arseniy",
// // 		solvedProblems: 5,
// // 		scoreForOneTask: 1.0,
// // 		passingScore: 100.0,
// // 	}
// // 	fmt.Print(student.IsExcellentStudent())
// // }

// package main

// type Student struct {
// 	name            string
// 	solvedProblems  int
// 	scoreForOneTask float64
// 	passingScore    float64
// }

// func (s Student) IsExcellentStudent() bool {
// 	TotalScore := s.solvedProblems * s.scoreForOneTask
// 	return TotalScore >= s.passingScore
// }


// package main

// import "time"

// type Task struct {
// 	summary     string
// 	description string
// 	deadline    time.Time
// 	priority    int
// }

// func (s Task) IsOverdue() bool {
// 	var a time.Time
// 	if  a < time.Now() {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func (c Task) IsTopPriority() bool {
// 	if c.priority == 4 || c.priority == 5 {
// 		return true
// 	} else {
// 		return false
// 	}
// } // true 4/5 false<4/5

// func main() {

// }


// // package main

// // // import (
// // // 	// "fmt"
// // // 	"time"
// // // )

// // // type Task struct {
// // // 	summary     string
// // // 	description string
// // // 	deadline    time.Time
// // // 	priority    int
// // // }

// // // func (t Task) IsOverdue() bool {
// // // 	return time.Now().After(t.deadline)
// // // }

// // // func (t Task) IsTopPriority() bool {
// // // 	return t.priority == 4 || t.priority == 5
// // // }

// // // type Note struct {
// // // name string

// // // tasks слайс Task

// // // notes слайс Note
// // // }

// // // package main

// // // import (
// // // 	"fmt"
// // // 	"time"
// // // )

// // // type Task struct {
// // // 	title     string
// // // 	priority  int
// // // 	deadline  time.Time
// // // 	completed bool
// // // }

// // // func (t Task) SetTitle(title string) { t.title = title }

// // // func (t Task) SetPriority(priority int) { t.priority = priority }

// // // func (t Task) SetDeadline(deadline time.Time) { t.deadline = deadline }

// // // func (t Task) SetCompleted(completed bool) { t.completed = completed }

// // // func (t Task) IsOverdue() bool { return !t.completed && time.Now().After(t.deadline) }

// // type Note struct {
// // 	title string
// // 	text  string
// // }

// // type ToDoList struct {
// // 	name  string
// // 	tasks []Task
// // 	notes []Note
// // }

// // func (tdl ToDoList) TasksCount() int {
// // 	return len(tdl.tasks)
// // }

// // func (tdl ToDoList) NotesCount() int {
// // 	return len(tdl.notes)
// // }

// // func (tdl ToDoList) CountTopPrioritiesTasks() int {
// // 	count := 0
// // 	for _, task := range tdl.tasks {
// // 		if task.priority > 5 {
// // 			count++
// // 		}
// // 	}
// // 	return count
// // }

// // func (tdl ToDoList) CountOverdueTasks() int {
// // 	count := 0
// // 	for _, task := range tdl.tasks {
// // 		if task.IsOverdue() {
// // 			count++
// // 		}
// // 	}
// // 	return count
// // }

// // // func main() {
// // // 	task1 := Task{title: "Task 1", priority: 3, deadline: time.Now().AddDate(0, 0, 1), completed: false} task2 := Task{title: "Task 2", priority: 7, deadline: time.Now().AddDate(0, 0, -1), completed: false}

// // // note1 := Note{title: "Note 1", text: "This is note 1"}
// // // note2 := Note{title: "Note 2", text: "This is note 2"}

// // // tdl := ToDoList{name: "My ToDo List", tasks: []Task{task1, task2}, notes: []Note{note1, note2}

// // // fmt.Println("Total tasks:", tdl.TasksCount())
// // // fmt.Println("Total notes:", tdl.NotesCount())
// // // fmt.Println("Top priority tasks:", tdl.CountTopPrioritiesTasks())
// // // fmt.Println("Overdue tasks:", tdl.CountOverdueTasks())
// // // }

// // // }



// package main

// import ( 
// 	"fmt" 
// "time" 
// )

// type Task struct { 
// 	title string 
// 	priority int 
// 	deadline time.Time 
// 	completed bool 
// }

// func (t *Task) SetCompleted() { 
// 	t.completed = true 
// }

// func (t *Task) IsOverdue() bool {
// 	 if !t.completed && time.Now().After(t.deadline) {
// 	 return true 
// 	} 
// return false 
// }

// type Note struct { 
// 	title string
// 	 text string 
// }

// type ToDoList struct { 
// 	name string 
// 	tasks []Task 
// 	notes []Note 
// }

// func (tdl *ToDoList) TasksCount() int { 
// 	return len(tdl.tasks) 
// }

// func (tdl *ToDoList) NotesCount() int { 
// 	return len(tdl.notes) 
// }

// func (tdl *ToDoList) CountTopPrioritiesTasks() int { 
// 	count := 0 for _, task := range tdl.tasks { 
// 	if task.priority > 5 { 
// 	count++ 
// 	} 
// } 
// return count
//  }

// func (tdl *ToDoList) CountOverdueTasks() int { 
// 	count := 0 
// 	for _, task := range tdl.tasks { 
// 	if task.IsOverdue() {
// 	 count++ 
// 	} 
// } 
// return count 
// }


// package main

// // import "math"

// type Shape interface {
// 	Area() float64
// }

// type Circle struct {
// 	radius float64
// }

// func (c Circle) Area() float64 {
// 	ac := 3.141592653589793 * c.radius * c.radius
// 	return ac
// }

// type Rectangle struct {
// 	width  float64
// 	height float64
// }

// func (r Rectangle) Area() float64 {
// 	rc := r.width * r.height
// 	return rc
// }


// package main

// // type Animal interface {
// // 	MakeSound() string
// // }

// // type Dog struct {
// // 	dog1 string
// // }

// // func (g Dog) MakeSound() string {
// // 	return g.dog1
// // }

// // type Cat struct {
// // 	cat1 string
// // }

// // func (c Cat) MakeSound() string {
// // 	return c.cat1
// // }

// // func main() {
// // 	cat1 := Cat{}
// // 	cat1.MakeSound()
// // 	dog1 := Dog{}
// // 	dog1.MakeSound()
// // }

// package main

// import "fmt"

// type Animal interface{ 
// 	MakeSound() 
// }

// type Dog struct{}

// func (d Dog) MakeSound() {
// 	 fmt.Println("Гав!") 
// 	}

// type Cat struct{}

// func (c Cat) MakeSound() { 
// 	fmt.Println("Мяу!") 
// }


//type LogLevel string

// const (
// 	Error LogLevel = "Error"
// 	Info  LogLevel = "Info"
// )

// type Logger interface {
// 	Log(message string)
// }

// type Log struct {
// 	Level LogLevel
// }

// func (l Log) Log(message string) {
// 	switch l.Level {
// 	case Error:
// 		fmt.Printf("ERROR: %s\n", message)
// 	case Info:
// 		fmt.Printf("INFO: %s\n", message)
// 	default:
// 		fmt.Printf("%s: %s\n", l.Level, message)
// 	}
// }


// package main

// import (
// 	"fmt"
// )

// type Money float64

// func (m Money) String() string {
// 	return fmt.Sprintf("$%.2f", m)
// }

// func main() {
// 	var salary Money = 2000.50

// 	fmt.Println(salary.String())
// }



// package main

// import "strconv"

// func ConcatStringsAndInt(str1, str2 string, num int) string {
// 	str12 := str1 + str2
// 	str3 := strconv.Itoa(num)

// 	str := str12 + str3
// 	return str

// }


// package main

// import "fmt"

// func DivideIntegers(a, b int) (float64, error) {
// 	c := a / b
// 	if b == 0 {
// 		fmt.Println(0)
// 		fmt.Println("division by zero is not allowed")
// 	} else {

// 		return nil
// 	}
// }

// func main() {

// }

// package main

// import (
// 	"errors"
// 	// "fmt"
// )

// func DivideIntegers(a, b int) (float64, error) {
// 	if b == 0 {
// 		return 0, errors.New("division by zero is not allowed")
// 	}

// 	return float64(a) / float64(b), nil
// }


// package main

// import (
// 	"fmt"
// 	"time"
// )

// func TimeAgo(pastTime time.Time) string {
// 	duration := time.Since(pastTime)

// 	if duration.Seconds() < 60 {
// 		return fmt.Sprintf("%.0f seconds ago", duration.Seconds())
// 	} else if duration.Minutes() < 60 {
// 		return fmt.Sprintf("%.0f minutes ago", duration.Minutes())
// 	} else if duration.Hours() < 24 {
// 		return fmt.Sprintf("%.0f hours ago", duration.Hours())
// 	} else if duration.Hours() < 24*7 {
// 		return fmt.Sprintf("%.0f days ago", duration.Hours()/24)
// 	} else if duration.Hours() < 24*30 {
// 		return fmt.Sprintf("%.0f weeks ago", duration.Hours()/(24*7))
// 	} else if duration.Hours() < 24*365 {
// 		return fmt.Sprintf("%.0f months ago", duration.Hours()/(24*30))
// 	} else {
// 		return fmt.Sprintf("%.0f years ago", duration.Hours()/(24*365))
// 	}
// }

// func main() {
// 	pastTime := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
// 	fmt.Println(TimeAgo(pastTime))
// }



// func DeleteVowels(s string) string {
// 	var result string
// 	for i := 0; i < len(s); i++ {
// 		switch unicode.ToLower(rune(s[i])) {
// 		case 'a':
// 			continue
// 		case 'e':
// 			continue
// 		case 'i':
// 			continue
// 		case 'o':
// 			continue
// 		case 'u':
// 			continue
// 		}
// 		result += string(s[i])
// 	}
// 	return result
// }

// func DeleteVowels(s string) string {
// 	var result string
// 	for i := 0; i < len(s); i++ {
// 		switch unicode.ToLower(rune(s[i])) {
// 		case 'a':
// 			continue
// 		case 'e':
// 			continue
// 		case 'i':
// 			continue
// 		case 'o':
// 			continue
// 		case 'u':
// 			continue
// 		}
// 		result += string(s[i])
// 	}
// 	return result
// }


// package main

// import (
// 	// "fmt"
// 	// "fmt"
// 	"testing"
// 	// "unicode"
// )

// func TestDeleteVowels(t *testing.T) { 
// 	tests := []struct { 
// 	input string
// 	 expected string }{ {"hello", "hll"}, {"world", "wrld"}, {"apple", "ppl"}, {"banana", "bnn"}, {"elephant", "lphnt"}, {"yellow", "yllw"}, {"but", "bt"}, {"nice", "nc"} }

// for _, tests := range tests {
// 	result := DeleteVowels(tests.input)
// 	if result != tests.expected {
// 		t.Errorf("For input %s, expected %s, but got %s", tests.input, tests.expected, result)
// 	} 
// }
// }



// этот код:
// package main

// import ( "testing" )

// func TestDeleteVowels(t *testing.T) { tests := []struct { 
// 	input string
// 	 expected string }{ {"hello", "hll"}, {"world", "wrld"}, {"apple", "ppl"}, {"banana", "bnn"}, {"elephant", "lphnt"}, {"yellow", "yllw"}, }

// for _, test := range tests {
// 	result := DeleteVowels(test.input)
// 	if result != test.expected {
// 		t.Errorf("For input %s, expected %s, but got %s", test.input, test.expected, result)
// 	}
// }
// }

// при вот такой функции DeleteVowels:

// package main

// import "unicode"

// func DeleteVowels(s string) string {
// 	var result string
// 	for i := 0; i < len(s); i++ {
// 		switch unicode.ToLower(rune(s[i])) {
// 		case 'a':
// 			continue
// 		case 'e':
// 			continue
// 		case 'i':
// 			continue
// 		case 'o':
// 			continue
// 		case 'u':
// 			continue
// 		}
// 		result += string(s[i])
// 	}
// 	return result
// } 

// выводит 
// PASS
// coverage: 80.0% of statements
// и сообщение
// Unknown error 
// как исправить ошибку, не изменяя функцию DeleteVowels