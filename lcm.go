package k3library

func LCM(x, y int) int {
	return x * y / GCD(x, y)
}
