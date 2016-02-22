package k3library

import (
	"fmt"
)

func Sgn(x float64) (ans float64, err error) {

	if x > 0 {
		ans = 1
	} else if x < 0 {
		ans = -1
	} else {
		ans = 0
	}

	return
}
