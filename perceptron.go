package main

import (
	"fmt"
	"math"
	"math/rand"
)

/*
perceptron is a struct that models the neuron in a brain.
It has inputs, weights, and eventually a function. We will
explain how this all works together in an eval function.

Attributes:
weights               you can think of this as "how much
                      does that particular input matter"
inputs                the input represented as a float64
                      from the function we are making
                      predictions for.
*/
type perceptron struct {
	weights [] float64
	inputs [] float64
}

/*
makePerceptron is a factory function that can produce
different kinds of perceptrons. Later, there will be
a differing configuration for them and this will be 
more useful then.
*/
func makePerceptron(number int) perceptron {
	weights := make([]float64, number)
	for i, _ := range weights {
		weights[i] = rand.Float64() * 10 - 5
	}
	
	inputs := make([]float64,number)
	return perceptron{weights,inputs}
}

/*
dotProduct is just as it seems, it takes two vectors and multiplies
each pair of numbers and sums them all up and returns that sum.
*/
func (p perceptron) dotProduct() float64 {
	var dotProduct float64
	for i, _ := range p.weights {
		dotProduct += p.weights[i] * p.inputs[i]
	}
	return dotProduct
}

/* TODO: make this an attribute of the perceptron.
sigmoid is a function that governs the output of the perceptron.
We can think of this as the 
*/
func sigmoid(input float64) float64 {
	return 1 / (1 + math.Exp(-input))
}

/* TODO: reword this description!
This is the eval function. We multiply the inputs against
the weights (the dot product) and run THAT answer through
a function that governs the perceptron.
*/
func (p perceptron) eval() float64 {
	dp := p.dotProduct()
	return sigmoid(dp)
}

func main() {
	ann := makePerceptron(5)
	fmt.Println(ann.eval())
}
