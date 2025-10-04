package human

import (
	"myServ/internal/entitie/human/citizen"
	"myServ/internal/entitie/human/villager"
)

type Human interface {
	Describe() string
	GetName() string
}

func NewHuman(humanType, city, name string, age int) Human {
	switch humanType {
	case "citizen":
		return citizen.NewCitizen(name, city, age)
	case "villager":
		return villager.NewVillager(name, city, age)
	default:
		return nil
	}
}
