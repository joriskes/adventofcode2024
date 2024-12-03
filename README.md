# Advent Of Code 2024

Advent of code challenges 2024, in GO. I'm not a GO developer, using this years aoc to
learn GO.

## Installation

Copy `.env.example` to `.env` and set the environment variable called `AOC_SESSION` with the session cookie of
adventofcode.com

## Running

To run all days run: `go run main.go run` from the root.
To run a spefic day add that to the run command: `go run main.go run 1`.

## Auto download / day creation script

To create / update a day run:
`go run main.go create <DAY_NUMBER>` from the root. It will create a new day as a directory, download the
AoC input and copy the template there
