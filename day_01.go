package advent

import (
	"fmt"
	"io"
	"strings"
)

func day01(input string) int {
	r := strings.NewReader(input)
	var value int
	var sum int
	for {
		_, err := fmt.Fscanln(r, &value)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		sum += value/3 - 2
	}
	return sum
}
