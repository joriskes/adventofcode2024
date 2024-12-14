package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Tile struct {
	x    int
	y    int
	crop string
}

type Region struct {
	tiles []Tile
	crop  string
	area  int
	fence int
	price int
}

func getTile(x, y int, layout map[int]map[int]Tile) (Tile, error) {
	if x < 0 || y < 0 || y >= len(layout) || x >= len(layout[y]) {
		return Tile{}, errors.New("Tile out of bounds")
	}
	return layout[y][x], nil
}

func calcArea(region Region) int {
	return len(region.tiles)
}

func calcFence(region Region, layout map[int]map[int]Tile) int {
	res := 0
	for _, t := range region.tiles {
		c := layout[t.y][t.x].crop
		f, err := getTile(t.x, t.y-1, layout)
		if err != nil || f.crop != c {
			res++
		}
		f, err = getTile(t.x, t.y+1, layout)
		if err != nil || f.crop != c {
			res++
		}
		f, err = getTile(t.x-1, t.y, layout)
		if err != nil || f.crop != c {
			res++
		}
		f, err = getTile(t.x+1, t.y, layout)
		if err != nil || f.crop != c {
			res++
		}
	}
	return res
}

func calcPrice(region Region) int {
	return region.area * region.fence
}

func collectRegionTiles(t Tile, layout map[int]map[int]Tile, visited *[]string) []Tile {
	c := t.crop
	*visited = append(*visited, strconv.Itoa(t.y)+","+strconv.Itoa(t.x))
	res := []Tile{t}

	f, err := getTile(t.x, t.y-1, layout)
	if err == nil && f.crop == c && !slices.Contains(*visited, strconv.Itoa(f.y)+","+strconv.Itoa(f.x)) {
		res = append(res, collectRegionTiles(f, layout, visited)...)
	}
	f, err = getTile(t.x, t.y+1, layout)
	if err == nil && f.crop == c && !slices.Contains(*visited, strconv.Itoa(f.y)+","+strconv.Itoa(f.x)) {
		res = append(res, collectRegionTiles(f, layout, visited)...)
	}
	f, err = getTile(t.x-1, t.y, layout)
	if err == nil && f.crop == c && !slices.Contains(*visited, strconv.Itoa(f.y)+","+strconv.Itoa(f.x)) {
		res = append(res, collectRegionTiles(f, layout, visited)...)
	}
	f, err = getTile(t.x+1, t.y, layout)
	if err == nil && f.crop == c && !slices.Contains(*visited, strconv.Itoa(f.y)+","+strconv.Itoa(f.x)) {
		res = append(res, collectRegionTiles(f, layout, visited)...)
	}

	return res
}

func main() {
	fmt.Println("====== Day 12 ======")
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
	regions := []Region{}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		char := strings.Split(line, "")
		layout[y] = map[int]Tile{}
		for x := 0; x < len(char); x++ {
			layout[y][x] = Tile{
				x, y, char[x],
			}
		}
		y++
	}

	for y = 0; y < len(layout); y++ {
		for x := 0; x < len(layout[y]); x++ {
			found := false
			for r := 0; r < len(regions); r++ {
				if slices.Contains(regions[r].tiles, layout[y][x]) {
					found = true
				}
			}

			// New tile, start new region
			if !found {
				visited := []string{}
				regions = append(regions, Region{
					tiles: collectRegionTiles(layout[y][x], layout, &visited),
					crop:  layout[y][x].crop,
				})
			}
		}
	}

	for r := 0; r < len(regions); r++ {
		regions[r].area = calcArea(regions[r])
		regions[r].fence = calcFence(regions[r], layout)
		regions[r].price = calcPrice(regions[r])
		part1 += regions[r].price
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 12 time:", time.Since(start))
}
