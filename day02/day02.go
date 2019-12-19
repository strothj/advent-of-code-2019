package day02

import (
	"strconv"
	"strings"
)

func day02(input string, programAlarm bool, nounVerbSearch bool) int {
	inputValues := strings.Split(input, ",")
	var memory []int

	loadProgram := func(noun, verb int) {
		memory = make([]int, 0, len(inputValues))
		for _, v := range inputValues {
			memoryValue, _ := strconv.Atoi(v)
			memory = append(memory, memoryValue)
		}
		if noun != -1 {
			memory[1] = noun
		}
		if verb != -1 {
			memory[2] = verb
		}
	}

	executeProgram := func() int {
		for instructionPointer := 0; instructionPointer < len(memory); instructionPointer += 4 {
			opCode := memory[instructionPointer]
			parameter0Address := 0
			parameter1Address := 0
			resultAddress := 0

			switch opCode {
			case 1:
				fallthrough
			case 2:
				if instructionPointer+3 >= len(memory) {
					return -1
				}
				parameter0Address = memory[instructionPointer+1]
				parameter1Address = memory[instructionPointer+2]
				resultAddress = memory[instructionPointer+3]
				break
			case 99:
				return memory[0]
			default:
				return -2
			}

			if parameter0Address >= len(memory) || parameter1Address >= len(memory) || resultAddress >= len(memory) {
				return -3
			}

			switch opCode {
			case 1:
				memory[resultAddress] = memory[parameter0Address] + memory[parameter1Address]
				break
			case 2:
				memory[resultAddress] = memory[parameter0Address] * memory[parameter1Address]
				break
			}
		}

		return -1
	}

	for j := 0; j <= 99; j++ {
		for k := 0; k <= 99; k++ {
			noun := -1
			verb := -1
			if programAlarm {
				noun = 12
				verb = 2
			}
			if nounVerbSearch {
				noun = j
				verb = k
			}
			loadProgram(noun, verb)
			result := executeProgram()
			if !nounVerbSearch {
				return result
			}
			if result == 19690720 {
				return 100*noun + verb
			}
		}
	}

	return -1
}
