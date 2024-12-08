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

	fmt.Println(antennas)
	fmt.Println(antiNodes)
	part1 = len(antiNodes)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 8 time:", time.Since(start))
}
