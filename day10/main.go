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

type Pos struct {
	x int
	y int
}

type MapNode struct {
	pos    Pos
	height int
	score  int
}

func walkTrail(pos Pos, heightMap map[int]map[int]MapNode, visited *[]string, part2 bool) int {
	h := heightMap[pos.y][pos.x].height
	// Don't visit other trailheads
	if len(*visited) > 0 && h == 0 {
		return 0
	}
	p := strconv.Itoa(pos.x) + "," + strconv.Itoa(pos.y)
	// Only visit every location once
	if slices.Contains(*visited, p) {
		return 0
	}
	// For part 2 simply not keeping visited is enough
	if !part2 {
		*visited = append(*visited, p)
	}
	s := 0
	if h == 9 {
		s = 1
	}

	if pos.y > 0 && heightMap[pos.y-1][pos.x].height-h == 1 {
		s += walkTrail(Pos{pos.x, pos.y - 1}, heightMap, visited, part2)
	}
	if pos.y < len(heightMap)-1 && heightMap[pos.y+1][pos.x].height-h == 1 {
		s += walkTrail(Pos{pos.x, pos.y + 1}, heightMap, visited, part2)
	}
	if pos.x > 0 && heightMap[pos.y][pos.x-1].height-h == 1 {
		s += walkTrail(Pos{pos.x - 1, pos.y}, heightMap, visited, part2)
	}
	if pos.x < len(heightMap[pos.y])-1 && heightMap[pos.y][pos.x+1].height-h == 1 {
		s += walkTrail(Pos{pos.x + 1, pos.y}, heightMap, visited, part2)
	}
	return s
}

func main() {
	fmt.Println("====== Day 10 ======")
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1 := 0
	part2 := 0

	heightMap := map[int]map[int]MapNode{}
	trailHeads := []Pos{}

	y := 0
	for scanner.Scan() {
		heightMap[y] = map[int]MapNode{}
		line := strings.Split(scanner.Text(), "")
		for x, l := range line {
			h, _ := strconv.Atoi(l)
			heightMap[y][x] = MapNode{
				pos:    Pos{x, y},
				height: h,
				score:  0,
			}

			if h == 0 {
				trailHeads = append(trailHeads, Pos{x, y})
			}
		}
		y++
	}

	//printMap(heightMap)
	for _, trailHead := range trailHeads {
		v := []string{}
		part1 += walkTrail(trailHead, heightMap, &v, false)
		v = []string{}
		part2 += walkTrail(trailHead, heightMap, &v, true)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 10 time:", time.Since(start))
}
