package main

import (
	_ "embed"
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
)

//go:embed data-example.txt
var example string

//go:embed data-puzzle.txt
var puzzle string

func main() {
	println("07")

	_, _ = fmt.Println("example1: ", solve01(example))
	_, _ = fmt.Println("example2: ", solve02(example))

	_, _ = fmt.Println("puzzle1: ", solve01(puzzle))
	_, _ = fmt.Println("puzzle2: ", solve02(puzzle))
}

type pos struct {
	x int
	y int
}

func solve01(input string) int {

	field := lo.Map(strings.Split(strings.Trim(input, "\n "), "\n"), func(d string, _ int) []string {
		return strings.Split(d, "")
	})

	_, start, _ := lo.FindIndexOf(field[0], func(item string) bool {
		return item == "S"
	})

	beams := []pos{pos{
		y: 0,
		x: start,
	}}

	splitter := mapset.NewSet[pos]()

	for len(beams) > 0 {
		beam := beams[0]
		beams = beams[1:]

		if beam.y >= len(field)-1 {
			continue
		}

		if field[beam.y+1][beam.x] == "^" {
			p := pos{y: beam.y + 1, x: beam.x}
			if splitter.Contains(p) {
				continue
			}
			splitter.Add(p)
			beams = append(beams, pos{y: beam.y + 1, x: beam.x - 1}, pos{y: beam.y + 1, x: beam.x + 1})
		} else {
			beams = append(beams, pos{y: beam.y + 1, x: beam.x})
		}

	}

	return splitter.Cardinality()
}

type post struct {
	x   int
	y   int
	ref pos
}

func solve02(input string) int {

	field := lo.Map(strings.Split(strings.Trim(input, "\n "), "\n"), func(d string, _ int) []string {
		return strings.Split(d, "")
	})

	_, start, _ := lo.FindIndexOf(field[0], func(item string) bool {
		return item == "S"
	})

	beams := []post{post{
		y: 0,
		x: start,
		ref: pos{
			y: 0,
			x: start,
		},
	}}

	splitter := map[pos]int{}

	splitter[pos{
		y: 0,
		x: start,
	}] = 1

	timelines := 0

	for len(beams) > 0 {
		beam := beams[0]
		beams = beams[1:]

		if beam.y >= len(field)-1 {
			timelines += splitter[beam.ref]
			continue
		}

		if field[beam.y+1][beam.x] == "^" {
			p := pos{y: beam.y + 1, x: beam.x}
			if _, ok := splitter[p]; ok {
				splitter[p] += splitter[beam.ref]
				continue
			}
			splitter[p] = splitter[beam.ref]
			beams = append(beams, post{y: beam.y + 1, x: beam.x - 1, ref: p}, post{y: beam.y + 1, x: beam.x + 1, ref: p})
		} else {
			beams = append(beams, post{y: beam.y + 1, x: beam.x, ref: beam.ref})
		}

	}

	return timelines
}
