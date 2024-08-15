package main

import (
	"fmt"
	"unsafe"
)

var name = "Global variable"

var (
	config string
	isHead int
	gender bool = true
)

const age = 4

func main() {
	var num1 int64 = 15
	var num2 int = 13

	getType(num1 + int64(num2))

	fmt.Println(unsafe.Sizeof(uint8(num1)))

	surname := "Vasia"
	hello := fmt.Sprintf("hello world %s", surname)
	_ = hello // игнорирует

	fmt.Println(hello)

	fmt.Println(name)
	fmt.Println(age)

	name := "No global variable"
	fmt.Println(name)

	var nameString = "Tolebi"
	getType(nameString)

	nameString2 := "Tolebi"
	getType(nameString2)

	var isBool = true
	getType(isBool)

	isBool2 := true
	getType(isBool2)
}

func getType(param any) {
	fmt.Printf("Type: %T | Value: %v\n", param, param)
}
