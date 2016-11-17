package main

import (
	"fmt"
)

func main() {
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// use '_' to ignore some value
	var d, _, e float32 = 1.1, 2.2, 3.3
	fmt.Println(d, "\t", e)

	f := 123.4
	fmt.Println(f)

	var g int = int(f)
	fmt.Println(g)

	h := int(f)
	fmt.Println(h)
	fmt.Println(float32(h), ", ", float64(h))

	var (
		i = 100
		j = 1.23
		k = "golang"
		l = true
		m bool
		n float64
	)
	fmt.Println(i, ", ", j, ", ", k, ", ", l, ", ", m, ", ", n)

	// force convert number to string
	var o int = 65
	p := string(o)
	fmt.Println(p)
}
