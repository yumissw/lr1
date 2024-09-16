package main

import (
	"fmt"
)

func main() {
	a := 15
	var b int = 4

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	r1 := a + b
	r2 := a - b
	r3 := a * b
	r4 := a / b

	fmt.Println("сложение = ", r1)
	fmt.Println("вычитание = ", r2)
	fmt.Println("умножение = ", r3)
	fmt.Println("деление = ", r4)
}
