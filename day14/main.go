package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
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

type Robot struct {
	pos      Pos
	velocity Pos
}

func strToInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return res
}

func moveRobot(r Robot, width int, height int) Pos {
	res := r.pos
	res.x += r.velocity.x
	res.y += r.velocity.y
	if res.x >= width {
		res.x -= width
	}
	if res.x < 0 {
		res.x += width
	}
	if res.y >= height {
		res.y -= height
	}
	if res.y < 0 {
		res.y += height
	}
	return res
}

func printLayout(robots []Robot, width int, height int, secondsElapsed int) {
	// Horiz line-ish frames: 24 127 = 103
	// Vert line-ish frames: 65 166 = 101

	// Only render images that have possible answer
	// Warning: i'm off by one on the file names since it's zero indexed (part2 will be one higher)
	if (secondsElapsed-24)%height == 0 || (secondsElapsed-64)%width == 0 {
		img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width * 2, height * 2}})
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				numBots := 0
				for r := 0; r < len(robots); r++ {
					if robots[r].pos.x == x && robots[r].pos.y == y {
						numBots++
					}
				}
				if numBots > 0 {
					img.Set(x*2, y*2, color.Black)
					img.Set((x*2)+1, y*2, color.Black)
					img.Set(x*2, (y*2)+1, color.Black)
					img.Set((x*2)+1, (y*2)+1, color.Black)
				} else {
					img.Set(x*2, y*2, color.White)
					img.Set((x*2)+1, y*2, color.White)
					img.Set(x*2, (y*2)+1, color.White)
					img.Set((x*2)+1, (y*2)+1, color.White)
				}
			}
		}

		f, _ := os.Create("output/" + strconv.Itoa(secondsElapsed) + ".png")
		png.Encode(f, img)
	}
}

func main() {
	fmt.Println("====== Day 14 ======")
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1 := 0
	part2 := 0

	width := 101
	height := 103

	robots := []Robot{}

	for scanner.Scan() {
		t := scanner.Text()

		re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

		matches := re.FindSubmatch([]byte(t))

		robots = append(robots, Robot{
			pos: Pos{
				x: strToInt(string(matches[1])),
				y: strToInt(string(matches[2])),
			},
			velocity: Pos{
				x: strToInt(string(matches[3])),
				y: strToInt(string(matches[4])),
			},
		})
	}

	secondsElapsed := 0

	for secondsElapsed < 100 {
		for r := 0; r < len(robots); r++ {
			robots[r].pos = moveRobot(robots[r], width, height)
		}
		printLayout(robots, width, height, secondsElapsed)
		secondsElapsed++
	}

	q1, q2, q3, q4 := 0, 0, 0, 0

	for r := 0; r < len(robots); r++ {
		if robots[r].pos.x < width/2 && robots[r].pos.y < height/2 {
			q1++
		}
		if robots[r].pos.x > width/2 && robots[r].pos.y < height/2 {
			q2++
		}
		if robots[r].pos.x < width/2 && robots[r].pos.y > height/2 {
			q3++
		}
		if robots[r].pos.x > width/2 && robots[r].pos.y > height/2 {
			q4++
		}
	}
	part1 = q1 * q2 * q3 * q4

	fmt.Println("Part 1:", part1)

	for secondsElapsed < width*height+1 {
		for r := 0; r < len(robots); r++ {
			robots[r].pos = moveRobot(robots[r], width, height)
		}

		printLayout(robots, width, height, secondsElapsed)
		secondsElapsed++
	}

	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 14 time:", time.Since(start))
}
