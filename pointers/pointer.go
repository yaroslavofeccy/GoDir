package main

import (
	"fmt"
)

func f_sum(a, b *int) {
	*a = *a + *b
	*b = *a
}

func main() {
	var a int = 10
	var b int = 15

	var a_ptr *int = &a

	fmt.Println(a_ptr)
	fmt.Println(*a_ptr)

	f_sum(&a, &b)

	fmt.Println(a)
	fmt.Println(b)

	// In Go char is rune pizde....
	var c rune = 'c'
	fmt.Println(c)
}
