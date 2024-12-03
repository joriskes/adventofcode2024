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
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(string(file), -1)

	part1 := 0
	for _, m := range matches {
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		part1 += a * b
	}

	fmt.Println("Part 1:", part1)
	//fmt.Println("Part 2:", part2)
}
