package main

import (
	"fmt"
	"time"
)

type OurString string
type OutInt int

type Person struct {
	Name string
	Age  int
}

func main() {
	var customString OurString
	fmt.Printf("%T %#v \n", customString, customString)

	customString = "Hello world"
	fmt.Printf("%T %#v \n", customString, customString)

	customInt := OutInt(5)
	fmt.Printf("%T %#v \n", customInt, customInt)
	fmt.Println(int(customInt))

	var John Person
	fmt.Printf("%T %#v \n", John, John)

	John = Person{}
	fmt.Printf("%T %#v \n", John, John)

	John.Name = "John"
	John.Age = 15
	fmt.Println(John)

	Brad := Person{
		Name: "Brad",
		Age:  20,
	}
	fmt.Println(Brad)

	Tolebi := Person{"Tolebi", 20}
	fmt.Println(Tolebi)

	jTolebi := &John
	fmt.Println((*jTolebi).Age)
	fmt.Println(jTolebi.Age)

	jIvan := &Person{"Ivan", 20}
	fmt.Println(jIvan)

	unnamedStruct := struct {
		Name, LastName, Birthday string
	}{
		Name:     "NoName",
		LastName: "LastName",
		//Birthday: fmt.Sprintf("%s", time.Now()),
		Birthday: time.Now().String(),
	}

	fmt.Println(unnamedStruct)
}
