package main

import (
	"fmt"
)

type Vehicle struct {
	Brand string
	Model string
}

func (v Vehicle) Start() {
	fmt.Printf("The %s %s is starting...\n", v.Brand, v.Model)
}

type Car struct {
	Vehicle // composition
	Doors   int
}

func (c Car) Honk() {
	fmt.Printf("The %s %s is honking! Beep beep!\n", c.Brand, c.Model)
}

type Drivable interface {
	Drive()
}

func (c Car) Drive() {
	fmt.Printf("Driving the %s %s with %d doors...\n", c.Brand, c.Model, c.Doors)
}

type Bike struct {
	Vehicle
	HasCarrier bool
}

func (b Bike) Drive() {
	fmt.Printf("Riding the %s %s bike...\n", b.Brand, b.Model)
}

func main() {

	myCar := Car{
		Vehicle: Vehicle{Brand: "Toyota", Model: "Corolla"},
		Doors:   4,
	}

	myBike := Bike{
		Vehicle:    Vehicle{Brand: "Giant", Model: "Escape 3"},
		HasCarrier: true,
	}

	myCar.Start()
	myCar.Honk()
	myCar.Drive()

	myBike.Start()
	myBike.Drive()

	var vehicle Drivable

	vehicle = myCar
	vehicle.Drive()

	vehicle = myBike
	vehicle.Drive()
}
