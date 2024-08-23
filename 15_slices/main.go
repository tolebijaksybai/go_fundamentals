package main

import "fmt"

func main() {
	//variadicFunctions()
	//convertToArrayPointer()
	//passToFunction()
	//sliceWithNew()
}

func sliceWithNew() {
	slicePointer := new([]int)

	fmt.Printf("Type: %T Value: %#v \n", slicePointer, *slicePointer)
	fmt.Printf("Length: %d Capacity: %d \n\n", len(*slicePointer), cap(*slicePointer))

	newSlice2 := append(*slicePointer, 1)
	fmt.Printf("Type: %T Value: %#v \n", newSlice2, newSlice2)
	fmt.Printf("Length: %d Capacity: %d \n\n", len(newSlice2), cap(newSlice2))
}

func passToFunction() {
	initialSlice := []int{1, 2}
	fmt.Printf("Type: %T Value: %#v \n", initialSlice, initialSlice)
	fmt.Printf("Length: %d Capacity: %d \n\n", len(initialSlice), cap(initialSlice))

	changeValue(initialSlice)

	fmt.Printf("Type: %T Value: %#v \n", initialSlice, initialSlice)
	fmt.Printf("Length: %d Capacity: %d \n\n", len(initialSlice), cap(initialSlice))

	newSlice := append(initialSlice, 3)
	fmt.Printf("Type: %T Value: %#v \n", newSlice, newSlice)
	fmt.Printf("Length: %d Capacity: %d \n\n", len(newSlice), cap(newSlice))

	newSlice2 := appendValue(newSlice)
	fmt.Printf("Type: %T Value: %#v \n", newSlice, newSlice)
	fmt.Printf("Length: %d Capacity: %d \n\n", len(newSlice), cap(newSlice))

	fmt.Printf("Type: %T Value: %#v \n", newSlice2, newSlice2)
	fmt.Printf("Length: %d Capacity: %d \n\n", len(newSlice2), cap(newSlice2))
}

func appendValue(slice []int) []int {
	slice = append(slice, 4, 5)
	fmt.Printf("Type: %T Value: %#v \n", slice, slice)
	fmt.Printf("Length: %d Capacity: %d \n\n", len(slice), cap(slice))

	return slice
}

func changeValue(slice []int) {
	slice[1] = 15
}

func convertToArrayPointer() {
	initialSlice := []int{1, 2, 3}
	fmt.Printf("Type: %T Value: %#v \n", initialSlice, initialSlice)
	fmt.Printf("Length: %d Capacity: %d \n\n", len(initialSlice), cap(initialSlice))

	intArray := (*[3]int)(initialSlice)
	fmt.Printf("Type: %T Value: %#v \n", intArray, intArray)
	fmt.Printf("Length: %d Capacity: %d \n", len(intArray), cap(intArray))
}

func variadicFunctions() {
	showAllElements(1, 2, 3)
	showAllElements(1, 2, 3, 4, 5, 6, 7)

	firstSlice := []int{5, 6, 7, 8}
	secondSlice := []int{9, 3, 2, 1}
	showAllElements(firstSlice...)

	newSlice := append(firstSlice, secondSlice...)
	fmt.Printf("Type: %T Value: %#v \n", newSlice, newSlice)
}

func showAllElements(values ...int) {
	for _, value := range values {
		fmt.Println("value:", value)
	}
	fmt.Println()
}
