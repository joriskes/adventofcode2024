package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1 := 0
	part2 := 0

	for scanner.Scan() {
		t := scanner.Text()
		// Todo: run day
		fmt.Println(t)

	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
