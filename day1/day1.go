package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calc_fuel(i int) int {
	return int(math.Floor(float64(i) / 3)) -2
}

func main()  {
	file, err := os.Open("C:\\workspace\\AdventOfCode2019\\day1\\input.txt")
	check(err)
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		check(err)
		fuel := calc_fuel(i)
		for fuel > 0 {
			result += fuel
			fuel = calc_fuel(fuel)
		}
	}
	fmt.Print(strconv.Itoa(result))
}
