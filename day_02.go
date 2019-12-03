package advent

import (
	"strconv"
	"strings"
)

func day02(input string, programAlarm bool) int {
	inputValues := strings.Split(input, ",")
	memory := make([]int, 0, len(inputValues))
	for _, v := range inputValues {
		memoryValue, _ := strconv.Atoi(v)
		memory = append(memory, memoryValue)
	}
	if programAlarm {
		memory[1] = 12
		memory[2] = 2
	}

	pc := 0
	exit := false
	for {
		switch memory[pc] {
		case 1:
			memory[memory[pc+3]] = memory[memory[pc+1]] + memory[memory[pc+2]]
			break
		case 2:
			memory[memory[pc+3]] = memory[memory[pc+1]] * memory[memory[pc+2]]
			break
		case 99:
			exit = true
			break
		}
		if exit {
			break
		}
		pc += 4
	}

	return memory[0]
}
