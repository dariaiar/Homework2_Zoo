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
		Animal: fox,
		Class:  "Predator",
		Meal:   "Meat",
	}
	foxClass.printClass()

	wolfClass := Classification{
		Animal: wolf,
		Class:  "Predator",
		Meal:   "Meat",
	}
	wolfClass.printClass()

	duckClass := Classification{
		Animal: duck,
		Class:  "Herbivore",
		Meal:   "Plants",
	}
	duckClass.printClass()

	elephantClass := Classification{
		Animal: elephant,
		Class:  "Herbivore",
		Meal:   "Plants",
	}
	elephantClass.printClass()

	println("-----------------------------------------------------------------------------------------------")

	Mike := Catcher{Name: "Mike", Animal: fox}
	Mike.CatchAnimal()

	Olson := Catcher{Name: "Olson", Animal: wolf}
	Olson.CatchAnimal()

	Rob := Catcher{Name: "Rob", Animal: duck}
	Rob.CatchAnimal()

	Nick := Catcher{Name: "Nick", Animal: elephant}
	Nick.CatchAnimal()

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

type Catcher struct {
	Name string
	Animal
}

func (h Catcher) CatchAnimal() {
	fmt.Printf("Catcher %s caught %s (%s) at %s\n", h.Name, h.AnimalName, h.AnimalType, h.AnimalLocation)
}
