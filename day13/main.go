package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Pos struct {
	x int
	y int
}

type Machine struct {
	AOffset Pos
	BOffset Pos
	Prize   Pos
}

func strToInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return res
}

func calculateTokenCost(offset Pos, finish Pos) int {
	x := 0
	y := 0
	steps := 0
	for x < finish.x && y < finish.y {
		x += offset.x
		y += offset.y
		steps++
	}
	if x != finish.x || y != finish.y {
		return -1
	}
	return steps
}

func main() {
	fmt.Println("====== Day 13 ======")
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1 := 0
	part2 := 0

	machines := []Machine{}
	for scanner.Scan() {
		t0 := scanner.Text()
		scanner.Scan()
		t1 := scanner.Text()
		scanner.Scan()
		t2 := scanner.Text()
		scanner.Scan()

		re := regexp.MustCompile(`X[+|=]{1}(\d+),\sY[+|=]{1}(\d+)`)

		a := re.FindSubmatch([]byte(t0))
		b := re.FindSubmatch([]byte(t1))
		p := re.FindSubmatch([]byte(t2))

		machines = append(machines, Machine{
			AOffset: Pos{
				strToInt(string(a[1][:])),
				strToInt(string(a[2][:])),
			},
			BOffset: Pos{
				strToInt(string(b[1][:])),
				strToInt(string(b[2][:])),
			},
			Prize: Pos{
				strToInt(string(p[1][:])),
				strToInt(string(p[2][:])),
			},
		})
	}

	for _, m := range machines {
		resA := calculateTokenCost(m.AOffset, m.Prize)
		resB := calculateTokenCost(m.BOffset, m.Prize)

		fmt.Println(m, resA, resB)

		if resA < 0 {
			part1 += resB
		} else if resB < 0 {
			part1 += resA
		} else {
			if resB*3 < resA {
				part1 += resB
			} else {
				part1 += resA
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 13 time:", time.Since(start))
}
