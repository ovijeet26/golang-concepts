package creational

import (
	"errors"
	"fmt"
)

/*
Factory Method is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.

It’s impossible to implement the classic Factory Method pattern in Go due to lack of OOP features such as classes and inheritance. However, we can still implement the basic version of the pattern, the Simple Factory.
*/

// Product Interface
type IGun interface {
	setName(string)
	setPower(int)
	getName() string
	getPower() int
}

// Concrete product
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(pow int) {
	g.power = pow
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

// Concrete implementation
type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun{
			name:  "AK47",
			power: 100,
		},
	}
}

// Concret Implementation
type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun{
			name:  "Musket",
			power: 10,
		},
	}
}

// FactoryMethod
func GetGunFactory(gunType string) (IGun, error) {

	switch gunType {
	case "AK47":
		obj := newAk47()
		return obj, nil
	case "Musket":
		obj := newMusket()
		return obj, nil
	default:
		return nil, errors.New("invalid gun type")
	}
}

// Client code
func FactoryClientCode() {

	ak47, _ := GetGunFactory("AK47")
	musket, _ := GetGunFactory("Musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
