// function
package pac

import (
	"fmt"
	"math"
)

func Unit8Format(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	} else {
		return 0, fmt.Errorf("%d is out of the rang", x)
	}
}
