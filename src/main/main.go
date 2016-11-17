package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var a int
	var b int8
	var c int16
	var d int32
	var e int64

	var f bool
	var g float32
	var h float64

	var i string

	var j [5]int

	var k uint
	var l uint8
	var m uint16
	var n uint32
	var o uint64

	fmt.Println(a, "\t", b, "\t", c, "\t", d, "\t", e, "\t", f, "\t", g, "\t", h, "\t", i, "\t", j)

	a = math.MinInt8
	b = math.MinInt8
	c = math.MinInt16
	d = math.MinInt32
	e = math.MinInt64

	fmt.Println(a, "\t", b, "\t", c, "\t", d, "\t", e)

	a = math.MaxInt8
	b = math.MaxInt8
	c = math.MaxInt16
	d = math.MaxInt32
	e = math.MaxInt64

	k = math.MaxUint8
	l = math.MaxUint8
	m = math.MaxUint16
	n = math.MaxUint32
	o = math.MaxUint64

	fmt.Println(a, "\t", b, "\t", c, "\t", d, "\t", e)
	fmt.Println(k, "\t", l, "\t", m, "\t", n, "\t", o)

	fmt.Println("Hello World, ", time.Now())

}
