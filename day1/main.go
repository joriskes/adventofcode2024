package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var leftList = [](string){}
	var rightList = [](string){}

	for scanner.Scan() {
		v := strings.Split(scanner.Text(), "   ")
		leftList = append(leftList, v[0])
		rightList = append(rightList, v[1])
	}

	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})
	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	total := 0.0

	for i := 0; i < len(leftList); i++ {
		l, err := strconv.ParseFloat(leftList[i], 64)
		if err != nil {
			panic(err)
		}
		r, err := strconv.ParseFloat(rightList[i], 64)
		if err != nil {
			panic(err)
		}
		total += math.Abs(l - r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1: ", int(total))
}
