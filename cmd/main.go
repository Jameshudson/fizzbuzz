package main

import (
	"github.com/jameshudson/fizzbuzz/internal/algorithms"
)

func main() {
	//fizzbuzz := interview.NewFizzBuzz()
	fizzbuzz := algorithms.NewModuloFizzBuzz()
	fizzbuzz.Generate()
}
