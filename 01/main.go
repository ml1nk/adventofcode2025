package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed data-example.txt
var example string

//go:embed data-puzzle.txt
var puzzle string

func main() {
	println("01")

	lines := strings.Split(example, "\n")

	_, _ = fmt.Println("example1: ", solve01(50, lines))
	_, _ = fmt.Println("example2: ", solve02(50, lines))

	lines = strings.Split(puzzle, "\n")

	_, _ = fmt.Println("puzzle1: ", solve01(50, lines))
	_, _ = fmt.Println("puzzle2: ", solve02(50, lines))
}

func solve01(pos int, lines []string) int {
	pw := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		i, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		if line[0] == 'R' {
			pos = (pos + i) % 100
		} else {
			pos = (pos - i) % 100
			if pos < 0 {
				pos = 100 + pos
			}
		}

		if pos == 0 {
			pw++
		}
	}

	return pw
}

func solve02(pos int, lines []string) int {
	pw := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		i, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		pw += int(math.Floor(float64(i) / 100.0))

		i = i % 100

		oldpos := pos

		if line[0] == 'R' {
			pos = (pos + i) % 100

			if pos > 0 && pos < oldpos && oldpos > 0 {
				pw++
			}
		} else {
			pos = (pos - i) % 100
			if pos < 0 {
				pos = 100 + pos
				if oldpos > 0 {
					pw++
				}
			}
		}

		if pos == 0 {
			pw++
		}
	}

	return pw
}
