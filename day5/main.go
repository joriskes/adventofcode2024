package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Rule struct {
	needsToBeBefore []int
}

func strToInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func isInOrder(update []int, rules map[int]Rule) bool {
	for i, u := range update {
		m, ok := rules[u]
		if ok {
			if len(m.needsToBeBefore) > 0 {
				for _, u := range update[0:i] {
					if slices.Contains(m.needsToBeBefore, u) {
						return false
					}
				}
			}
		}
	}
	return true
}

func fixOrder(update []int, rules map[int]Rule) []int {
	newOrder := update
	move := 0
	for move > -1 {
		move = -1
		i := 0
	findMove:
		for move == -1 && i < len(newOrder) {
			r, ok := rules[newOrder[i]]
			if ok && len(r.needsToBeBefore) > 0 {
				for _, c := range r.needsToBeBefore {
					if slices.Contains(newOrder[0:i], c) {
						move = i
						break findMove
					}
				}
			}
			i++
		}
		if move > -1 {
			// Found failure, move that number left one and retry
			c := newOrder[move-1]
			newOrder[move-1] = newOrder[move]
			newOrder[move] = c
		}
	}

	return update
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

	rules := map[int](Rule){}
	updates := [][]int{}
	emptyLineFound := false
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			emptyLineFound = true
		} else {
			if !emptyLineFound {
				s := strings.Split(t, "|")
				n1 := strToInt(s[0])
				n2 := strToInt(s[1])

				m, ok := rules[n1]
				if ok {
					m.needsToBeBefore = append(rules[n1].needsToBeBefore, n2)
					rules[n1] = m
				} else {
					rules[n1] = Rule{
						needsToBeBefore: []int{n2},
					}
				}
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
		if isInOrder(updates[u], rules) {
			part1 += updates[u][len(updates[u])/2]
		} else {
			updates[u] = fixOrder(updates[u], rules)
			part2 += updates[u][len(updates[u])/2]
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 5 time:", time.Since(start))
}
