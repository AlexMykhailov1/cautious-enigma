package main

import "github.com/AlexMykhailov1/cautious-enigma/patterns/behavioral/observer"

func main() {
	observable := observer.NewConcreteObservable()
	emailObserver := observer.EmailObserver{Email: "test@gmail.com"}
	smsObserver := observer.SMSObserver{Phone: "123456789"}

	observable.Subscribe(emailObserver)
	observable.Subscribe(smsObserver)

	observable.Notify()

	observable.Unsubscribe(emailObserver)
	observable.Unsubscribe(smsObserver)
}
