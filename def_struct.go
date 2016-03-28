package k3library

import (
	"fmt"
)

type Vec struct {
	V   []float64
	Row int
}

type Mat struct {
	M        []Vec
	Col, Row int
}

func NewVec(row int) Vec {
	if row < 0 {
		row = 0
	}

	v := make([]float64, row)
	return Vec{v, row}
}

func NewMat(col, row int) Mat {
	if col <= 0 {
		col = 1
	}

	m := make([]Vec, col)
	for i := range m {
		m[i] = NewVec(row)
	}
	return Mat{m, col, row}
}

func (v Vec) String() string {
	return fmt.Sprint(v.V)
}

func (m Mat) String() string {
	str := ""
	for i, v := range m.M {
		str += "|"
		for _, item := range v.V {
			str += fmt.Sprintf(" %v", item)
		}
		if i != m.Col {
			str += " |\n"
		} else {
			str += " |"
		}
	}

	return str
}
