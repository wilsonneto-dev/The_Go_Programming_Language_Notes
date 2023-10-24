package main

import "fmt"

func main() {
	// common way
	m1 := map[string]int{
		"Texas":      10,
		"California": 5,
	}
	m1["other"] = 8
	fmt.Println(m1, len(m1))

	// with make
	m2 := make(map[string]int, 10)
	m2["Test"] = 1
	m2["OtherTest"] = 2
	fmt.Println(m2, len(m2))

	// deleting an item:
	delete(m2, "Test")
	fmt.Printf("After deleting m2: %v \n", m2)

	// checking existence
	v1, ok1 := m2["OtherTest"]
	v2, ok2 := m2["OtherTestWrong"]

	fmt.Printf("Not found v1: %v ok1: %v\n", v1, ok1)
	fmt.Printf("Not found v2: %v ok2: %v\n", v2, ok2)

	_, exists := m2["OtherTestWrong"]
	if exists {
		fmt.Println("exists")
	} else {
		fmt.Println("It doesn't exists")
	}

	// when you pass you pass by reference
}
