// test1 project main.go
package main

import "fmt"
import "test1/pac"

func main() {
	fmt.Println("Hello World!")
	x, _ := pac.Unit8Format(123)
	fmt.Println(x)
}
