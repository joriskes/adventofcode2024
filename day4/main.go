package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func countInLine(line string) int {
	return strings.Count(line, "XMAS") + strings.Count(line, "SAMX")
}

func main() {
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
	size := width + height
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
	fmt.Println("Part 2:", part2)
}
