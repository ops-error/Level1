package main

import "fmt"

type Human struct {
	name string
}

type Action struct {
	Human
	age int
}

func (h Human) AnsweringMachine() {
	fmt.Printf("Yo, yo, yo, %v! 148, 3-to-the-3-to-the-6-to-the-9. Representin’ the ABQ. What's up, biatch? Leave it at the tone!", h.name)
}

func (a Action) Greeting() {
	fmt.Printf("Hi! My name is %v, I'm %v years old\n", a.name, a.age)
}

func main() {

	userData := Action{
		Human: Human{"Anastasiia"},
		age:   23,
	}

	userData.Greeting()
	userData.AnsweringMachine()
}
