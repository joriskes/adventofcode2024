package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Pos struct {
	x int
	y int
}

type Antenna struct {
	pos       Pos
	frequency string
}

type Antinode struct {
	pos       Pos
	frequency string
}

func calcAntiNode(p1 Pos, p2 Pos) Pos {
	res := Pos{x: p1.x + (p1.x - p2.x), y: p1.y + (p1.y - p2.y)}
	return res
}

func dist(p1 Pos, p2 Pos) int {
	return int(math.Abs(float64(p2.x-p1.x)) + math.Abs(float64(p2.y-p1.y)))
}

func main() {
	fmt.Println("====== Day 8 ======")
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1 := 0
	part2 := 0

	y := 0
	width := 0
	height := 0
	antennas := []Antenna{}
	for scanner.Scan() {
		line := scanner.Text()
		char := strings.Split(line, "")
		for x := 0; x < len(char); x++ {
			if char[x] != "." {
				antennas = append(antennas, Antenna{Pos{x: x, y: y}, char[x]})
			}
		}
		y++
		width = len(char)
		height = y
	}

	antiNodes := map[string]Antinode{}
	for i := 0; i < len(antennas); i++ {
		for j := 0; j < len(antennas); j++ {
			if i != j {
				if antennas[i].frequency == antennas[j].frequency {
					nodePos := calcAntiNode(antennas[i].pos, antennas[j].pos)
					if nodePos.x > -1 && nodePos.y > -1 && nodePos.x < width && nodePos.y < height {
						antiNodes[strconv.Itoa(nodePos.x)+","+strconv.Itoa(nodePos.y)] = Antinode{calcAntiNode(antennas[i].pos, antennas[j].pos), antennas[i].frequency}
					}
				}
			}
		}
	}
	part1 = len(antiNodes)

	clear(antiNodes)
	for i := 0; i < len(antennas); i++ {
		for j := 0; j < len(antennas); j++ {
			if i != j {
				if antennas[i].frequency == antennas[j].frequency {
					a1 := antennas[i].pos
					a2 := antennas[j].pos
					nodePos := calcAntiNode(a1, a2)
					for nodePos.x > -1 && nodePos.y > -1 && nodePos.x < width && nodePos.y < height {
						antiNodes[strconv.Itoa(nodePos.x)+","+strconv.Itoa(nodePos.y)] = Antinode{calcAntiNode(antennas[i].pos, antennas[j].pos), antennas[i].frequency}
						if dist(nodePos, a1) < dist(nodePos, a2) {
							a2 = a1
						}
						a1 = nodePos
						nodePos = calcAntiNode(a1, a2)
					}
				}
			}
		}
	}
	// The layout is correct, but since antennas and antinodes can be (and mostly are) in the same spot for P2
	// i calculate them while (optionally printing the map)
	printMap := false
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			done := false
			for i := 0; i < len(antennas); i++ {
				if antennas[i].pos.x == x && antennas[i].pos.y == y {
					part2++
					done = true
					if printMap {
						fmt.Print(antennas[i].frequency)
					}
				}
			}
			if !done {
				_, ok := antiNodes[strconv.Itoa(x)+","+strconv.Itoa(y)]
				if ok {
					if printMap {
						fmt.Print("#")
					}
					part2++
				} else {
					if printMap {
						fmt.Print(".")
					}
				}
			}
		}
		if printMap {
			fmt.Println()
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 8 time:", time.Since(start))
}
