package k3library

func LinearFunc(v Vec) func(float64) float64 {
	return linearFunc(v, 0)
}

func linearFunc(v Vec, idx int) func(float64) float64 {
	if v.row-idx == 2 {
		a := v.At(idx + 1)
		b := v.At(idx)
		return func(x float64) float64 {
			return a*x + b
		}
	} else if v.row-idx == 1 {
		return func(x float64) float64 {
			return v.At(idx)
		}
	}

	a := linearFunc(v, idx+1)
	b := v.At(idx)

	return func(x float64) float64 {
		return a(x)*x + b
	}

}
