package service

import (
	"fmt"
)

type Obj interface {
	Describe() string
}

func Do(obj Obj) {
	fmt.Printf("%s\n", obj.Describe())
}
