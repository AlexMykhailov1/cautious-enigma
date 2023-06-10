package strategy

import "fmt"

type animal struct {
	name string
	eatBehavior
	walkBehavior
	sleepBehavior
}

type eatBehavior interface {
	Eat()
}

type walkBehavior interface {
	Walk()
}

type sleepBehavior interface {
	Sleep()
}

func NewAnimal(name string, eatBehavior eatBehavior, walkBehavior walkBehavior, sleepBehavior sleepBehavior) *animal {
	return &animal{
		name:          name,
		eatBehavior:   eatBehavior,
		walkBehavior:  walkBehavior,
		sleepBehavior: sleepBehavior,
	}
}

func (a *animal) Describe() {
	fmt.Printf("Description of %s:\n", a.name)
	a.Eat()
	a.Walk()
	a.Sleep()
	fmt.Println()
}
