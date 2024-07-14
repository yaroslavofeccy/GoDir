package main

import (
	"fmt"
	"maps"
	"slices"
	"unicode/utf8"
)

const PI float32 = 3.14

func plusplus(a int, b int) int {
	return a + b
}

func plusng(a, b int) int {
	return a + b
}

func for_goods(a, b int) (int, int) {
	return b, a
}

// ...type give us a like C++ template<typename T>, cool thing
// template<typename T>
//
//	T sum(T a, T, b){
//		...
//	}
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Lambda func in Go is like a...
func mult() func(a, b int) int {
	return func(a, b int) int {
		return a * b
	}
}

type biba struct {
	boba         bool
	two_dolboebs bool
}

type rect struct {
	wight, height float64
}

func (r *rect) area() float64 {
	return r.wight * r.height
}

func (r rect) n_area() float64 {
	return r.wight * r.height
}

func (r *rect) make_things() {
	r.wight, r.height = 10.0, 15.0
}

func main() {
	fmt.Println("Hello, World!")

	var msg string
	msg = "Hello"

	var huinya = 3.15 + 5

	fmt.Println(msg)
	fmt.Println(huinya)
	fmt.Println(int64(huinya))

	i := 0
	for i <= 100 {
		fmt.Print(i)
		i++
	}
	for j := 0; j <= 100; j++ {
		fmt.Print(j)
	}

	for k := 0; k <= 100; k++ {
		if k%2 == 0 {
			fmt.Println(k)
		}
	}

	var num int = 10

	num = num + 5

	switch num {
	case 10:
		fmt.Println("Yap :)")
	case 15:
		fmt.Println("Not Yap :(")
	}
	var a [5]int
	a = [...]int{1, 2, 3, 4, 5}

	b := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(a, len(a))
	fmt.Println(b, len(b))

	for i := range a {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	// In golang Null, None and other Pizdez is -- nil !

	s := make([]string, 3)

	s[0] = "bebrick"
	s[2] = "Zhopick"

	s = append(s, "Pidorasik")
	s = append(s, "Dolboyebick")

	var c int = 0
	for i := range s {
		c++

		fmt.Println(s[i])
	}
	fmt.Println(c)

	h := make([]string, len(s))
	copy(h, s)

	fmt.Println("s ", s)
	fmt.Println("h ", h)

	var nn []int
	for l := range 1000 {
		nn = append(nn, l)
	}
	fmt.Println(nn[5:10])

	slices.Reverse(nn)
	fmt.Println(nn[5:10])

	if slices.Equal(h, s) {
		fmt.Println("Yap")
	}

	m := make(map[string]int)

	m["p1"] = 10
	m["p2"] = 3

	// the first one(tp) get the value of "p1":10 and the other one(pp) got true if its not nil or false if else
	tp, pp := m["p1"]

	fmt.Println(m)

	fmt.Println(pp)
	fmt.Println(tp)

	mmm := map[string]int{"pipi": 16, "popo": 18, "pypy": 21, "papapa": 30}
	delete(mmm, "papapa")
	fmt.Println(mmm)

	fmt.Println(m)

	if maps.Equal(m, mmm) {
		fmt.Println("Yap")
	} else {
		fmt.Println("Not YAP")
	}

	m2 := make(map[string]int, len(m))
	maps.Copy(m2, m)

	if maps.Equal(m, m2) {
		fmt.Println("Yap")
	} else {
		fmt.Println("Not YAP")
	}

	clear(m)

	for name, num := range m2 {
		fmt.Println(name, num)
	}

	fmt.Println(plusplus(1000, -7))
	fmt.Println(plusng(1000, -7))

	aa, bb := for_goods(10, 15)

	fmt.Println(aa, bb)

	sum(1, 2)
	sum(1, 2, 3, 4, 5, 6)

	nnn := []int{1, 2, 3, 4}

	// if wanna do sum of arr just add ... to it and all done!
	// Very cool lang
	sum(nnn...)

	// And declare lambda_func hmm...
	lambda_f := mult()
	fmt.Println(lambda_f(2, 2))

	// This is pizdec too...
	var bbb string = "Helli Biba"
	r, indx := utf8.DecodeRuneInString(bbb[0:])
	fmt.Printf("%#U bim bim at %d\n", r, indx)

	// For declare struct in Go plz use {}, not () or []
	bi := biba{true, true}
	fmt.Println(bi.boba, bi.two_dolboebs)

	// Chesno ne eby nahuiya tam v func nuzhen pointer ato pizdec kakoi to...
	rec := rect{wight: 2.0, height: 2.0}
	fmt.Println(rec.area())
	fmt.Println(rec.n_area())

	rec.make_things()
	fmt.Println(rec.area())
	fmt.Println(rec.n_area())
}
