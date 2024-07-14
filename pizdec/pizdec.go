package main

import "fmt"

type people interface {
	eat() bool
	shit() bool

	iq_counting() int
}
type params struct {
	p_type string
	p_age  int
}

func (p params) eat() bool {
	return true
}

func (p params) shit() bool {
	return false
}
func (p params) iq_counting() int {
	return 160
}

func check_life(peo people) {
	fmt.Println(peo)
	fmt.Println(peo.eat())
	fmt.Println(peo.shit())
	fmt.Println(peo.iq_counting())
}

func main() {
	p := params{"huesos", 21}

	fmt.Println(p.eat())

	check_life(p)
}
