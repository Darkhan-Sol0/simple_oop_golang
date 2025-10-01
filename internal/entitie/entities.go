package entitie

type person struct {
	name string
	age  int
}

type human struct {
	*person
	city string
}

type pet struct {
	*person
}

type cat struct {
	*pet
	owner string
}

type dog struct {
	*pet
	owner string
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newPet(name string, age int) *pet {
	return &pet{
		person: newPerson(name, age),
	}
}

func newDog(name, owner string, age int) Pet {
	return &dog{
		pet:   newPet(name, age),
		owner: owner,
	}
}

func newCat(name, owner string, age int) Pet {
	return &cat{
		pet:   newPet(name, age),
		owner: owner,
	}
}
