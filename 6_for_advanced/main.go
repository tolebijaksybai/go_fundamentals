package main

import "fmt"

func main() {
Label1:
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println("I:", i, "J:", j)
			if i >= 10 {
				continue Label1
			}
		}
	}

	for i := 0; i <= 20; i++ {
		if i >= 10 {
			break
		}
		fmt.Println(i)
	}

Label2:
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println("I:", i, "J:", j)
			if i >= 10 {
				break Label2
			}
		}
	}

}
