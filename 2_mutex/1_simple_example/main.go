package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	//The 2 go routines in main are trying to access and modify msg
	//This can cause bugs
	//Here, by using the Mutex library, we are locking all other routines from modifying msg
	m.Lock()
	msg = s
	//After this routines is done with modifying msg, we unlock for other routines to modify as well
	//By doing so we avoid race condition bugs
	m.Unlock()
}

func main() {
	msg = "Hello, world!"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, universe!", &mutex)
	go updateMessage("Hello, cosmos!", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
