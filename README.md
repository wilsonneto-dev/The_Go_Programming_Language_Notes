Studying GO by the book and documentation

- [ ] 🟡 [https://www.amazon.com.br/Go-Programming-Language-Brian-Kernighan](https://www.amazon.com.br/Go-Programming-Language-Brian-Kernighan/dp/0134190440/ref=sr_1_1?crid=3O3XQFDCFYMID&keywords=the+go+programming+language&qid=1695784756&sprefix=The+GO+pro%2Caps%2C235)
- [ ] 🟡 https://www.youtube.com/watch?v=jFfo23yIWac
- [ ] https://www.youtube.com/watch?v=YS4e4q9oBaU

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

