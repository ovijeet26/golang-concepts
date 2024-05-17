package creational

import "fmt"

/*
Builder is a creational design pattern, which allows constructing complex objects step by step.
Unlike other creational patterns, Builder doesn’t require products to have a common interface. That makes it possible to produce different products using the same construction process.
*/

// Base Builder intrface
type House struct {
	windowType string
	doorType   string
	floor      int
}

type IBuilder interface {
	seWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

// Concrete implementation of the buildeer interface -> Normal House
type NormalHouse struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalHouseBuilder() *NormalHouse {
	return new(NormalHouse)
}

func (b *NormalHouse) seWindowType() {

	b.windowType = "wooden window"
}

func (b *NormalHouse) setDoorType() {
	b.doorType = "wooden door"
}

func (b *NormalHouse) setNumFloor() {
	b.floor = 2
}

func (b *NormalHouse) getHouse() House {

	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}

// Concrete implementattion of th builder interface -> Igloo House
type IglooHouse struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooHousBuilder() *IglooHouse {
	return new(IglooHouse)
}

func (b *IglooHouse) seWindowType() {

	b.windowType = "ice window"
}

func (b *IglooHouse) setDoorType() {
	b.doorType = "ice door"
}

func (b *IglooHouse) setNumFloor() {
	b.floor = 1
}

func (b *IglooHouse) getHouse() House {

	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}

func getBuilder(builderType string) IBuilder {

	switch builderType {
	case "nomal":
		return newNormalHouseBuilder()
	case "igloo":
		return newIglooHousBuilder()
	default:
		return newNormalHouseBuilder()
	}
}

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {

	return &Director{
		builder: b,
	}
}

func (d *Director) setDirector(b IBuilder) {

	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.seWindowType()
	d.builder.setDoorType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func BuilderClientCaller() {

	var houseBuilder IBuilder

	houseBuilder = getBuilder("normal")

	director := newDirector(houseBuilder)

	builder := director.buildHouse()

	fmt.Printf("Door type -> %s \nWindow type -> %s \n Number of floors -> %d \n", builder.doorType, builder.windowType, builder.floor)

	iglooBuilder := getBuilder("igloo")
	director.setDirector(iglooBuilder)

	builder = director.buildHouse()
	fmt.Printf("Door type -> %s \nWindow type -> %s \n Number of floors -> %d \n", builder.doorType, builder.windowType, builder.floor)
}
