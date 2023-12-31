package main

import "fmt"

func main() {

	// fixed length
	grades := [2]int{1, 2}
	fmt.Printf("Grades %v", grades)

	// also
	grades2 := [...]int{1, 2, 3}
	fmt.Printf("Grades %v", grades2)

	// also
	var students [3]string
	fmt.Printf("Students %v", students)

	// get the length
	fmt.Printf("Length %v", len(students))

	// 3d array
	var matrix [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("Matrix %v\n", matrix)

	// copies
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Printf("a %v b %v\n", a, b)
	// prints a [1 2 3] b [1 5 3]

	c := &a
	c[1] = 5
	fmt.Printf("a %v c %v\n", a, c)
	// prints a [1 5 3] c &[1 5 3]

	// print the pointer value
	fmt.Printf("*c: %T %v / c: %T %v\n", *c, *c, c, c)

	// Slices
	z := []int{1, 2, 3}
	fmt.Printf("slice 1: %v len: %v\n", z, len(z))

	z1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	z2 := z1[:]
	z3 := z1[1:4]
	z3[0] = 0
	fmt.Printf("z1 %v z2 %v z3 %v\n", z1, z2, z3)

	// make
	sl := make([]int, 3)
	fmt.Printf("%v - len %v\n", sl, len(sl))

	//capacity
	sl2 := make([]int, 3, 4)
	fmt.Printf("%v - len %v - cap %v\n", sl2, len(sl2), cap(sl2))
	sl2 = append(sl2, 5)
	sl2 = append(sl2, 6)
	fmt.Printf("%v - len %v - cap %v\n", sl2, len(sl2), cap(sl2))

	// testing append with common arrays
	sl10 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("bafore append in normal -> %v, cap %v, len %v\n", sl10, cap(sl10), len(sl10))
	sl10 = append(sl10, 10, 11, 12, 13)
	fmt.Printf("after append in normal -> %v, cap %v, len %v\n", sl10, cap(sl10), len(sl10))

}
