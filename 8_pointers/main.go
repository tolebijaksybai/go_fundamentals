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
}
