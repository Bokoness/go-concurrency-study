package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	// the standard output of os, the program will write its logs here
	// the purpose of this variable is to return the Stdout to its original form when down with testing
	stdOut := os.Stdout

	// Creating our own custom standard out
	r, w, _ := os.Pipe()
	// Setting the os.Stdout to our custom stdout
	os.Stdout = w

	// define a new WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("epsilon", &wg)

	wg.Wait()

	//Closing our custom Stdout writing
	_ = w.Close()

	//Reading our data from custom Stdout
	result, _ := io.ReadAll(r)
	output := string(result)

	//Returning the os.Stdout to original
	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Expected to find epsilon, but it is not there")
	}
}
