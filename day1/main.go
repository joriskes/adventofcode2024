package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("====== Day 1 ======")
	start := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var leftList []float64
	var rightList []float64

	for scanner.Scan() {
		v := strings.Split(scanner.Text(), "   ")
		l, err := strconv.ParseFloat(v[0], 64)
		if err != nil {
			panic(err)
		}
		r, err := strconv.ParseFloat(v[1], 64)
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, l)
		rightList = append(rightList, r)
	}

	// Part 1

	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})
	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	part1 := 0.0

	for i := 0; i < len(leftList); i++ {
		part1 += math.Abs(leftList[i] - rightList[i])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", int(part1))

	// Part 2
	part2 := 0.0
	for i := 0; i < len(leftList); i++ {
		foundCount := 0.0
		for j := 0; j < len(rightList); j++ {
			if leftList[i] == rightList[j] {
				foundCount++
			}
		}
		part2 += leftList[i] * foundCount
	}
	fmt.Println("Part 2:", int(part2))

	fmt.Println("⏱️ Day 1 time:", time.Since(start))
}
