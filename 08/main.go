package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
)

//go:embed data-example.txt
var example string

//go:embed data-puzzle.txt
var puzzle string

func main() {
	println("08")

	_, _ = fmt.Println("example1: ", solve01(example, 10))
	_, _ = fmt.Println("example2: ", solve02(example))

	_, _ = fmt.Println("puzzle1: ", solve01(puzzle, 1000))
	_, _ = fmt.Println("puzzle2: ", solve02(puzzle))
}

type posd struct {
	a pos
	b pos
	d int
}

type pos struct {
	x int
	y int
	z int
}

func (p pos) distance(pp pos) int {
	dx := pp.x - p.x
	dx = dx * dx

	dy := pp.y - p.y
	dy = dy * dy

	dz := pp.z - p.z
	dz = dz * dz

	return dx + dy + dz
}

func (p pos) print() string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
}

func solve01(input string, iterations int) int {

	groups := map[pos]mapset.Set[pos]{}

	positions := lo.Map(strings.Split(strings.Trim(input, "\n "), "\n"), func(d string, _ int) pos {
		splitted := strings.Split(d, ",")
		return pos{
			x: toNum(splitted[0]),
			y: toNum(splitted[1]),
			z: toNum(splitted[2]),
		}
	})

	distances := []posd{}
	for i, a := range positions[0 : len(positions)-1] {
		for _, b := range positions[i+1:] {
			distances = append(distances, posd{
				a: a,
				b: b,
				d: a.distance(b),
			})
		}
	}

	slices.SortStableFunc(distances, func(a posd, b posd) int {
		return a.d - b.d
	})

	for _, a := range positions {
		groups[a] = mapset.NewSet(a)
	}

	for i := range iterations {

		p := distances[i]

		// already in the same group
		if groups[p.a] == groups[p.b] {
			continue
		}

		for t := range groups[p.b].Iter() {
			groups[p.a].Add(t)
			groups[t] = groups[p.a]
		}
	}

	res := lo.Map(
		lo.Uniq(
			lo.MapToSlice(groups, func(key pos, val mapset.Set[pos]) mapset.Set[pos] {
				return val
			}),
		),
		func(v mapset.Set[pos], _ int) int {
			return v.Cardinality()
		},
	)

	slices.Sort(res)

	return res[len(res)-3] * res[len(res)-2] * res[len(res)-1]
}

func solve02(input string) int {

	groups := map[pos]mapset.Set[pos]{}

	positions := lo.Map(strings.Split(strings.Trim(input, "\n "), "\n"), func(d string, _ int) pos {
		splitted := strings.Split(d, ",")
		return pos{
			x: toNum(splitted[0]),
			y: toNum(splitted[1]),
			z: toNum(splitted[2]),
		}
	})

	distances := []posd{}
	for i, a := range positions[0 : len(positions)-1] {
		for _, b := range positions[i+1:] {
			distances = append(distances, posd{
				a: a,
				b: b,
				d: a.distance(b),
			})
		}
	}

	slices.SortStableFunc(distances, func(a posd, b posd) int {
		return a.d - b.d
	})

	for _, a := range positions {
		groups[a] = mapset.NewSet(a)
	}

	for i := range len(distances) {

		p := distances[i]

		// already in the same group
		if groups[p.a] == groups[p.b] {
			continue
		}

		for t := range groups[p.b].Iter() {
			groups[p.a].Add(t)
			groups[t] = groups[p.a]
		}

		if groups[p.a].Cardinality() == len(positions) {
			return p.a.x * p.b.x
		}
	}

	return 0
}

func toNum(num string) int {
	n, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
