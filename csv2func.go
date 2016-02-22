package k3library

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func Csv2func(filename string) (f func(float64) float64, ok bool) {
	read_file, _ := os.OpenFile(filename, os.O_RDONLY, 0666)
	reader := csv.NewReader(read_file)

	var aa []string

	for {
		col, err := reader.Read()

		if err == io.EOF {
			break
		}

		aa = append(col)
	}

}
