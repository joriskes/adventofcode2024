package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func countInLine(line string) int {
	return strings.Count(line, "XMAS") + strings.Count(line, "SAMX")
}

func charAt(horizStrings []string, x int, y int) string {
	if x < 0 || y < 0 || len(horizStrings) <= y || len(horizStrings[y]) <= x {
		return ""
	}
	return horizStrings[y][x : x+1]
}

func main() {
	fmt.Println("====== Day 4 ======")
	start := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1 := 0
	part2 := 0

	var horizStrings []string
	var vertStrings []string

	// Joris: this is slow
	for scanner.Scan() {
		horizStrings = append(horizStrings, scanner.Text())
	}

	for lnum, horizLine := range horizStrings {
		part1 += countInLine(horizLine)

		if lnum == 0 {
			vertStrings = make([]string, len(horizLine))
		}
		for i := 0; i < len(horizLine); i++ {
			// Build vertStrings in reverse order (for diagonal use later)
			vertStrings[i] = horizLine[i:i+1] + vertStrings[i]
		}
	}

	for _, vertLine := range vertStrings {
		part1 += countInLine(vertLine)
	}

	width := len(vertStrings)
	height := len(horizStrings)
	size := width + height - 1
	horizDiagStrings := make([]string, size)
	vertDiagStrings := make([]string, size)

	for x := 0; x < width; x++ {
		for y := height - 1; y >= 0; y-- {
			horizDiagStrings[x+y] += horizStrings[y][x : x+1]
			vertDiagStrings[x+y] += vertStrings[x][y : y+1]
		}
	}

	for _, horizDiagLine := range horizDiagStrings {
		part1 += countInLine(horizDiagLine)
	}

	for _, vertDiagLine := range vertDiagStrings {
		part1 += countInLine(vertDiagLine)
	}

	fmt.Println("Part 1:", part1)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := charAt(horizStrings, x, y)
			if c == "A" {
				c1 := charAt(horizStrings, x-1, y-1) + charAt(horizStrings, x+1, y+1)
				c2 := charAt(horizStrings, x+1, y-1) + charAt(horizStrings, x-1, y+1)
				if (c1 == "MS" || c1 == "SM") && (c2 == "MS" || c2 == "SM") {
					part2++
				}

			}
		}

	}
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 4 time:", time.Since(start))
}
