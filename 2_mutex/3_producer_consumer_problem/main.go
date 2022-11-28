package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NUmberOfPizzas = 10

var pizzasMade, pizzasFaild, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making

	// run forver or untill we receive a quil notification
	// try to make pizzas
	for {
		// try to make a pizza
		// decision
	}
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// print out a message
	color.Cyan("The Pizzeria is open for business")
	color.Cyan("------------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	// run the producer in the background
	go pizzeria(pizzaJob)
	//create and run consumer
}
