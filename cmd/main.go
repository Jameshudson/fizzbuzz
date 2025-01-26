package main

import "github.com/jameshudson/fizzbuzz/internal/algorithms"

func main() {
	fizzbuzz := algorithms.NewFizzBuzz()
	//fizzbuzz := algorithms.NewModuloFizzBuzz()
	fizzbuzz.Generate()
}
