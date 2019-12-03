package advent

import (
	"fmt"
	"io"
	"strings"
)

func day01(input string, calculateFuelMass bool) int {
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
		fuelMass := value
		for {
			fuelMass = fuelMass/3 - 2
			if fuelMass < 0 {
				break
			}
			sum += fuelMass
			if !calculateFuelMass {
				break
			}
		}
	}
	return sum
}
