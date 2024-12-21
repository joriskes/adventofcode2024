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

type Tile struct {
	pos    Pos
	isWall bool
}

type Node struct {
	pos    Pos
	weight int
}

func posToIndex(pos Pos) string {
	return strconv.Itoa(pos.x) + "," + strconv.Itoa(pos.y)
}

func indexToPos(index string) Pos {
	parts := strings.Split(index, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return Pos{x: x, y: y}
}
func searchNode(nodesToVisit *map[string]Node, current Pos, endIndex Pos, accWeight int, heading int) {
	if current.x < 0 || current.y < 0 {
		return
	}
	currentIndex := posToIndex(current)
	currentNode, ok := (*nodesToVisit)[currentIndex]

	if !ok {
		return
	}
	//fmt.Println("Visit", currentNode)

	if currentNode.weight > -1 && currentNode.weight < accWeight {
		return
	}

	// Update current node's weight
	(*nodesToVisit)[currentIndex] = Node{
		currentNode.pos,
		accWeight,
	}

	// Stop after reaching the end
	if current.x == endIndex.x && current.y == endIndex.y {
		return
	}

	if heading == 0 {
		searchNode(nodesToVisit, Pos{x: current.x, y: current.y - 1}, endIndex, accWeight+1, 0)
	} else {
		searchNode(nodesToVisit, Pos{x: current.x, y: current.y - 1}, endIndex, accWeight+1001, 0)
	}
	if heading == 1 {
		searchNode(nodesToVisit, Pos{x: current.x + 1, y: current.y}, endIndex, accWeight+1, 1)
	} else {
		searchNode(nodesToVisit, Pos{x: current.x + 1, y: current.y}, endIndex, accWeight+1001, 1)
	}
	if heading == 2 {
		searchNode(nodesToVisit, Pos{x: current.x, y: current.y + 1}, endIndex, accWeight+1, 2)
	} else {
		searchNode(nodesToVisit, Pos{x: current.x, y: current.y + 1}, endIndex, accWeight+1001, 2)
	}
	if heading == 3 {
		searchNode(nodesToVisit, Pos{x: current.x - 1, y: current.y}, endIndex, accWeight+1, 3)
	} else {
		searchNode(nodesToVisit, Pos{x: current.x - 1, y: current.y}, endIndex, accWeight+1001, 3)
	}

}

func main() {
	fmt.Println("====== Day 16 ======")
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
	startNodeIndex := ""
	endNodeIndex := ""
	layout := map[int]map[int]Tile{}
	nodesToVisit := map[string]Node{}
	for scanner.Scan() {
		line := scanner.Text()
		char := strings.Split(line, "")
		layout[y] = map[int]Tile{}
		for x := 0; x < len(char); x++ {
			newTile := Tile{
				pos: Pos{x, y},
			}
			posIndex := posToIndex(newTile.pos)
			switch char[x] {
			case "S":
				startNodeIndex = posIndex
				newTile.isWall = false
			case "E":
				endNodeIndex = posIndex
				newTile.isWall = false
			case ".":
				newTile.isWall = false
			case "#":
				newTile.isWall = true
			}

			layout[y][x] = newTile
			if !newTile.isWall {
				nodesToVisit[posIndex] = Node{
					pos:    newTile.pos,
					weight: -1,
				}
			}
		}
		y++
	}

	searchNode(&nodesToVisit, indexToPos(startNodeIndex), indexToPos(endNodeIndex), 0, 1)

	part1 = nodesToVisit[endNodeIndex].weight

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 16 time:", time.Since(start))
}
