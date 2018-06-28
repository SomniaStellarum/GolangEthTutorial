package main

import "fmt"

type Person struct {
	Name string
	age int
}

func NewPerson(_name string, _age int) (*Person) {
	return &Person{Name: _name, age: _age}
}

func (p *Person) SayName() {
	fmt.Print("My name is ", p.Name, "\n")
}

func (p *Person) SayAge() {
	fmt.Print("My age is ", p.age - 5, "\n")
}

func main() {
	person1 := NewPerson("Somnia", 34)
	person1.SayName()
	person1.SayAge()
}