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
	_, _ = fmt.Println("example2: ", solve02(example))

	_, _ = fmt.Println("puzzle1: ", solve01(puzzle))
	_, _ = fmt.Println("puzzle2: ", solve02(puzzle))
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

func solve02(input string) int {
	data := strings.Split(strings.Trim(input, "\n "), "\n\n")

	ranges := []*freshRange{}

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

		ranges = append(ranges, &freshRange{
			from: from,
			to:   to,
		})
	}

	for i, e := range ranges {
		for p, f := range ranges {
			if i == p || e == nil || f == nil {
				continue
			}

			merged := false

			if e.from <= f.from && e.to >= f.from {
				f.from = e.from
				merged = true
			}

			if e.to >= f.to && e.from <= f.to {
				f.to = e.to
				merged = true
			}

			if merged {
				ranges[i] = nil
			}
		}
	}

	fresh := 0

	for _, e := range ranges {
		if e == nil {
			continue
		}

		fresh += e.to - e.from + 1
	}

	return fresh
}
