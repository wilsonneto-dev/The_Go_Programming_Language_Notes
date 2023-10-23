package main

import (
	"fmt"
	"strconv"
)

func main() {
	// just printing
	fmt.Println("Hellooow!")

	// parses
	var i int = 40
	fmt.Printf("%v, %T\n", i, i)

	var j float32 = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var k string = string(i) // it will convert to the ASCII equivalent
	fmt.Printf("%v, %T\n", k, k)

	var l string = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", l, l)

	// --------------- ways to declare variables
	var foo int = 30
	var foo2 int
	foo3 := 31
	fmt.Printf("%v  %v %v", foo, foo2, foo3)
}
