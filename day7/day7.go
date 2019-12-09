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

type Amplifier struct {
	intercode []int
	index     int
}

func (amplifier *Amplifier) execute_intercode(input1, input2 int) (int, bool) {
	i := amplifier.index
	inputs := 0
	for {
		instruction := amplifier.intercode[i]
		opcode := instruction % 100
		//mode3 := (instruction / 10000) % 10
		mode2 := (instruction / 1000) % 10
		mode1 := (instruction / 100) % 10
		if opcode == 99 {
			i += 1
			amplifier.index = i
			return -1, true
		} else {
			if opcode == 1 {
				operand1, operand2 := calc_operands(amplifier.intercode, i, mode1, mode2)
				target := amplifier.intercode[i+3]
				amplifier.intercode[target] = operand1 + operand2
				i += 4
			}
			if opcode == 2 {
				operand1, operand2 := calc_operands(amplifier.intercode, i, mode1, mode2)
				target := amplifier.intercode[i+3]
				amplifier.intercode[target] = operand1 * operand2
				i += 4
			}
			if opcode == 3 {
				target := amplifier.intercode[i+1]
				if inputs == 0 {
					amplifier.intercode[target] = input1
					inputs++
				} else {
					amplifier.intercode[target] = input2
				}
				i += 2
			}
			if opcode == 4 {
				target := amplifier.intercode[amplifier.intercode[i+1]]
				if mode1 == 0 {
					abs := int(math.Abs(float64(amplifier.intercode[i+1])))
					target = amplifier.intercode[abs]
				}
				i += 2
				amplifier.index = i
				return target, false
			}
			if opcode == 5 {
				operand1, operand2 := calc_operands(amplifier.intercode, i, mode1, mode2)
				if operand1 != 0 {
					i = operand2
				} else {
					i += 3
				}
			}
			if opcode == 6 {
				operand1, operand2 := calc_operands(amplifier.intercode, i, mode1, mode2)
				if operand1 == 0 {
					i = operand2
				} else {
					i += 3
				}
			}
			if opcode == 7 {
				operand1, operand2 := calc_operands(amplifier.intercode, i, mode1, mode2)
				target := amplifier.intercode[i+3]
				if operand1 < operand2 {
					amplifier.intercode[target] = 1
				} else {
					amplifier.intercode[target] = 0
				}
				i += 4
			}
			if opcode == 8 {
				operand1, operand2 := calc_operands(amplifier.intercode, i, mode1, mode2)
				target := amplifier.intercode[i+3]
				if operand1 == operand2 {
					amplifier.intercode[target] = 1
				} else {
					amplifier.intercode[target] = 0
				}
				i += 4
			}
		}
	}
}

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
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
	file, err := os.Open("C:\\workspace\\AdventOfCode2019\\day7\\input.txt")
	check(err)
	defer file.Close()

	highest := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ",")
		Perm([]rune("56789"), func(runes []rune) {
			in := 0
			out := 0
			var amplifiers map[int]Amplifier
			amplifiers = make(map[int]Amplifier)
			for i := 0; i < 5; i++ {
				amplifiers[i] = Amplifier{intercode: convert_to_int_array(input), index: 0}
			}
			halt := false
			loops := 0
			//First loop with phases
			for i, char := range runes {
				amplifier := amplifiers[i]
				out, _ = amplifier.execute_intercode(int(char-'0'), in)
				amplifiers[i] = amplifier
				in = out
			}
			//Continue with feedback loops until halt
			for halt == false {
				for i := 0; i < 5; i++ {
					amplifier := amplifiers[i]
					out, halt = amplifier.execute_intercode(in, in)
					amplifiers[i] = amplifier
					if halt {
						break
					}
					in = out
				}
				loops++
			}
			if in > highest {
				highest = in
			}
		})
	}
	fmt.Print(strconv.Itoa(highest))
}
