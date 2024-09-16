package main

import (
	"fmt"
)

func func1(a, b float64) (float64, float64) {
	r1 := a + b
	r2 := a - b

	return r1, r2
}

func main() {
	a := 15.5
	b := 6.3

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	r1, r2 := func1(a, b)

	fmt.Println("сложение = ", r1)
	fmt.Println("вычитание = ", r2)
}
