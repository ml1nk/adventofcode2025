package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

//go:embed data-example.txt
var example string

//go:embed data-puzzle.txt
var puzzle string

func main() {
	println("09")

	_, _ = fmt.Println("example1: ", solve01(example))
	_, _ = fmt.Println("example2: ", solve02(example, true))
	_, _ = fmt.Println("puzzle1: ", solve01(puzzle))
	_, _ = fmt.Println("puzzle2: ", solve02(puzzle, false)) // tested with false and true, autodetection is trivial
}

type pos struct {
	x int
	y int
}

type posx struct {
	a         pos
	b         pos
	rectangle int
}

func (p pos) rectangle(pp pos) int {
	dx := max(pp.x-p.x, p.x-pp.x) + 1
	dy := max(pp.y-p.y, p.y-pp.y) + 1
	return dx * dy
}

func solve01(input string) int {

	positions := lo.Map(strings.Split(strings.Trim(input, "\n "), "\n"), func(d string, _ int) pos {
		splitted := strings.Split(d, ",")
		return pos{
			x: toNum(splitted[0]),
			y: toNum(splitted[1]),
		}
	})

	largest := 0

	for i, a := range positions[0 : len(positions)-1] {
		for _, b := range positions[i+1:] {
			cur := a.rectangle(b)
			if largest < cur {
				largest = cur
			}
		}
	}

	return largest
}

func solve02(input string, positive bool) int {

	positions := lo.Map(strings.Split(strings.Trim(input, "\n "), "\n"), func(d string, _ int) pos {
		splitted := strings.Split(d, ",")
		return pos{
			x: toNum(splitted[0]),
			y: toNum(splitted[1]),
		}
	})

	ranges := getRanges(positions, positive)
	data := getSortedResults(positions)

fail:
	for _, d := range data {
		minX := min(d.a.x, d.b.x)
		maxX := max(d.a.x, d.b.x)

		minY := min(d.a.y, d.b.y)
		maxY := max(d.a.y, d.b.y)

	match:
		for q := range maxX - minX + 1 {
			curRange := ranges[minX+q]

			for i, q := range curRange {
				// no later range will match
				if minY < q && i%2 == 0 {
					continue fail
				}

				if minY >= q && i%2 == 0 && minY <= curRange[i+1] {
					if maxY <= curRange[i+1] {
						continue match
					} else {
						continue fail
					}
				}
			}
			continue fail
		}

		return d.rectangle
	}

	return 0
}

func xdiffvalid(valid map[int][]int, a pos, b pos, ignoreFirst bool, ignoreLast bool) {

	aLower := a.x < b.x

	c := a.x - b.x + 1
	if aLower {
		c = b.x - a.x + 1
	}

	mult := 1
	if !aLower {
		mult = -1
	}

	for q := range c {
		if q == 0 && ignoreFirst {
			continue
		}

		if q == c-1 && ignoreLast {
			continue
		}

		pos := a.x + q*mult

		if _, ok := valid[pos]; !ok {
			valid[pos] = []int{a.y}
		} else {
			valid[pos] = append(valid[pos], a.y)
		}
	}
}

func toNum(num string) int {
	n, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func getRelative(positions []pos, i int) pos {
	if i >= len(positions) {
		return positions[i%len(positions)]
	}
	if i < 0 {
		return positions[i+len(positions)]
	}
	return positions[i]
}

func getRanges(positions []pos, positive bool) map[int][]int {

	ranges := map[int][]int{}

	for i, cur := range positions {

		llcur := getRelative(positions, i-2)
		lcur := getRelative(positions, i-1)
		ncur := getRelative(positions, i+1)

		if lcur.y == cur.y {

			xdiffvalid(ranges, lcur, cur,
				(llcur.y < lcur.y && positive) || (llcur.y > lcur.y && !positive),
				(cur.y > ncur.y && positive) || (cur.y < ncur.y && !positive),
			)

			if lcur.x > cur.x {
				if cur.y > ncur.y {
					positive = !positive
				}
			} else {
				if cur.y < ncur.y {
					positive = !positive
				}
			}
		} else {
			if lcur.y > cur.y {
				if cur.x > ncur.x {
					positive = !positive
				}
			} else {
				if cur.x < ncur.x {
					positive = !positive
				}
			}
		}
	}

	for y, p := range ranges {
		slices.Sort(p)
		cleaned := []int{}

		i := 0
		open := false

		for i < len(p) {
			if !open {
				cleaned = append(cleaned, p[i])
				open = true
				i++
				continue
			}

			if i+1 < len(p) && p[i]+1 >= p[i+1] {
				i++
				continue
			}

			cleaned = append(cleaned, p[i])
			open = false
			i++
		}

		ranges[y] = cleaned
	}

	return ranges
}

func getSortedResults(positions []pos) []posx {
	data := []posx{}

	for i, a := range positions[0 : len(positions)-1] {
		for _, b := range positions[i+1:] {
			cur := a.rectangle(b)
			data = append(data, posx{
				a:         a,
				b:         b,
				rectangle: cur,
			})
		}
	}

	slices.SortStableFunc(data, func(a posx, b posx) int {
		return b.rectangle - a.rectangle
	})

	return data
}
