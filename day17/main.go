package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func strToInt(s string) uint64 {
	res, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return res
}

func main() {
	fmt.Println("====== Day 17 ======")
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1 := ""
	part2 := 0

	mem := map[string]uint64{}
	regs := []string{"A", "B", "C"}
	var prog []uint64
	var output []uint64

	instructionPointer := uint64(0)
	i := 0
	for scanner.Scan() {
		t := scanner.Text()

		if i < 3 {
			re := regexp.MustCompile(`Register [A|B|C]{1}: (\d+)`)
			matches := re.FindSubmatch([]byte(t))
			mem[regs[i]] = strToInt(string(matches[1]))
		} else if t != "" {
			s := strings.Replace(t, "Program: ", "", -1)
			a := strings.Split(s, ",")
			for _, v := range a {
				prog = append(prog, strToInt(v))
			}
		}

		i++
	}

	for instructionPointer < uint64(len(prog)) {
		opcode := prog[instructionPointer]
		literalOperand := prog[instructionPointer+1]

		comboOperand := literalOperand
		switch {
		case literalOperand == 4:
			comboOperand = mem["A"]
		case literalOperand == 5:
			comboOperand = mem["B"]
		case literalOperand == 6:
			comboOperand = mem["C"]
		}

		switch opcode {
		case 0: // adv
			mem["A"] = mem["A"] >> comboOperand // / uint64(math.Pow(2, float64(comboOperand)))
		case 1: // bxl
			mem["B"] = mem["B"] ^ literalOperand
		case 2: // bst
			mem["B"] = comboOperand % 8
		case 3: // jnz
			if mem["A"] != 0 {
				instructionPointer = literalOperand
				continue
			}
		case 4: // bxc
			mem["B"] = mem["B"] ^ mem["C"]
		case 5: // out
			output = append(output, comboOperand%8)
		case 6: // bdv
			mem["B"] = mem["A"] >> comboOperand // / uint64(math.Pow(2, float64(comboOperand)))
		case 7: // cdv
			mem["C"] = mem["A"] >> comboOperand // / uint64(math.Pow(2, float64(comboOperand)))

		}
		instructionPointer += 2
	}

	for _, num := range output {
		part1 += strconv.FormatUint(num, 10) + ","
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 17 time:", time.Since(start))
}
