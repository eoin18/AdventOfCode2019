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

type Vertex struct {
	X int
	Y int
}

type Node struct {
	Visitors uint
	Steps int
}

//Part1
func manhattan_distance(v1 Vertex, v2 Vertex) int{
	return int(math.Abs(float64(v1.X-v2.X)) + math.Abs(float64(v1.Y-v2.Y)))
}

func update_node(current Vertex, grid map[Vertex]Node, steps int, count uint) {
	if grid[current].Visitors == 0 {
		grid[current] = Node{Visitors:0|count, Steps:steps}
	} else if grid[current].Visitors & count == 1 {
		grid[current] = Node{Visitors:grid[current].Visitors|count, Steps:steps}
	} else {
		grid[current] = Node{Visitors:grid[current].Visitors|count, Steps:grid[current].Steps + steps}
	}
}

func run_path(start Vertex, path []string, grid map[Vertex]Node, count uint){
	current := start
	steps := 0
	for i := 0 ; i < len(path) ; i++ {
		instruction := path[i]
		dir := instruction[0]
		distance, _ := strconv.Atoi(instruction[1:])
		if dir == 'U' {
			for j := 0 ; j < distance ; j++{
				steps += 1
				current = Vertex{X:current.X, Y:current.Y + 1}
				update_node(current, grid, steps, count)
			}
		} else if dir == 'D' {
			for j := 0 ; j < distance ; j++{
				steps += 1
				current = Vertex{X:current.X, Y:current.Y - 1}
				update_node(current, grid, steps, count)
			}
		} else if dir == 'R' {
			for j := 0 ; j < distance ; j++{
				steps += 1
				current = Vertex{X:current.X + 1, Y:current.Y}
				update_node(current, grid, steps, count)
			}
		} else if dir == 'L' {
			for j := 0 ; j < distance ; j++{
				steps += 1
				current = Vertex{X:current.X - 1, Y:current.Y}
				update_node(current, grid, steps, count)
			}
		}
	}
}

func main()  {
	file, err := os.Open("C:\\workspace\\AdventOfCode2019\\day3\\input.txt")
	check(err)
	defer file.Close()

	var grid map[Vertex]Node
	grid = make(map[Vertex]Node)
	scanner := bufio.NewScanner(file)
	var i uint = 1
	var j uint = 1
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ",")
		run_path(Vertex{X:0, Y:0}, input, grid, i)
		j = j | i
		i = i << 1
	}
	shortest := 0
	for _,v := range grid {
		if v.Visitors == j {
			//dist := manhattan_distance(Vertex{X:0, Y:0}, k)
			steps := v.Steps
			if shortest == 0 || steps < shortest {
				shortest = steps
			}
		}
	}
	fmt.Print(strconv.Itoa(shortest))
}
