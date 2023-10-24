package main

import "fmt"

func main() {
	// if with initializer
	states := map[string]int{
		"SP": 1,
		"RJ": 2,
		"GO": 3,
		// ...
	}

	if v, ok := states["SP"]; ok {
		fmt.Printf("%v %v\n", ok, v)
	}

	// switch, same initializer way
	switch i := 1 + 1; i {
	case 1:
		fmt.Println(1)
		fallthrough // to go to the next one
	case 2:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}

	// type switch
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("st")
	}
}
