package main

import "fmt"

/*
func testComment(params...) (return values)  {
	function body
}
*/

func main() {
	first, second := 1, 2

	Greet()

	PersonGreet("Tolebi")

	FioGreet("Tolebi", "Zhaksybay")

	sum := Sum(first, second)
	fmt.Printf("Sum is: %d\n", sum)

	sum, multiply := SumMultiply(first, second)
	fmt.Printf("Sum is: %d | Multiply is: %d\n", sum, multiply)

	_, multiply64 := namedSumAndMultiply(first, second)
	fmt.Println(multiply64)
}

func Greet() {
	fmt.Println("Greet() function")
}

func PersonGreet(name string) {
	fmt.Printf("PersonGreet() function, %s\n", name)
}

func FioGreet(name, surname string) {
	fmt.Printf("FioGreet() function, Name: %s | Surname: %s\n", name, surname)
}

func Sum(first, second int) int {
	return first + second
}

func SumMultiply(first, second int) (int, int) {
	return first + second, first * second
}

func namedSumAndMultiply(first, second int) (sum int64, multiply int64) {
	sum = int64(first + second)
	multiply = int64(first) * int64(second)
	return // sum, multiply
}
