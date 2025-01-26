package algorithms

import "fmt"

type FizzBuzz interface {
	Generate()
}

func NewModuloFizzBuzz() FizzBuzz {
	return &moduloFizzBuzz{}
}

func NewFizzBuzz() FizzBuzz {
	return &fizzBuzz{}
}

type moduloFizzBuzz struct{}

func (f *moduloFizzBuzz) Generate() {
	fizzCount, buzzCount := 0, 0
	for i := 1; i <= 100; i++ {
		fizzCount++
		buzzCount++
		if fizzCount == 3 && buzzCount == 5 {
			fmt.Println("FizzBuzz")
			fizzCount = 0
			buzzCount = 0
		} else if fizzCount == 3 {
			fmt.Println("Fizz")
			fizzCount = 0
		} else if buzzCount == 5 {
			fmt.Println("Buzz")
			buzzCount = 0
		} else {
			fmt.Println(i)
		}
	}
}

type fizzBuzz struct{}

func (f *fizzBuzz) Generate() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
