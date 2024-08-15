package main

import "fmt"

func main() {
	var intPointer *int

	fmt.Printf("%T %#v \n", intPointer, intPointer)

	// panic: runtime error:
	//fmt.Printf("%T %#v %#v \n", intPointer, intPointer, *intPointer)

	var a int64 = 7
	fmt.Printf("%T %#v \n", a, a)

	var pointerA *int64 = &a
	fmt.Printf("%T %#v %#v \n", pointerA, pointerA, *pointerA)

	var newPointer = new(float32)
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)

	*newPointer = 3
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)

	num := 3
	square(num)
	fmt.Printf("%T %#v %#v \n", num, num, num)

	squarePointer(&num)
	fmt.Printf("%T %#v %#v \n", num, num, num)

	var wallet1 *int
	fmt.Println(hasWallet(wallet1))

	wallet2 := 0
	fmt.Println(hasWallet(&wallet2))

	wallet3 := 100
	fmt.Println(hasWallet(&wallet3))
}

func square(num int) int {
	return num * num
}

func squarePointer(num *int) {
	*num = *num * *num
}

func hasWallet(money *int) bool {
	return money != nil
}
