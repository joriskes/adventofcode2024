package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var blinkCache = map[string]int{}

func blinkNumber(number int, timesLeft int) int {
	if c, ok := blinkCache[strconv.Itoa(number)+","+strconv.Itoa(timesLeft)]; ok {
		return c
	}

	stoneCount := 1
	num := number
	for j := 0; j < timesLeft; j++ {
		numWidth := len(strconv.Itoa(num))
		if num == 0 {
			num = 1
		} else if numWidth%2 == 0 {
			leftNum, _ := strconv.Atoi(strconv.Itoa(num)[0 : numWidth/2])
			rightNum, _ := strconv.Atoi(strconv.Itoa(num)[numWidth/2:])
			stoneCount += blinkNumber(leftNum, timesLeft-j-1)
			num = rightNum
		} else {
			num *= 2024
		}
	}
	blinkCache[strconv.Itoa(number)+","+strconv.Itoa(timesLeft)] = stoneCount
	return stoneCount
}

func main() {
	fmt.Println("====== Day 11 ======")
	start := time.Now()
	b, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	part1 := 0
	part2 := 0

	in := strings.Split(strings.TrimSpace(string(b)), " ")

	for i := 0; i < len(in); i++ {
		num, err := strconv.Atoi(in[i])
		if err != nil {
			fmt.Println(err)
			break
		}
		part1 += blinkNumber(num, 25)
	}

	for i := 0; i < len(in); i++ {
		num, err := strconv.Atoi(in[i])
		if err != nil {
			fmt.Println(err)
			break
		}
		part2 += blinkNumber(num, 75)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 11 time:", time.Since(start))
}
