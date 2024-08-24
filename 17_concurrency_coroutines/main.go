package main

import (
	"fmt"
)

func main() {
	//defer fmt.Println(1)
	//defer fmt.Println(2)

	//fmt.Println(sum(2, 3))

	//deferValues()

	makePanic()

	//runtime.GOMAXPROCS(1)
	//fmt.Printf("NumCPU: %d\n\n", runtime.NumCPU())

	//go showNumbers(100)

	//runtime.Gosched()
	//time.Sleep(time.Second)

	//fmt.Println("exit")
}

func makePanic() {
	defer func() {
		panicValue := recover()
		fmt.Println(panicValue)
	}()

	panic("some panic")
}

func deferValues() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("first", i)
	}

	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println("second", i)
		}()
	}

	for i := 0; i < 10; i++ {
		k := i
		defer func() {
			fmt.Println("third", k)
		}()
	}

	for i := 0; i < 10; i++ {
		defer func(k int) {
			fmt.Println("fourth", k)
		}(i)
	}
}

func showNumbers(number int) {
	for i := 0; i < number; i++ {
		fmt.Println(i)
	}
	fmt.Println()
}

func sum(x, y int) (sum int) {
	defer func() {
		sum *= 2
	}()

	sum = x + y
	return sum
}
