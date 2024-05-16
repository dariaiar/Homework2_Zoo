package main

import "fmt"

type Meal string

func main() {
	fox := Animal{AnimalType: "fox", AnimalName: "Dixy", AnimalAge: 2, AnimalColour: "red", AnimalLocation: "Section1"}
	fox.print()
	wolf := Animal{AnimalType: "wolf", AnimalName: "Richard", AnimalAge: 5, AnimalColour: "grey", AnimalLocation: "Section2"}
	wolf.print()
	var duck = Animal{AnimalType: "duck", AnimalName: "Daisy", AnimalAge: 8, AnimalColour: "white", AnimalLocation: "Swimming pool"}
	duck.print()
	var elephant = Animal{AnimalType: "elephant", AnimalName: "Edward", AnimalAge: 20, AnimalColour: "grey", AnimalLocation: "Garden"}
	elephant.print()

	println("-----------------------------------------------------------------------------------------------")

	foxClass := Classification{
		Animal: Animal{AnimalType: "fox", AnimalName: "Dixy", AnimalAge: 2, AnimalColour: "red", AnimalLocation: "Section1"},
		Class:  "Predator",
		Meal:   "Meat",
	}
	foxClass.printClass()

	wolfClass := Classification{
		Animal: Animal{AnimalType: "wolf", AnimalName: "Richard", AnimalAge: 5, AnimalColour: "grey", AnimalLocation: "Section2"},
		Class:  "Predator",
		Meal:   "Meat",
	}
	wolfClass.printClass()

	duckClass := Classification{
		Animal: Animal{AnimalType: "duck", AnimalName: "Daisy", AnimalAge: 8, AnimalColour: "white", AnimalLocation: "Swimming pool"},
		Class:  "Herbivore",
		Meal:   "Plants",
	}
	duckClass.printClass()

	elephantClass := Classification{
		Animal: Animal{AnimalType: "elephant", AnimalName: "Edward", AnimalAge: 20, AnimalColour: "grey", AnimalLocation: "Garden"},
		Class:  "Herbivore",
		Meal:   "Plants",
	}
	elephantClass.printClass()

	println("-----------------------------------------------------------------------------------------------")

	findDuck()
}

type Animal struct {
	AnimalType     string
	AnimalName     string
	AnimalAge      int
	AnimalColour   string
	AnimalLocation string
}

func (a Animal) print() {
	fmt.Printf("Animal %s %s, by the name %s, %d years old is lost in area: %s \n", a.AnimalColour, a.AnimalType, a.AnimalName, a.AnimalAge, a.AnimalLocation)
}

type Classification struct {
	Animal
	Class string
	Meal  string
}

func (c Classification) printClass() {
	fmt.Printf("%s %s by the name %s is %s and eats %s\n", c.AnimalColour, c.AnimalType, c.AnimalName, c.Class, c.Meal)
}

func findDuck() {
	var ducklocation string = "Garden"
	var newDuckLocation *string = &ducklocation
	println("Duck is here:", *newDuckLocation)
}
