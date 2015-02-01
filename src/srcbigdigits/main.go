// srcbigdigits project main.go
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var i int
	var x1 uint8
	i, _ = strconv.Atoi(os.Args[1])
	x1, _ = Unit8format(i)
	fmt.Println(x1)
	fmt.Println("Hello World!")
}
func Unit8format(x int) (uint8, error) {
	if x >= 0 && x <= math.MaxInt8 {
		return uint8(x), nil
	}

	return 0, fmt.Errorf("%d is out of the uint8", x)
}
