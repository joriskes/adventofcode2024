package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func remove(s []int, index int) []int {
	// Make copy
	ret := make([]int, 0)
	// Append all before index
	ret = append(ret, s[:index]...)
	// Append all after index
	return append(ret, s[index+1:]...)
}

func isSafe(numbers []int) bool {
	direction := 0
	safe := true
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		if diff == 0 || math.Abs(float64(diff)) > 3 {
			safe = false
			break
		}
		if diff > 0 && direction < 0 {
			safe = false
			break
		}
		if diff < 0 && direction > 0 {
			safe = false
			break
		}
		direction = diff
	}
	return safe
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

	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), " ")
		numbers := make([]int, len(lineSplit))
		for i := 0; i < len(lineSplit); i++ {
			numbers[i], _ = strconv.Atoi(lineSplit[i])
		}

		safe := isSafe(numbers)
		if safe {
			part1++
			part2++
		} else {
			for i := 0; i < len(lineSplit); i++ {
				try := remove(numbers, i)
				safe2 := isSafe(try)
				if safe2 {
					part2++
					break
				}
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
