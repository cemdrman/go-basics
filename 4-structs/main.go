package structs

import "fmt"

type Person struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	person := Person{
		Name:   "Cem",
		Age:    18,
		Active: true,
	}

	fmt.Println("Name", person.Name)
	fmt.Println("Age", person.Age)
	fmt.Println("Active", person.Active)

	person2 := Person{
		Name:   "Emir",
		Age:    18,
		Active: true,
	}

	var personList [2]Person

	personList[0] = person
	personList[1] = person2

	fmt.Print(personList)
}
