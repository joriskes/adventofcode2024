package main

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Advent of code 2024 usage:")
		fmt.Println("\ngo run main.go create <DAY_NUMBER>")
		fmt.Println("Generates a new day using supplied number, copies template.go there and downloads the day input (if available)")
		fmt.Println("\ngo run main.go run")
		fmt.Println("Runs all available days in sequence")
		fmt.Println("\ngo run main.go run <DAY_NUMBER>")
		fmt.Println("Runs the supplied day")
	}

	switch os.Args[1] {
	case "create":
		create(os.Args[2])
	case "run":
		if len(os.Args) > 2 && os.Args[2] != "" {
			run(os.Args[2])
		} else {
			for i := 1; i < 26; i++ {
				d := strconv.Itoa(i)
				if _, err := os.Stat("./day" + d); !os.IsNotExist(err) {
					run(d)
				}
			}

		}
	default:
		fmt.Println("Unknown command")
	}

}

func run(day string) {
	cmd := exec.Command("go", "run", "main.go")
	// Swap working directory before running
	cmd.Dir = "day" + day
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(string(output))
}

func create(day string) {
	// Read .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, copy .env.example to .env and update the env vars")
		os.Exit(1)
	}

	year := os.Getenv("AOC_YEAR")
	session := os.Getenv("AOC_SESSION")

	// Get input file
	url := "https://adventofcode.com/" + year + "/day/" + day + "/input"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", "session="+session)
	res, _ := client.Do(req)

	if res.StatusCode != 200 {
		fmt.Println("Failed grabbing input from " + url + ": " + res.Status + " is your AOC_SESSION valid and is " + year + " day " + day + " already unlocked?")
		os.Exit(1)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)

	// Make day directory if not exists
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Could not determine own path", err)
		os.Exit(1)
	}

	path := filepath.Join(pwd, "day"+day)
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("Could create path "+path, err)
		os.Exit(1)
	}

	// (Over)Write input
	inputFile := path + string(os.PathSeparator) + "input.txt"
	if err := os.WriteFile(inputFile, resBody, os.ModePerm); err != nil {
		fmt.Printf("Could create input file "+inputFile, err)
		os.Exit(1)
	}

	// Read template
	template, err := os.ReadFile("template.go")
	if err != nil {
		fmt.Print(err)
	}

	template = []byte(strings.Replace(string(template), "$DAY", day, -1))

	// Write template to new .go file (if non-existing)
	goFile := path + string(os.PathSeparator) + "main.go"
	if _, err := os.Stat(goFile); errors.Is(err, os.ErrNotExist) {
		if err := os.WriteFile(goFile, template, os.ModePerm); err != nil {
			fmt.Printf("Could create go file "+goFile, err)
			os.Exit(1)
		}
	}
}
