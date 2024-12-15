package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Pos struct {
	x, y int
}

type Tile struct {
	pos    Pos
	isWall bool
}

type Box struct {
	pos Pos
}

type Robot struct {
	pos       Pos
	moves     []string
	moveIndex int
}

func printLayout(layout map[int]map[int]Tile, boxes []Box, robot Robot) {
	for y := 0; y < len(layout); y++ {
		for x := 0; x < len(layout[y]); x++ {
			char := "."
			if layout[y][x].isWall {
				char = "#"
			}
			for b := 0; b < len(boxes); b++ {
				if boxes[b].pos.x == x && boxes[b].pos.y == y {
					if char != "." {
						fmt.Println()
						fmt.Print("F: Box on " + char)
						fmt.Println()
					}
					char = "O"
				}
			}
			if robot.pos.x == x && robot.pos.y == y {
				if char != "." {
					fmt.Println()
					fmt.Println("F: Robot on " + char)
					fmt.Println()
				}
				char = "@"
			}
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func moveBox(layout map[int]map[int]Tile, boxes *[]Box, boxIndex int, xoffset int, yoffset int) bool {
	newX := (*boxes)[boxIndex].pos.x + xoffset
	newY := (*boxes)[boxIndex].pos.y + yoffset
	movePossible := false
	for b := 0; b < len(*boxes); b++ {
		// Box found in new position, move it as well (if possible)
		if b != boxIndex && (*boxes)[b].pos.x == newX && (*boxes)[b].pos.y == newY {
			movePossible = moveBox(layout, boxes, b, xoffset, yoffset)
			if !movePossible {
				return false
			}
		}
	}
	if !movePossible {
		movePossible = !layout[newY][newX].isWall
	}
	if movePossible {
		(*boxes)[boxIndex].pos.x = newX
		(*boxes)[boxIndex].pos.y = newY
	}
	return movePossible
}

func moveRobot(layout map[int]map[int]Tile, boxes *[]Box, robot *Robot, xoffset int, yoffset int) {
	newX := robot.pos.x + xoffset
	newY := robot.pos.y + yoffset
	movePossible := false
	for b := 0; b < len(*boxes); b++ {
		// Box found in new position, move it as well (if possible)
		if (*boxes)[b].pos.x == newX && (*boxes)[b].pos.y == newY {
			movePossible = moveBox(layout, boxes, b, xoffset, yoffset)
			if !movePossible {
				return
			}
		}
	}
	if !movePossible {
		movePossible = !layout[newY][newX].isWall
	}
	if movePossible {
		robot.pos.x = newX
		robot.pos.y = newY
	}

}

func calcGPS(pos Pos) int {
	return (pos.y * 100) + pos.x
}

func main() {
	fmt.Println("====== Day 15 ======")
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
	boxes := []Box{}

	robot := Robot{
		Pos{0, 0},
		[]string{},
		0,
	}

	y := 0

	readingMap := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingMap = false
		} else {
			if readingMap {
				char := strings.Split(line, "")
				layout[y] = map[int]Tile{}
				for x := 0; x < len(char); x++ {
					switch char[x] {
					case "#":
						layout[y][x] = Tile{
							Pos{x, y}, true,
						}
					case ".":
						layout[y][x] = Tile{
							Pos{x, y}, false,
						}
					case "O":
						boxes = append(boxes, Box{Pos{x, y}})
						layout[y][x] = Tile{
							Pos{x, y}, false,
						}
					case "@":
						robot.pos.x = x
						robot.pos.y = y
						layout[y][x] = Tile{
							Pos{x, y}, false,
						}
					}
				}
			} else {
				char := strings.Split(line, "")
				for m := 0; m < len(char); m++ {
					robot.moves = append(robot.moves, char[m])
				}
			}
		}
		y++
	}

	for robot.moveIndex < len(robot.moves) {
		move := robot.moves[robot.moveIndex]
		switch move {
		case "^":
			moveRobot(layout, &boxes, &robot, 0, -1)
		case "v":
			moveRobot(layout, &boxes, &robot, 0, 1)
		case "<":
			moveRobot(layout, &boxes, &robot, -1, 0)
		case ">":
			moveRobot(layout, &boxes, &robot, 1, 0)
		}

		robot.moveIndex++
	}

	for b := 0; b < len(boxes); b++ {
		part1 += calcGPS(boxes[b].pos)
	}

	//printLayout(layout, boxes, robot)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 15 time:", time.Since(start))
}
