package k3library

import (
	"fmt"
	"math"
	"os"
)

func Zbrac(start, end float64, f func(float64) float64) (x1, x2 float64, ok bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Fprintln(os.Stderr, x)
			ok = false
		}
	}()

	if start == end {
		panic("Zbrac: Bad initial range")
	} else if start > end {
		start, end = end, start
	}

	const POW float64 = 1.6
	const ITERATE int = 100
	fs := f(start)
	fe := f(end)
	x1, x2 = start, end

	for i := 0; i <= ITERATE; i++ {
		if fs*fe < 0 {
			ok = true
			return
		}
		if math.Abs(fs) < math.Abs(fe) {
			x1 += POW * (x1 - x2)
			fs = f(x1)
		} else {
			x2 += POW * (x2 - x1)
			fe = f(x2)
		}
	}
	return
}
