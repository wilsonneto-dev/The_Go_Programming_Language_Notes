package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ooook")

	var a bool = true

	var b int = 1
	var c int8 = 1
	var d int16 = 1
	var e int32 = 1
	var f int64 = 1

	var g uint = 1
	var h uint8 = 1
	var i uint16 = 1
	var j uint32 = 1
	var k uint64 = 1

	var l float32 = 1
	var m float64 = 1

	var n complex64 = 5 + 2i
	var o complex128 = 5 + 2i

	var p string = ""
	var ch rune = 'a'

	fmt.Printf("%v", a)
	fmt.Printf("%v %v %v %v %v %v %v %v %v %v %v %v %v %v %v", b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
	fmt.Printf("%v", ch)
}
