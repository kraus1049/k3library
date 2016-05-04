package k3library

func GCD(x, y int) int {
	if y > x {
		x, y = y, x
	}

	for {
		ans := x % y

		if ans == 0 {
			return y
		} else {
			y = ans
		}
	}
}
