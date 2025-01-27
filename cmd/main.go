package main

import "github.com/jameshudson/fizzbuzz/internal/algorithms"

func main() {
	//fizzbuzz := algorithms.NewModuloFizzBuzz()
	fizzbuzz := algorithms.NewNonModuloFizzBuzz()
	fizzbuzz.Generate()
}
