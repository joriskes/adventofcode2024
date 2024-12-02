package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("====== Day $DAY ======")
	start := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		t := scanner.Text()
		// Todo: run day
		fmt.Println(t)

	}

	fmt.Println("⏱️ Day $DAY time:", time.Since(start))
}
