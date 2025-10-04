package main

import (
	"myServ/internal/entitie/human"
	"myServ/internal/entitie/pet"
	"myServ/internal/service"
)

func main() {
	citizen := human.NewHuman("citizen", "Yakutsk", "Baska", 22)
	villager := human.NewHuman("villager", "Maya", "Uyban", 42)
	cat := pet.NewPet("cat", "Barsik", citizen.GetName(), 5)
	dog := pet.NewPet("dog", "Bobbik", villager.GetName(), 3)

	service.Do(citizen)
	service.Do(villager)
	service.Do(cat)
	service.Do(dog)
}
