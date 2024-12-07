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

type Tile struct {
	obstacle       bool
	visited        bool
	visitDirection int // To detect loops we need the directon of visit stored
	visitId        string
}

type Guard struct {
	direction int
	x         int
	startX    int
	y         int
	startY    int
	inBounds  bool
}

func printLayout(layout map[int]map[int]Tile) {
	for y := 0; y < len(layout); y++ {
		toPrint := ""
		for x := 0; x < len(layout[y]); x++ {
			c := "."
			if layout[y][x].obstacle {
				c = "#"
			}
			if layout[y][x].visited {
				c = "X"
			}
			toPrint += c
		}
		fmt.Println(toPrint)
	}
}

func runLayoutPart1(layout map[int]map[int]Tile, guard Guard) int {
	res := 1 // First location is also a step
	steps := 0
	for guard.inBounds {
		steps++
		// This is ugly
		layout[guard.y][guard.x] = Tile{
			obstacle: false,
			visited:  true, // Mark current location visited
		}
		nextX := guard.x
		nextY := guard.y
		switch guard.direction {
		case 0:
			nextY--
		case 1:
			nextX++
		case 2:
			nextY++
		case 3:
			nextX--
		}
		if nextY < 0 || nextY >= len(layout) || nextX < 0 || nextX >= len(layout[nextY]) {
			guard.inBounds = false
		} else {
			if layout[nextY][nextX].obstacle {
				guard.direction = (guard.direction + 1) % 4
			} else {
				if !layout[nextY][nextX].visited {
					res++
				}
				guard.x = nextX
				guard.y = nextY
			}
		}
	}
	return res
}

// We don't want to deep copy every bruteforce attempt so the id serves as a check if the visitdirection belongs to this attempt or a past one
func runLayoutPart2(layout map[int]map[int]Tile, guard Guard, id string) bool {
	steps := 0
	for guard.inBounds {
		steps++
		// Detect loop
		if layout[guard.y][guard.x].visited && layout[guard.y][guard.x].visitId == id && layout[guard.y][guard.x].visitDirection == guard.direction {
			return true
		}
		// This is ugly
		layout[guard.y][guard.x] = Tile{
			obstacle:       false,
			visited:        true, // Mark current location visited
			visitDirection: guard.direction,
			visitId:        id,
		}
		nextX := guard.x
		nextY := guard.y

		// Find next obstacle in direction the guard is going
		switch guard.direction {
		case 0:
			for !layout[nextY-1][nextX].obstacle && nextY >= 0 {
				steps++
				nextY--
			}
		case 1:
			for !layout[nextY][nextX+1].obstacle && nextX < len(layout[nextY]) {
				steps++
				nextX++
			}
		case 2:
			for !layout[nextY+1][nextX].obstacle && nextY < len(layout) {
				steps++
				nextY++
			}
		case 3:
			for !layout[nextY][nextX-1].obstacle && nextX >= 0 {
				steps++
				nextX--
			}
		}
		if nextY < 0 || nextY >= len(layout) || nextX < 0 || nextX >= len(layout[nextY]) {
			guard.inBounds = false
		} else {
			guard.direction = (guard.direction + 1) % 4
			guard.x = nextX
			guard.y = nextY
		}

		if steps > 10000 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("====== Day 6 ======")
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1 := 0
	part2 := 0

	layout := map[int]map[int]Tile{}
	y := 0
	guard := Guard{}
	for scanner.Scan() {
		line := scanner.Text()
		char := strings.Split(line, "")
		layout[y] = map[int]Tile{}
		for x := 0; x < len(char); x++ {
			layout[y][x] = Tile{
				obstacle: char[x] == "#",
				visited:  false,
			}
			if char[x] == "^" {
				guard.inBounds = true
				guard.direction = 0
				guard.x = x
				guard.y = y
				guard.startX = x
				guard.startY = y
			}
		}
		y++
	}

	part1 = runLayoutPart1(layout, guard)

	// Only place obstacles on visited spots
	todoList := []int{}
	for y := 0; y < len(layout); y++ {
		for x := 0; x < len(layout[y]); x++ {
			if layout[y][x].visited {
				todoList = append(todoList, y*len(layout)+x)
			}
		}
	}

	for y := 0; y < len(layout); y++ {
		for x := 0; x < len(layout[y]); x++ {
			if slices.Contains(todoList, y*len(layout)+x) && !layout[y][x].obstacle {
				// Place obstacle
				layout[y][x] = Tile{
					obstacle: true,
					visited:  false,
				}
				// Test run if loop
				if runLayoutPart2(layout, guard, strconv.Itoa(x)+"-"+strconv.Itoa(y)) {
					part2++
				}
				// Reset layout and guard
				layout[y][x] = Tile{
					obstacle: false,
					visited:  false,
				}
				guard.x = guard.startX
				guard.y = guard.startY
				guard.direction = 0
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 6 time:", time.Since(start))
}
