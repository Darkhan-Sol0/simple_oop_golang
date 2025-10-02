package service

import (
	"fmt"
	"myServ/internal/entitie/human"
	"myServ/internal/entitie/pet"
)

func Do(obj any) {
	switch v := obj.(type) {
	case human.Human:
		fmt.Printf("%s, %s, %d \n", v.GetCity(), v.GetName(), v.GetAge())
	case pet.Pet:
		fmt.Printf("%s, %s, %d \n", v.Speak(), v.GetName(), v.GetAge())
	default:
		fmt.Printf("hui\n")
	}
}
