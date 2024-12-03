package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	r := regexp.MustCompile(`((mul|do|don't)\((\d{1,3})?,?(\d{1,3})?\))`)
	matches := r.FindAllStringSubmatch(string(file), -1)

	part1 := 0
	part2 := 0
	enabled := true
	for _, m := range matches {
		instruction, m1, m2 := m[2], m[3], m[4]
		switch instruction {
		case "mul":
			a, _ := strconv.Atoi(m1)
			b, _ := strconv.Atoi(m2)
			part1 += a * b
			if enabled {
				part2 += a * b
			}
		case "do":
			enabled = true
		case "don't":
			enabled = false
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
