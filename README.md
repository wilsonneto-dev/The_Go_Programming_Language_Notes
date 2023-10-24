Studying GO by the book and documentation

- [ ] 🟡 https://www.youtube.com/watch?v=YS4e4q9oBaU
- [ ] https://www.youtube.com/watch?v=jFfo23yIWac
- [ ] [https://www.amazon.com.br/Go-Programming-Language-Brian-Kernighan](https://www.amazon.com.br/Go-Programming-Language-Brian-Kernighan/dp/0134190440/ref=sr_1_1?crid=3O3XQFDCFYMID&keywords=the+go+programming+language&qid=1695784756&sprefix=The+GO+pro%2Caps%2C235)

![image](https://github.com/wilsonneto-dev/The_Go_Programming_Language_Notes/assets/20674439/83878282-193a-4f89-b15e-9ff2a82bbbf0)

## Notes

### To create a new go module

````shell
go mod init example/hello
````
So just add the main like below:
````go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
````

Let's run it:
````shell
go run .
````

### Declaring variables

````go
var foo int = 30
var foo2 int
foo3 := 31
````

### Visibility / Naming Conventions

- Lower Case first letter: Package Scope
- Upper Case first letter: Export from Package / Exposed

### Type Conversions

- destinationType(variable)
- use strconv package for strings

## Primitives

### Types

````go
var b bool = true

var b int = 1
var b int8 = 1
var b int16 = 1
var b int32 = 1
var b int64 = 1

var b uint = 1
var b uint8 = 1
var b uint16 = 1
var b uint32 = 1
var b uint64 = 1

var b float32 = 1
var b float64 = 1

var b complex64 = 5+2i
var b complex128 = 5+2i

var s string = ""
var ch rune = 'a'
````

### Constants

````go
// same way
const a int = 3
const a = 3
````

### IOTA

````go
const (
    a = iota  // a = 0
    b         // b = 1
    c         // c = 2
)

const (
    a = iota * 2  // a = 0 * 2 = 0
    b             // b = 1 * 2 = 2
    c             // c = 2 * 2 = 4
)

const (
    a = iota  // a = 0
    b         // b = 1
)

const (
    c = iota  // c = 0, not 2!
    d         // d = 1
)

````

### Arrays / Slices

In Go arrays are types, if you copy a array to another variable, it will be a new array, not a reference.

- `len()` returns the length of the array/slice
- `cap()` returns the capacity of the array/slice
- You can declare one array in three ways: `var a [3]int`, `a := [3]int{1, 2, 3}` or `a := [...]int{1, 2, 3}`
- You can declare a slice in two ways: `var a []int` or `a := []int{1, 2, 3}` (Literal Style)
- You can use make to create a slice: `make([]int, 3, 4)` where 3 is the length and 4 is the capacity (Via make)
- Or you can create a slice form an existing array: `a := [3]int{1, 2, 3}` and `b := a[:]` or `b := a[1:2]` (Via existing array)

append() is a function that can be used to append a slice to another slice, for example:
```go
a := []int{1, 2, 3}
b := []int{4, 5, 6}
a = append(a, b...)
```

Some examples
````go

Some examples:
````go
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
````

### Spread Opeartor:

You can use the spread operator to append a slice to another slice for example

```go
appednd(sl1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
```

## Maps and Strcuts

### Maps

#### Introduction
In Go, a map is a built-in data type that provides a way to store unordered key-value pairs. It offers efficient lookup, addition, and deletion operations. The keys within a map are unique and can be of any type that is comparable (e.g., int, float, string, etc.), while values can be of any type.

#### Declaration and Initialization
Using Composite Literal

This is a concise way to declare and initialize a map with a predefined set of key-value pairs:

```go
m1 := map[string]int{
	"Texas":      10,
	"California": 5,
}
m1["other"] = 8
fmt.Println(m1, len(m1))

```

Using make

The make function provides a way to create a map. Optionally, you can provide an initial space allocation hint:

```go
m2 := make(map[string]int, 10)
m2["Test"] = 1
m2["OtherTest"] = 2
fmt.Println(m2, len(m2))
```

The number 10 here is a hint to the compiler about the initial allocation size. It does not limit the number of items the map can store.

#### Map Operations

Add/Update

To add or update a key-value pair in the map:

```go
m1["NewKey"] = 12
```

Delete

```go
delete(m2, "Test")
fmt.Printf("After deleting m2: %v \n", m2)
```
Check Existence

When accessing a map, you can retrieve both the value and a boolean indicating whether the key was present:

```go
v1, ok1 := m2["OtherTest"]
v2, ok2 := m2["OtherTestWrong"]

fmt.Printf("Found v1: %v ok1: %v\n", v1, ok1)
fmt.Printf("Not found v2: %v ok2: %v\n", v2, ok2)
```

A common idiom to check if a key exists is:

```go
_, exists := m2["OtherTestWrong"]
if exists {
	fmt.Println("exists")
} else {
	fmt.Println("It doesn't exist")
}
```

#### Map References

```go
func changeMap(m map[string]int) {
    m["Changed"] = 100
}

func main() {
    m := map[string]int{"Original": 50}
    changeMap(m)
    fmt.Println(m) // Outputs: map[Changed:100 Original:50]
}
```

#### Iterating Over a Map

Iterating over key and value pairs

```go
m := map[string]int{
    "a": 1,
    "b": 2,
    "c": 3,
}

for key, value := range m {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
}
```

Iterating over  only keys:

```go
for key := range m {
    fmt.Println("Key:", key)
}
```

Iterating over the values

```go
for _, value := range m {
    fmt.Println("Value:", value)
}
```

#### Conclusion

Maps are powerful data structures in Go that allow you to store and manage key-value pairs efficiently. By understanding the basic operations and behavior of maps, you can effectively use them in various applications.

### Structs

In Go, a struct is a composite data type that groups together zero or more fields (properties) with differing data types. Structs are incredibly versatile and are often used as the foundation for creating complex data models in Go.

Anonymous Structs

Structs don't necessarily need to have a defined name. You can create an anonymous struct, declaring it on-the-fly:

```go
aDeveloper := struct{ name string }{name: "Wilson"}
fmt.Println(aDeveloper)
```

Structs are Value Types

By default, when you assign one struct to another or pass it to a function, the entire struct is copied. This behavior is consistent with Go's other value types like int and float64.

```go
anotherDeveloper := aDeveloper
anotherDeveloper.name = "Martin Fowler"
fmt.Println(aDeveloper)       // Wilson
fmt.Println(anotherDeveloper) // Martin Fowler
```

Passing Structs by Reference

If you don't want to work with a copy, you can pass a pointer to the struct:

```go
anotherDeveloper2 := &aDeveloper
anotherDeveloper2.name = "Martin Fowler"
fmt.Println(aDeveloper)        // Martin Fowler
fmt.Println(anotherDeveloper2) // Martin Fowler
```

Declaring Struct Types

```go
type Developer struct {
	Name  string
	Age   int
	Stack []string
}
dev := Developer{Name: "Wilson", Age: 30, Stack: []string{"Go Lang", "C#", "Python", "JS/TS", "Azure", ".Net", "React.js", "Next.js"}}
fmt.Println(dev)
```

Embedding

Go doesn't have the traditional class inheritance that you'd find in object-oriented languages, but it offers a type embedding mechanism. This allows you to include one struct type inside another, which promotes code reuse and a form of polymorphism.

```go
type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	CanFly bool
}
```

Embedded types can be initialized either explicitly:

```go
bird := Bird{Animal: Animal{Name: "Name", Origin: "BR"}, CanFly: true}
fmt.Println(bird)
```

Or implicitly:

```go
bird2 := Bird{}
bird2.Name = "Bird 2"
bird2.Origin = "US"
bird2.CanFly = true
fmt.Println(bird2)
```

#### Tags

Struct field tags are annotations attached to struct field declarations. They are commonly used for metadata, such as configuring serialization and validation methods:

```go
type City struct {
	Name  string `required max:"100" json:"name,omitempty"`
	State string `json:"state,omitempty"`
}

r := reflect.TypeOf(City{})
field, _ := r.FieldByName("Name")
fmt.Println(field.Tag) //required max:"100" json:"name,omitempty"
```

#### Conclusion

Structs in Go provide a flexible way to organize and model your data. They form the foundation for creating user-defined types and can be combined, embedded, or extended in versatile ways to suit the specific requirements of your program.

