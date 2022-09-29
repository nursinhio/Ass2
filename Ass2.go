package main

import "fmt"

type Cloth struct {
	person IPerson
}

func (c *Cloth) ShowClothes() string {
	return c.person.ShowClothes() + " + " + "cloth"
}

type Jacket struct {
	person IPerson
}

func (c *Jacket) ShowClothes() string {
	return c.person.ShowClothes() + " + " + "T-Shirt"
}

type IPerson interface {
	ShowClothes() string
}

type Person struct {
	name  string
	cloth string
}

func (p *Person) ShowClothes() string {
	return p.name
}

type AddCloth func() string
func main() {

	person := &Person{
		name: "BOB",
	}

	personWithCloth := &Cloth{
		person: person,
	}
	personWithCloth.ShowClothes()
	personWithClothAndJacket := Jacket{
		person: personWithCloth,
	}
	fmt.Println(personWithClothAndJacket.ShowClothes())
}