package service

import (
	"fmt"
	"myServ/internal/entitie"
)

func Do(obj any) {
	switch v := obj.(type) {
	case entitie.Pet:
		fmt.Printf("Name: %s Owner: %s Age: %d Say: %s \n", v.GetName(), v.GetOwner(), v.GetAge(), v.Speak())
	case entitie.Human:
		fmt.Printf("Name: %s City: %s Age: %d Say: %s \n", v.GetName(), v.From(), v.GetAge(), v.Speak())
	default:
		fmt.Println("none date")
	}
}
