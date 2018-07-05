package main

import "fmt"

type NameAger interface {
	SayName()
	SayAge()
}

func WhoDis(_nameAger NameAger) {
	_nameAger.SayName()
	_nameAger.SayAge()
}

// Person implements NameAger interface
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

// Dog implements NameAger interface
type Dog struct {
}

func (d *Dog) SayName() {
	fmt.Print("Woof!\n")
}

func (d *Dog) SayAge() {
	fmt.Print("Woof! Woof!!\n")
}

func main() {
	person := NewPerson("Somnia", 34)
	dog := &Dog{}
	WhoDis(person)
	WhoDis(dog)
}