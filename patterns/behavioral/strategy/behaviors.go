package strategy

import "fmt"

type EatMeatBehavior struct {
}

func (e EatMeatBehavior) Eat() {
	fmt.Println("eat meat")
}

type EatGrassBehavior struct {
}

func (e EatGrassBehavior) Eat() {
	fmt.Println("eat grass")
}

type WalkFastBehavior struct {
}

func (w WalkFastBehavior) Walk() {
	fmt.Println("walk fast")
}

type WalkSlowBehavior struct {
}

func (w WalkSlowBehavior) Walk() {
	fmt.Println("walk slow")
}

type SleepLongBehavior struct {
}

func (s SleepLongBehavior) Sleep() {
	fmt.Println("sleep long")
}

type SleepLittleBehavior struct {
}

func (s SleepLittleBehavior) Sleep() {
	fmt.Println("sleep little")
}
