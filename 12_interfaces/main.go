package main

import "fmt"

type Runner interface {
	Run() string
}

type Swimmer interface {
	Swim() string
}

type Flyer interface {
	Fly()
}

type Ducker interface {
	Runner
	Swimmer
	Flyer
}

type Human struct {
	Name string
}

func (h Human) Run() string {
	return fmt.Sprintf("Человек %s бегает", h.Name)
}

func (h Human) writeCode() string {
	return fmt.Sprintf("Человек %s пишет код", h.Name)
}

type Duck struct {
	Name, Surname string
}

func (d Duck) Run() string {
	return "Утка бегает"
}

func (d Duck) Fly() string {
	return "Утка летает"
}

func main() {
	var runner Runner
	fmt.Printf("Type: %T | Value: %#v \n", runner, runner)

	if runner == nil {
		fmt.Println("Runner is nil")
	}

	var unnamedRunner *Human
	fmt.Printf("Type: %T | Value: %#v \n", unnamedRunner, unnamedRunner)

	runner = unnamedRunner
	fmt.Printf("Type: %T | Value: %#v \n", runner, runner)
	if runner == nil {
		fmt.Println("Runner is nil")
	}

	namedRunner := &Human{Name: "Jack"}
	fmt.Printf("Type: %T | Value: %#v \n", namedRunner, namedRunner)

	runner = namedRunner
	fmt.Printf("Type: %T | Value: %#v \n", runner, runner)

	var emptyInterface interface{} = unnamedRunner
	fmt.Printf("Type: %T | Value: %#v \n", emptyInterface, emptyInterface)

	emptyInterface = runner
	fmt.Printf("Type: %T | Value: %#v \n", emptyInterface, emptyInterface)

	emptyInterface = int64(1)
	fmt.Printf("Type: %T | Value: %#v \n", emptyInterface, emptyInterface)

	emptyInterface = true
	fmt.Printf("Type: %T | Value: %#v \n", emptyInterface, emptyInterface)

	typeAssertionAndPolymorphism()
}

func typeAssertionAndPolymorphism() {
	var runner Runner
	fmt.Printf("Type: %T | Value: %#v \n", runner, runner)

	john := &Human{"John"}
	runner = john
	polymorphism(john)
	typeAssertion(john)

	donald := &Duck{Name: "Donald", Surname: "Duck"}
	runner = donald
	polymorphism(donald)
	typeAssertion(donald)
}

func polymorphism(runner Runner) {
	fmt.Println(runner.Run())
}

func typeAssertion(runner Runner) {
	fmt.Printf("Type: %T | Value: %#v \n", runner, runner)
	if human, ok := runner.(*Human); ok {
		fmt.Printf("Type: %T | Value: %#v \n", human, human)
		fmt.Println(human.writeCode())
	}

	if duck, ok := runner.(*Duck); ok {
		fmt.Printf("Type: %T | Value: %#v \n", duck, duck)
		fmt.Printf(duck.Fly())
	}

	switch v := runner.(type) {
	case *Human:
		fmt.Println(v.writeCode())
	case *Duck:
		fmt.Println(v.Fly())
	default:
		fmt.Printf("Type: %T | Value: %#v \n", v, v)
	}
}
