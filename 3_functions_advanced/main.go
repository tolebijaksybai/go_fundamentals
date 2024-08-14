package main

import "fmt"

func main() {
	first, second := 1, 2

	var multiplier func(x, y int) int

	multiplier = func(x, y int) int { return x + y }
	fmt.Printf("multiplier(): %d\n", multiplier(first, second))

	divider := func(x, y int) int { return x / y }
	fmt.Printf("divider(): %d\n", divider(second, first))

	sumFunc := func(x, y int) int {
		return x + y
	}

	subtractFunc := func(x, y int) int {
		return x - y
	}

	fmt.Printf("calculate(sumFunc): %d\n", calculate(second, first, sumFunc))
	fmt.Printf("calculate(subtractFunc): %d\n", calculate(second, first, subtractFunc))

	divideBy2 := createDivider(2)
	fmt.Printf("divideBy2(createDivider): %d\n", divideBy2(50))

	dollar := 30

	getDollarValue := func() int {
		return dollar
	}

	fmt.Printf("getDollarValue(): %d\n", getDollarValue())

	dollar = 40

	fmt.Printf("getDollarValue(): %d\n", getDollarValue())
}

func calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)
}

func createDivider(divider int) func(x int) int {
	dividerFunc := func(x int) int {
		return x / divider
	}
	return dividerFunc
}
