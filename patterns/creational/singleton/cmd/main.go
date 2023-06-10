package main

import (
	"github.com/AlexMykhailov1/cautious-enigma/patterns/creational/singleton"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			wg.Add(1)
			singleton.GetInstance()
			wg.Done()
		}(wg)
	}
	wg.Wait()
}
