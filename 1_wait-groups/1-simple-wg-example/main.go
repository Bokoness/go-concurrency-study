package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() // Using one WaitGroup item for each iteration

	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup //WaitGroup variable

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words)) // adding 9 items to wait group

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait() // wait for the loop routine to finish
}
