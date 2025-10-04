package base

type person struct {
	name string
	age  int
}

type Person interface {
	GetAge() int
	GetName() string
}

func NewPerson(name string, age int) Person {
	return &person{
		name: name,
		age:  age,
	}
}

func (p *person) GetName() string {
	return p.name
}

func (p *person) GetAge() int {
	return p.age
}
