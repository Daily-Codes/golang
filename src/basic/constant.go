package main

import (
	"fmt"
)

const a int = 100
const b = a
const (
	c  = 'A'
	c1 // eq c
	d  = b
	d1 // eq d
)

const e, f, g = 1, "HelloWorld.go", true

const h = len(f)

const (
	j1, j2 = true, 12.34
	// k1 eq j1, k2 eq j2
	k1, k2
	l1, l2
)

//
// go enum
//
const (
	p0 = iota
	p1 = 'A'
	p2 = iota
	p3 = '0'
	p4 = iota
)

func main() {
	fmt.Println(a, "\t", b)
	fmt.Println(c, "\t", d)
	fmt.Println(c1, "\t", d1)
	fmt.Println(e, "\t", f, "\t", g)
	fmt.Println("len(f) = ", h, "\tf = ", f)

	fmt.Println(j1, "\t", j2)
	fmt.Println(k1, "\t", k2)
	fmt.Println(l1, "\t", l2)

	fmt.Println(p0, p1, p2, p3, p4)
}
