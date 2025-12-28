package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
)

//go:embed data-example.txt
var example string

//go:embed data-puzzle.txt
var puzzle string

type field struct {
	x   int
	y   int
	els [6]int
}

func main() {
	println("12")

	//_, _ = fmt.Println("example1: ", solve01(example))
	_, _ = fmt.Println("puzzle1: ", solve01(puzzle))
}

type element struct {
	variations [][3][3]bool
	filled     int
}

func (e element) print() {
	for _, v := range e.variations {
		for _, l := range v {
			print(strings.Join(lo.Map(l[:], func(l bool, _ int) string {
				if l {
					return "#"
				} else {
					return "."
				}
			}), ""))
			println()
		}
		println()
	}
	println()
}

func solve01(input string) int {
	lines := lo.Map(strings.Split(strings.Trim(input, "\n "), "\n"), func(d string, _ int) []string {
		return strings.Split(d, " ")
	})

	fieldsRaw := lines[30:]
	fields := []field{}

	for _, f := range fieldsRaw {
		d := strings.Split(f[0][:len(f[0])-1], "x")

		els := [6]int{}

		for i := range els {
			els[i] = toNum(f[i+1])
		}

		fields = append(fields, field{
			x:   toNum(d[0]),
			y:   toNum(d[1]),
			els: els,
		})
	}

	elements := []element{}

	base := 1
	offset := 5
	for i := range 6 {
		start := offset*i + base
		count := 0
		el := [3][3]bool{}
		for p := range 3 {
			if lines[start+p][0][0] == '#' {
				el[p][0] = true
				count++
			}
			if lines[start+p][0][1] == '#' {
				el[p][1] = true
				count++
			}
			if lines[start+p][0][2] == '#' {
				el[p][2] = true
				count++
			}
		}

		ele := element{
			variations: vars(el),
			filled:     count,
		}

		ele.print()

		elements = append(elements, ele)

	}

	res := 0
	for _, f := range fields {
		if solveField(f, elements) {
			res++
		}
	}

	return res
}

func solveField(f field, els []element) bool {

	minElements := ((f.x - f.x%3) * (f.y - f.y%3)) / 9
	elements := lo.Sum(f.els[:])

	usedField := lo.Reduce(f.els[:], func(a int, m int, i int) int {
		return a + els[i].filled*m
	}, 0)

	if f.x*f.y < usedField {
		println("trivialmax", f.x*f.y, usedField)
		return false
	}

	if elements <= minElements {
		println("trivialmin", minElements, elements)
		return true
	}

	panic("nontrivial")
}

func vars(orig [3][3]bool) [][3][3]bool {
	variations := mapset.NewThreadUnsafeSet(orig)
	flip(variations, orig)
	rotate(variations, orig)

	return variations.ToSlice()
}

func rotate(variations mapset.Set[[3][3]bool], orig [3][3]bool) {
	next := [3][3]bool{}
	for i := range 3 {
		for p := range 3 {
			next[p][i] = orig[i][p]
		}
	}
	variations.Add(next)
	flip(variations, next)
}

func flip(variations mapset.Set[[3][3]bool], orig [3][3]bool) {
	next := [3][3]bool{}
	for i := range 3 {
		for p := range 3 {
			pp := (p-1)*-1 + 1
			next[i][pp] = orig[i][p]
		}
	}
	variations.Add(next)

	next = [3][3]bool{}
	for i := range 3 {
		for p := range 3 {
			pp := (p-1)*-1 + 1
			next[pp][i] = orig[p][i]
		}
	}
	variations.Add(next)
}

func toNum(num string) int {
	n, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
