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

type Rule struct {
	num1 int
	num2 int
}

func strToInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func isInOrder(update []int, rules []Rule) bool {
	for rule := range rules {
		i1 := -1
		i2 := -1

		for i, v := range update {
			if v == rules[rule].num1 {
				i1 = i
			}
			if v == rules[rule].num2 {
				i2 = i
			}
		}

		if i1 > -1 && i2 > -1 {
			// Both found, check order
			if i1 > i2 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println("====== Day 5 ======")
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1 := 0
	part2 := 0

	rules := [](Rule){}
	updates := [][]int{}
	emptyLineFound := false
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			emptyLineFound = true
		} else {
			if !emptyLineFound {
				s := strings.Split(t, "|")
				rules = append(rules, Rule{
					num1: strToInt(s[0]),
					num2: strToInt(s[1]),
				})
			} else {
				s := strings.Split(t, ",")
				ar := make([]int, len(s))
				for i, v := range s {
					ar[i] = strToInt(v)
				}
				updates = append(updates, ar)
			}
		}
	}

	for u := 0; u < len(updates); u++ {
		fmt.Println(updates[u])
		if isInOrder(updates[u], rules) {
			fmt.Println("In order")
			part1 += updates[u][len(updates[u])/2]
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 5 time:", time.Since(start))
}
