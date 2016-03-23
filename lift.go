package k3library

func Lift(f func(float64) float64) func(float64) (float64, error) {
	g := func(x float64) (float64, error) { return f(x), nil }
	return g
}
