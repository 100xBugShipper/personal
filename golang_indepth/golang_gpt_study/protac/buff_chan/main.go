package main

import "fmt"

type Car interface {
	Move()
}

type bmw struct {
	price int
}

func (b *bmw) Move() {
	fmt.Println("BMW moving")
}

type audi struct {
	price int
}

func (a *audi) Move() {
	fmt.Println("AUDI moving")
}

func moveTheCar(c Car) {
	c.Move()
}

func main() {
	bmwee := &audi{
		price:10,
	}

	moveTheCar(bmwee)
}
