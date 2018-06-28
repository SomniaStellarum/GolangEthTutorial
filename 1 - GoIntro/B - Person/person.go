package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) SayName() {
	fmt.Print("My name is ", p.Name, "\n")
}

func main() {
	person1 := Person{Name: "Somnia"}
	person1.SayName()
}