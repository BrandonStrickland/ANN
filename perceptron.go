package main

import (
	"fmt"
	//"math"
	"math/rand"
)

type perceptron struct {
	weights [] float64
}

func makePerceptron(number int) perceptron {
	weights := make([]float64, number)
	for i, _ := range weights {
		weights[i] = rand.Float64() * 10 - 5
	}
	return perceptron{weights}
}

func main() {
	ann := makePerceptron(5)
	fmt.Println(ann.weights)
}
