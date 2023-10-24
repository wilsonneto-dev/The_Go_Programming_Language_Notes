package main

import (
	"fmt"
	"reflect"
)

func main() {
	// annonymous struct
	aDeveloper := struct{ name string }{name: "Wilson"}
	fmt.Println(aDeveloper)

	// structs are passed by copies
	anotherDeveloper := aDeveloper
	anotherDeveloper.name = "Martin Fowler"
	fmt.Println(aDeveloper)       // Wilson
	fmt.Println(anotherDeveloper) // M<artin Fowler

	// if you want to pass by ref
	anotherDeveloper2 := &aDeveloper
	anotherDeveloper2.name = "Martin Fowler"
	fmt.Println(aDeveloper)        // Martin Fowler
	fmt.Println(anotherDeveloper2) // Martin Fowler

	// declaring the type
	type Developer struct {
		Name  string
		Age   int
		Stack []string
	}
	dev := Developer{Name: "Wilson", Age: 30, Stack: []string{"Go Lang", "C#", "Python", "JS/TS", "Azure", ".Net", "React.js", "Next.js"}}
	fmt.Println(dev)

	// Embeddings
	type Animal struct {
		Name   string
		Origin string
	}

	type Bird struct {
		Animal
		CanFly bool
	}

	// using the below sintax we need to explicity the embeding
	bird := Bird{Animal: Animal{Name: "Name", Origin: "BR"}, CanFly: true}
	fmt.Println(bird)

	// Example below we can sue as the other properties
	bird2 := Bird{}
	bird2.Name = "Bird 2"
	bird2.Origin = "US"
	bird2.CanFly = true
	fmt.Println(bird2)

	// tags
	type City struct {
		Name  string `required max:"100" json:"name,omitempty"`
		State string `json:"state,omitempty"`
	}

	r := reflect.TypeOf(City{})
	field, _ := r.FieldByName("Name")
	fmt.Println(field.Tag) //required max:"100" json:"name,omitempty"
}
