package k3library

func Csv2func(filepath string, parser func(string) (map[string]string, [][2]float64, error)) (f func(float64) (float64, error), err error) {
}
