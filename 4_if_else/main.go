package main

import "fmt"

func main() {
	age := 15

	// basic if
	if age < 18 {
		fmt.Println("You are too young (full)")
	}

	// sort syntax
	if isChild := isChildren(age); isChild == true {
		fmt.Println("You are very young (short)")
		fmt.Println(isChild)
	}

	age = 17
	if age < 18 {
		fmt.Println("You are too young")
	} else {
		fmt.Println("You are an adult")
	}

	// Таблица истинности для: && 1 * 0 = 0
	if age >= 7 && age <= 18 {
		fmt.Println("You are in school")
	}

	// Таблица истинности для: || 1 + 0 + 1 = 1/0
	if age == 7 || age == 20 || age == 40 {
		fmt.Println("You have to get new passport")
	}

	// Отрицание: !
	// !true = false
	// !false = true

}

func isChildren(age int) bool {
	return age < 18
}
