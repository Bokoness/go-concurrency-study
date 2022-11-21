package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup

	myMsg := "Hello Ness"
	wg.Add(1)
	go updateMessage(myMsg, &wg)
	wg.Wait()
	printMessage()

	//Closing our custom Stdout writing
	_ = w.Close()

	//Reading our data from custom Stdout
	result, _ := io.ReadAll(r)
	output := string(result)

	//Returning the os.Stdout to original
	os.Stdout = stdOut

	if !strings.Contains(output, myMsg) {
		t.Errorf("Expected to find %s, but found %s", myMsg, output)
	}
}
