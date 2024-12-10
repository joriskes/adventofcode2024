package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	free bool
	id   int
	size int
}

type File struct {
	start int
	size  int
}

func printDisk(disk map[int]Block) {
	for i := 0; i < len(disk); i++ {
		if disk[i].free {
			fmt.Print(".")
		} else {
			fmt.Print(disk[i].id)
		}
	}
	fmt.Println()
}

func checkSum(disk map[int]Block) int {
	sum := 0
	for i := 0; i < len(disk); i++ {
		if !disk[i].free {
			sum += (i * disk[i].id)
		}
	}
	return sum
}

func main() {
	fmt.Println("====== Day 9 ======")
	start := time.Now()
	b, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	part1 := 0
	part2 := 0

	in := strings.Split(string(b), "")
	disk1 := map[int]Block{}
	disk2 := map[int]Block{}
	files := map[int]File{}

	p := 0
	id := 0
	isFree := false
	for i := 0; i < len(in); i++ {
		num, err := strconv.Atoi(in[i])
		if err != nil {
			fmt.Println(err)
			break
		}
		for j := 0; j < num; j++ {
			disk1[p] = Block{isFree, id, num - j}
			disk2[p] = Block{isFree, id, num - j}
			if !isFree && j == 0 {
				files[id] = File{
					p,
					num,
				}
			}
			p++
		}
		if !isFree {
			id++
		}
		isFree = !isFree
	}

	// Part 1
	last := len(disk1) - 1
	for i := 0; i < len(disk1); i++ {
		if disk1[i].free {
			for disk1[last].free {
				last--
			}
			// Done check
			if last <= i {
				break
			}
			//Swap last with current
			l := disk1[last]
			disk1[last] = disk1[i]
			disk1[i] = l
		}
	}
	part1 = checkSum(disk1)

	// Part 2
	startSearchFrom := 0
	for i := len(files) - 1; i > 0; i-- {
		file := files[i]
		move := -1
		foundFree := false
		for d := startSearchFrom; d < file.start; d++ {
			if disk2[d].free {
				if !foundFree {
					// Next time skip all non-free starting space
					startSearchFrom = d
					foundFree = true
				}
				if disk2[d].free && disk2[d].size >= file.size {
					move = d
					break
				}
			}
			d += disk2[d].size - 1 // Jump over complete space for speed
		}
		if move > -1 && move < file.start {
			for d := 0; d < file.size; d++ {
				b := disk2[move+d]
				disk2[move+d] = disk2[file.start+d]
				disk2[file.start+d] = b
			}
		}
	}
	part2 = checkSum(disk2)

	fmt.Println("Done")

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2) //6418529470362
	fmt.Println("⏱️ Day 9 time:", time.Since(start))
}
