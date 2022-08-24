package observer

import "fmt"

type Observer interface {
	Update()
	GetID() string
}

type EmailObserver struct {
	Email string
}

func (eo EmailObserver) Update() {
	fmt.Println("sent email to", eo.Email)
}

func (eo EmailObserver) GetID() string {
	return eo.Email
}

type SMSObserver struct {
	Phone string
}

func (so SMSObserver) Update() {
	fmt.Println("sent sms to", so.Phone)
}

func (so SMSObserver) GetID() string {
	return so.Phone
}
