package main

import (
	"fmt"
	"strconv"
)

//part 1
func apply_rules(value string) bool {
	digits := len(value) == 6
	adjacent := false
	ascending := true
	prev := 0
	for _, char := range value {
		curr := int(char - '0')
		if prev != 0 {
			if prev > curr {
				ascending = false
				break
			}
			if prev == curr {
				adjacent = true
			}
		}
		prev = curr
	}
	return digits && adjacent && ascending
}

//part 2
func apply_rules_2(value string) bool {
	digits := len(value) == 6
	ascending := true
	adjacent := false
	prev := 0
	adjcount := 0
	for _, char := range value {
		curr := int(char - '0')
		if prev != 0 {
			if prev > curr {
				ascending = false
				break
			}
			if prev == curr {
				adjcount += 1
			}
			if prev != curr {
				if adjcount == 1 {
					adjacent = true
				}
				adjcount = 0
			}
		}
		prev = curr
	}
	if adjcount == 1 {
		adjacent = true
	}
	return digits && adjacent && ascending
}

func main() {
	count := 0
	low := 134564
	high := 585159
	for i := low; i <= high; i++ {
		current := strconv.Itoa(i)
		if apply_rules_2(current) {
			count++
		}
	}

	fmt.Print(strconv.Itoa(count))
}
