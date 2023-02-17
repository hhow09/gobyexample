package main

import (
	"fmt"
)

// Parent class
type Animal struct {
	name    string
	subject string
	food    string
}

// Child Class
type Cat struct {
	// inherited from animal
	Animal
	// member specific for cat
	color string
}

// Child Class
type Fish struct {
	// inherited from animal
	Animal
	// member specific for Fish
	length int
}

// method specific for cat
func (c Cat) talk() {
	fmt.Printf("my name is %s. and my color is %s\n", c.name, c.color)
}

// method for all Animal
func (a Animal) eat() {
	fmt.Println(a.name + " only eat " + a.food + ", it belongs to" + a.subject)
}

// in practice, we usually use interface for abstraction
type CanEatInterface interface {
	eat()
}

type Alien struct{}

// Alien also implement CanEatInterface
func (a Alien) eat() {
	fmt.Println("nom nom")
}

// should accept AnimalInterface
func dinnerTime(a CanEatInterface) {
	fmt.Print("dinnerTime: ")
	a.eat()
}

func main() {
	a := Animal{name: "mammal", subject: "carnivorous", food: "maet"}
	// animal.eat()
	a.eat()

	// declare Cat
	cat := Cat{Animal: Animal{name: "meow", subject: "cat", food: "fish"}, color: "black"}
	cat.eat()
	cat.talk()

	// declare Fish
	fish := Fish{Animal: Animal{name: "blublu", subject: "deep ocean", food: "seeweed"}, length: 6}

	dinnerTime(a)
	dinnerTime(cat)
	dinnerTime(fish)

	al := Alien{}
	dinnerTime(al)
}
