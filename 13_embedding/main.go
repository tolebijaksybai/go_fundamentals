package main

import "fmt"

type Builder interface {
	Build()
}

type Person struct {
	Name string
	Age  int
}

type WorkExperience struct {
	Name string
	Age  int
}

type WoodBuilder struct {
	Person
	//Name string
	//WorkExperience
}

func (p Person) printName() {
	fmt.Println(p.Name)
}

//func (wb WoodBuilder) printName() {
//	fmt.Println(wb.Name)
//}

func (wb WoodBuilder) Build() {
	fmt.Println("Строю дом из дерева")
}

type BrickBuilder struct {
	Person
}

func (bb BrickBuilder) Build() {
	fmt.Println("Строю из кирпича")
}

type Building struct {
	Builder
	Name string
}

func main() {
	//explanation()
	usecase()
}

func usecase() {
	woodenBuilding := Building{
		Builder: WoodBuilder{Person{
			Name: "Vasia",
			Age:  49,
		}},
		Name: "Деревянная изба",
	}

	woodenBuilding.Build()

	brickBuilding := Building{
		Builder: BrickBuilder{Person{
			Name: "Vasia2",
			Age:  59,
		}},
		Name: "Кирпичный дом",
	}

	brickBuilding.Build()
}

func explanation() {
	builder := WoodBuilder{Person{Name: "Vasya", Age: 30}}
	//builder := WoodBuilder{Person{Name: "Vasya", Age: 30}, "Tasya"}
	//builder := WoodBuilder{
	//	Person{Name: "Vasya", Age: 30},
	//	"Tolya",
	//	WorkExperience{Name: "Tasya", Age: 14}}

	fmt.Printf("Type: %T Value: %#v \n", builder, builder)

	// shorthands
	fmt.Println(builder.Person.Age)
	//fmt.Println(builder.WorkExperience.Age)

	// shadowing
	fmt.Println(builder.Name)

	builder.printName()
	builder.Person.printName()
}
