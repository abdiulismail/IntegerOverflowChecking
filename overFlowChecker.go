package main

import (
	"fmt"
	"math"
)
import "errors"

var ErrOverFlow = errors.New("integer overflow")

func Add32(left, right int32) (int32, error) {
	if right > 0 {
		if left > math.MaxInt32-right {
			return 0, ErrOverFlow
		}
	} else {
		if left < math.MaxInt32-right {
			return 0, ErrOverFlow
		}
	}
	return left + right, nil
}

func main() {

	var a int32 = 2147483327
	var b int32 = 2147483327

	c, err := Add32(a, b)

	if err != nil {
		//handle overflow
		fmt.Println(err, a, b, c)
	}

}
