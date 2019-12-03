package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func execute_intercode(noun, verb int, intercode []int) int {
	intercode[1] = noun
	intercode[2] = verb
	for i := 0 ; i < len(intercode) ; i+= 4 {
		opcode := intercode[i]
		if opcode == 99 {
			break
		} else {
			operand1 := intercode[i+1]
			operand2 := intercode[i+2]
			target := intercode[i+3]
			if opcode == 1 {
				intercode[target] = intercode[operand1] + intercode[operand2]
			}
			if opcode == 2 {
				intercode[target] = intercode[operand1] * intercode[operand2]
			}
		}
	}
	return intercode[0]
}

func convert_to_int_array(strcode []string) []int {
	result := make([]int, len(strcode))
	i := 0
	for i < len(strcode){
		data, err := strconv.Atoi(strcode[i])
		check(err)
		result[i] = data
		i++
	}
	return result
}

func main()  {
	file, err := os.Open("C:\\workspace\\AdventOfCode2019\\day2\\input.txt")
	check(err)
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ",")
		i := 0
		j := 0
		for i < 99 {
			for j < 99 {
				intercode := convert_to_int_array(input)
				result = execute_intercode(i, j, intercode)
				if result == 19690720 {
					fmt.Print(strconv.Itoa(100*i + j))
					return
				}
				j++
			}
			j = 0
			i++
		}
	}
}
