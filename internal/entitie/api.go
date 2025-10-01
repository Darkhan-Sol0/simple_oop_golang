package entitie

func NewHuman(name, city string, age int) Human {
	return &human{
		person: newPerson(name, age),
		city:   city,
	}
}

func NewPet(petType, name, owner string, age int) Pet {
	switch petType {
	case "dog":
		return newDog(name, owner, age)
	case "cat":
		return newCat(name, owner, age)
	default:
		return nil
	}
}

func (h *human) From() string {
	return h.city
}

func (h *human) Speak() string {
	return "HUMAN!"
}

func (d *dog) GetOwner() string {
	return d.owner
}

func (d *dog) Speak() string {
	return "Bark!"
}

func (c *cat) GetOwner() string {
	return c.owner
}

func (c *cat) Speak() string {
	return "Mao!"
}

func (p *person) GetName() string {
	return p.name
}

func (p *person) GetAge() int {
	return p.age
}
