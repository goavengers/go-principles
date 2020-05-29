package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type AnimalBase struct {
	name string
}

type Lion struct{}

func (lion *Lion) MakeSound() string {
	return "roar"
}

type Squirrel struct{}

func (squirrel *Squirrel) MakeSound() string {
	return "squeak"
}

type Snake struct{}

func (snake *Snake) MakeSound() string {
	return "hiss"
}

func AnimalSoundsWrong() {
	animals := []AnimalBase{
		AnimalBase{name: "lion"},
		AnimalBase{name: "mouse"},
		AnimalBase{name: "snake"},
	}

	for _, animal := range animals {
		if animal.name == "lion" {
			fmt.Println("roar")
		} else if animal.name == "mouse" {
			fmt.Println("squeak")
		} else if animal.name == "snake" {
			fmt.Println("hiss")
		}
	}
}

func AnimalSoundsTrust() {
	animals := []Animal{
		&Lion{},
		&Squirrel{},
		&Snake{},
	}

	for _, animal := range animals {
		fmt.Println(animal.MakeSound())
	}
}

func main() {
	AnimalSoundsWrong()
	AnimalSoundsTrust()
}
