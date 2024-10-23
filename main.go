package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Read .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, copy .env.example to .env and update the env vars")
	}

	// Ask day number
	fmt.Println("Creates or updates a day setup for advent of code")
	fmt.Println("Enter day number: ")

	reader := bufio.NewReader(os.Stdin)
	day, _ := reader.ReadString('\n')
	day = strings.TrimSpace(day)

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

	// Write .go file (if non-existing)
	goFile := path + string(os.PathSeparator) + "main.go"
	if _, err := os.Stat(goFile); errors.Is(err, os.ErrNotExist) {
		if err := os.WriteFile(goFile, []byte("todo"), os.ModePerm); err != nil {
			fmt.Printf("Could create go file "+goFile, err)
			os.Exit(1)
		}
	}
}
