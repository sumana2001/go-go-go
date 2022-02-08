package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Goodmorning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("Tina")
	sayBye("John")
	cycleNames([]string{"tina", "mark", "john", "nia"}, sayGreeting)
	a1 := circleArea(10.5)
	a2 := circleArea(5)
	fmt.Printf("Area 1: %0.3f and Area 2: %0.3f", a1, a2)
}
