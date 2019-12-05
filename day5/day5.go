package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func execute_intercode(input int, intercode []int) int {
	i := 0
	for {
		instruction := intercode[i]
		opcode := instruction % 100
		//mode3 := (instruction / 10000) % 10
		mode2 := (instruction / 1000) % 10
		mode1 := (instruction / 100) % 10
		if opcode == 99 {
			return intercode[0]
		} else {
			if opcode == 1 {
				operand1, operand2 := calc_operands(intercode, i, mode1, mode2)
				target := intercode[i+3]
				intercode[target] = operand1 + operand2
				i += 4
			}
			if opcode == 2 {
				operand1, operand2 := calc_operands(intercode, i, mode1, mode2)
				target := intercode[i+3]
				intercode[target] = operand1 * operand2
				i += 4
			}
			if opcode == 3 {
				target := intercode[i+1]
				intercode[target] = input
				i += 2
			}
			if opcode == 4 {
				target := intercode[intercode[i+1]]
				if mode1 == 0 {
					abs := int(math.Abs(float64(intercode[i+1])))
					target = intercode[abs]
				}
				fmt.Print(strconv.Itoa(target) + "\n")
				i += 2
			}
			if opcode == 5 {
				operand1, operand2 := calc_operands(intercode, i, mode1, mode2)
				if operand1 != 0 {
					i = operand2
				} else {
					i += 3
				}
			}
			if opcode == 6 {
				operand1, operand2 := calc_operands(intercode, i, mode1, mode2)
				if operand1 == 0 {
					i = operand2
				} else {
					i += 3
				}
			}
			if opcode == 7 {
				operand1, operand2 := calc_operands(intercode, i, mode1, mode2)
				target := intercode[i+3]
				if operand1 < operand2 {
					intercode[target] = 1
				} else {
					intercode[target] = 0
				}
				i += 4
			}
			if opcode == 8 {
				operand1, operand2 := calc_operands(intercode, i, mode1, mode2)
				target := intercode[i+3]
				if operand1 == operand2 {
					intercode[target] = 1
				} else {
					intercode[target] = 0
				}
				i += 4
			}
		}
	}
}

func calc_operands(intercode []int, i int, mode1 int, mode2 int) (int, int) {
	operand1 := intercode[i+1]
	if mode1 == 0 {
		abs := int(math.Abs(float64(intercode[i+1])))
		operand1 = intercode[abs]
	}
	operand2 := intercode[i+2]
	if mode2 == 0 {
		abs := int(math.Abs(float64(intercode[i+2])))
		operand2 = intercode[abs]
	}
	return operand1, operand2
}

func convert_to_int_array(strcode []string) []int {
	result := make([]int, len(strcode))
	i := 0
	for i < len(strcode) {
		data, err := strconv.Atoi(strcode[i])
		check(err)
		result[i] = data
		i++
	}
	return result
}

func main() {
	file, err := os.Open("C:\\workspace\\AdventOfCode2019\\day5\\input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ",")
		intercode := convert_to_int_array(input)
		execute_intercode(5, intercode)
	}
}
