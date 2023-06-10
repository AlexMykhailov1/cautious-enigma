package main

import "github.com/AlexMykhailov1/cautious-enigma/patterns/behavioral/strategy"

func main() {
	cat := strategy.NewAnimal("cat",
		strategy.EatMeatBehavior{},
		strategy.WalkFastBehavior{},
		strategy.SleepLongBehavior{})
	cow := strategy.NewAnimal("cow",
		strategy.EatGrassBehavior{},
		strategy.WalkSlowBehavior{},
		strategy.SleepLittleBehavior{})

	cat.Describe()
	cow.Describe()
}
