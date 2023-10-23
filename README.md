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

### Visibility

- Lower Case first letter: Package Scope
- Upper Case first letter: Export from Package / Exposed

