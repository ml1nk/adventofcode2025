package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed data-example.txt
var example string

//go:embed data-puzzle.txt
var puzzle string

func main() {
	println("05")

	_, _ = fmt.Println("example1: ", solve01(example))

	_, _ = fmt.Println("puzzle1: ", solve01(puzzle))
}

type freshRange struct {
	from int
	to   int
}

func solve01(input string) int {
	data := strings.Split(strings.Trim(input, "\n "), "\n\n")

	ranges := []freshRange{}

	for e := range strings.SplitSeq(data[0], "\n") {
		a := strings.Split(e, "-")
		from, err := strconv.Atoi(string(a[0]))
		if err != nil {
			log.Fatal(err)
		}

		to, err := strconv.Atoi(string(a[1]))
		if err != nil {
			log.Fatal(err)
		}

		ranges = append(ranges, freshRange{
			from: from,
			to:   to,
		})
	}

	fresh := 0

	for e := range strings.SplitSeq(data[1], "\n") {
		val, err := strconv.Atoi(string(e))
		if err != nil {
			log.Fatal(err)
		}

		if isValid(ranges, val) {
			fresh++
		}

	}

	return fresh
}

func isValid(ranges []freshRange, val int) bool {
	for _, r := range ranges {
		if val >= r.from && val <= r.to {
			return true
		}
	}
	return false
}
