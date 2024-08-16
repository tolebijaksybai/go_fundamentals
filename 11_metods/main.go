package main

import (
	"fmt"
	"time"
)

type OurType string

/*
func (t OurType) Hello() {
	fmt.Printf("Hello")
}
*/

type Square struct {
	Side int
}

func (s Square) Perimeter() {
	fmt.Printf("%T, %#v \n", s, s)
	fmt.Printf("Perimeter: %d\n", s.Side*4)
}

func (s *Square) Scale(multiplier int) {
	fmt.Printf("%T, %#v \n", s, s)
	s.Side *= multiplier
}

func (s Square) WrongScale(multiplier int) {
	fmt.Printf("%T, %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T, %#v \n", s, s)
}

func main() {
	definition()
	rules()
}

func definition() {
	square := Square{Side: 4}
	pSquare := &square

	square2 := Square{Side: 2}

	square.Perimeter()
	square2.Perimeter()

	pSquare.Scale(2)

	pSquare.Perimeter()
	pSquare.Scale(2)
	pSquare.Perimeter()

	square.WrongScale(2)
	square.Perimeter()
}

func rules() {
	now := time.Now()
	fmt.Printf("%T, %#v \n", now, now)

	var ourString OurType = "hello"
	fmt.Printf("%T, %#v \n", ourString, ourString)
}
