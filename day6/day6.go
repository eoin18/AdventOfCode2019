package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("C:\\workspace\\AdventOfCode2019\\day6\\input.txt")
	defer file.Close()

	var grid map[string]string
	grid = make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ")")
		grid[input[1]] = input[0]
	}

	orbits := 0
	// Part 1
	//for _ ,v := range grid {
	//	orbits += 1
	//	current := grid[v]
	//	for current != "" {
	//		orbits += 1
	//		current=grid[current]
	//	}
	//}

	// Part 2

	start := "YOU"
	var mypath map[string]int
	mypath = make(map[string]int)

	current := grid[start]
	for current != "" {
		orbits += 1
		mypath[current] = orbits
		current = grid[current]
	}

	var santapath map[string]int
	santapath = make(map[string]int)

	orbits = 0
	start = "SAN"
	current = grid[start]
	for current != "" {
		orbits += 1
		santapath[current] = orbits
		current = grid[current]
	}

	shortestpath := 0
	for k, v := range mypath {
		if santapath[k] != 0 {
			length := (v - 1) + (santapath[k] - 1)
			if shortestpath == 0 || length < shortestpath {
				shortestpath = length
			}
		}
	}

	fmt.Print(strconv.Itoa(shortestpath))
}
