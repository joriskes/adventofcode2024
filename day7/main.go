package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func testOperators(numbers []uint64, runningTotal uint64, wantedTotal uint64, operatorList []string) bool {
	// Numbers go up BRRR
	if runningTotal > wantedTotal {
		return false
	}
	if len(numbers) < 1 {
		return runningTotal == wantedTotal
	}
	n1 := numbers[0]
	res := false
	for _, operator := range operatorList {
		switch operator {
		case "+":
			if testOperators(numbers[1:], runningTotal+n1, wantedTotal, operatorList) {
				res = true
			}
		case "*":
			if testOperators(numbers[1:], runningTotal*n1, wantedTotal, operatorList) {
				res = true
			}
		default:
			fmt.Println("Unknown operator")
			return false
		}
	}
	return res
}

func main() {
	fmt.Println("====== Day 7 ======")
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var part1 uint64 = 0
	var part2 uint64 = 0

	for scanner.Scan() {
		t := scanner.Text()

		ar := strings.Split(t, ": ")
		total, err := strconv.ParseUint(ar[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		operators := []string{
			"*",
			"+",
		}
		ar2 := strings.Split(ar[1], " ")
		numbers := []uint64{}
		for _, v := range ar2 {
			n, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, n)
		}

		if testOperators(numbers[1:], numbers[0], total, operators) {
			part1 += total
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 7 time:", time.Since(start))
}
