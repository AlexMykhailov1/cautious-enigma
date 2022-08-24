package observer

import "fmt"

type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Notify()
}

type concreteObservable struct {
	observers map[string]Observer
}

func NewConcreteObservable() *concreteObservable {
	return &concreteObservable{observers: make(map[string]Observer)}
}

func (co *concreteObservable) Subscribe(o Observer) {
	id := o.GetID()
	co.observers[id] = o
	fmt.Println("subscribed observer with id:", id)
}

func (co *concreteObservable) Unsubscribe(o Observer) {
	id := o.GetID()
	delete(co.observers, o.GetID())
	fmt.Println("unsubscribed observer with id:", id)
}

func (co *concreteObservable) Notify() {
	for _, observer := range co.observers {
		observer.Update()
	}
	fmt.Println("notified all subscribers")
}
