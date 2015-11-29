package main

import "fmt"

type Cat struct {
	name string
}

func (c Cat) Pet() {
	fmt.Println("hi, I'm cat")
}

func (c Cat) Name() string {
	return c.name
}

type Dog struct {
	name string
}

func (d Dog) Pet() {
	fmt.Println("hi, I'm dog")
}

func (d Dog) Name() string {
	return d.name
}

type Animal interface {
	Pet()
	Name() string
}

func Compliment(a Animal) {
	fmt.Println("Great Job", a.Name())
	a.Pet()
}

func main() {
	c := Cat{"Dora Chen"}
	Compliment(c)

	d := Dog{"Will Huang"}
	Compliment(d)
}
