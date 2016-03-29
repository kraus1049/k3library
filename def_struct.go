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
		if i == 0 {
			str += "\n|"
		} else {
			str += "|"
		}
		for _, item := range v.V {
			str += fmt.Sprintf(" %v", item)
		}
		if i != m.Col-1 {
			str += " |\n"
		} else {
			str += " |"
		}
	}

	return str
}

func (v *Vec) Set(x []float64) {
	if len(x) < v.Row {
		zero := make([]float64, v.Row-len(x))
		x = append(x, zero...)
	}

	v.V = x[0:v.Row]
}

func (m *Mat) Set(x [][]float64) {

	if len(x) >= m.Col {
		for i := 0; i < m.Col; i++ {
			m.M[i].Set(x[i])
		}
	} else {
		for i := 0; i < len(x); i++ {
			m.M[i].Set(x[i])
		}
	}

}

func (v *Vec) At(i int) float64 {
	return v.V[i]
}

func (m *Mat) At(i, j int) float64 {
	return m.M[i].V[j]
}

func (v *Vec) Write(i int, num float64) {
	if i >= 0 {
		v.V[i] = num
	} else {
		for j := 0; j < v.Row; j++ {
			v.V[j] = num
		}
	}
}

func (m *Mat) Write(i, j int, num float64) {
	if i >= 0 {
		if j >= 0 {
			m.M[i].V[j] = num
		} else {
			m.M[i].Write(-1, num)
		}
	} else {
		if j >= 0 {
			for k := 0; k < m.Col; k++ {
				m.M[k].V[j] = num
			}
		} else {
			idx := 0
			if m.Col < m.Row {
				idx = m.Col
			} else {
				idx = m.Row
			}

			for k := 0; k < idx; k++ {
				m.M[k].V[k] = num
			}
		}
	}
}
