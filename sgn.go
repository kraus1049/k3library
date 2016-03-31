package k3library

func Sgn(x float64) (float64, error) {
	ans := 0.0

	if x > 0 {
		ans = 1
	} else if x < 0 {
		ans = -1
	}

	return ans, nil
}
