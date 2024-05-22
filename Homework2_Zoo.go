package main

import "fmt"

type Meal string

func main() {
	fox := Animal{Type: "fox", Name: "Dixy", Age: 2, Colour: "red", Location: "Section1"}
	fox.print()
	wolf := Animal{Type: "wolf", Name: "Richard", Age: 5, Colour: "grey", Location: "Section2"}
	wolf.print()
	var duck = Animal{Type: "duck", Name: "Daisy", Age: 8, Colour: "white", Location: "Swimming pool"}
	duck.print()
	var elephant = Animal{Type: "elephant", Name: "Edward", Age: 20, Colour: "grey", Location: "Garden"}
	elephant.print()
	panda := Animal{Type: "panda", Name: "Saya", Age: 15, Colour: "black/white", Location: "Office"}
	panda.print()

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

	pandaClass := Classification{
		Animal: panda,
		Class:  "Herbivore",
		Meal:   "Plants",
	}
	pandaClass.printClass()

	println("-----------------------------------------------------------------------------------------------")

	mike := Catcher{Name: "Mike"}
	mike.CatchAnimal(fox)

	olson := Catcher{Name: "Olson"}
	olson.CatchAnimal(wolf)

	rob := Catcher{Name: "Rob"}
	rob.CatchAnimal(duck)

	nick := Catcher{Name: "Nick"}
	nick.CatchAnimal(elephant)

	george := Catcher{Name: "George"}
	george.CatchAnimal(panda)

}

type Animal struct {
	Type     string
	Name     string
	Age      int
	Colour   string
	Location string
}

func (a Animal) print() {
	fmt.Printf("Animal %s %s, by the name %s, %d years old is lost in area: %s \n", a.Colour, a.Type, a.Name, a.Age, a.Location)
}

type Classification struct {
	Animal
	Class string
	Meal  string
}

func (c Classification) printClass() {
	fmt.Printf("%s %s by the name %s is %s and eats %s\n", c.Colour, c.Type, c.Name, c.Class, c.Meal)
}

type Catcher struct {
	Name string
}

func (h Catcher) CatchAnimal(a Animal) {
	fmt.Printf("Catcher %s caught %s (%s) at %s\n", h.Name, a.Name, a.Type, a.Location)
}
