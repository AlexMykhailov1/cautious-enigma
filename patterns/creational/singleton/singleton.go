package singleton

import (
	"fmt"
	"sync"
)

type instance struct {
}

var singleInstance *instance

var lock = &sync.Mutex{}

func GetInstance() *instance {
	if singleInstance != nil {
		fmt.Println("instance already created")
	}
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleInstance = &instance{}
		fmt.Println("created instance")
	}
	return singleInstance
}
