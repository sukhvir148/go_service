package main

import (
	"fmt"
	"sync"
)

type Account struct {
	// mx sync.Mutex
	money *int
}

var mx sync.Mutex

func (a *Account) credit() {
	mx.Lock()
	*a.money++
	fmt.Println("creditings")
	mx.Unlock()
}
