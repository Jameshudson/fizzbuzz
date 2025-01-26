package algorithms

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
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			println("FizzBuzz")
		} else if i%3 == 0 {
			println("Fizz")
		} else if i%5 == 0 {
			println("Buzz")
		} else {
			println(i)
		}
	}
}

type fizzBuzz struct{}

func (f *fizzBuzz) Generate() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			println("FizzBuzz")
		} else if i%3 == 0 {
			println("Fizz")
		} else if i%5 == 0 {
			println("Buzz")
		} else {
			println(i)
		}
	}
}
