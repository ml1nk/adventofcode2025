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

//go:embed data-example2.txt
var example2 string

//go:embed data-puzzle.txt
var puzzle string

func main() {
	println("11")

	_, _ = fmt.Println("example1: ", solve01(example))
	_, _ = fmt.Println("puzzle1: ", solve01(puzzle))

	_, _ = fmt.Println("example2: ", solve02(example2))
	_, _ = fmt.Println("puzzle2: ", solve02(puzzle))
}

func solve01(input string) int {
	lines := lo.Map(strings.Split(strings.Trim(input, "\n "), "\n"), func(d string, _ int) []string {
		return strings.Split(d, " ")
	})

	connections := map[string][]string{}

	for _, line := range lines {
		key := line[0][:len(line[0])-1]
		connections[key] = line[1:]
	}

	res := 0
	queue := []string{"you"}

	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]

		for _, dest := range connections[el] {
			if dest == "out" {
				res++
				continue
			}

			queue = append(queue, dest)
		}

	}

	return res
}

type st struct {
	p       string
	visited mapset.Set[string]
}

type re struct {
	from    mapset.Set[string]
	to      mapset.Set[string]
	value   int
	reached int
}

func solve02(input string) int {
	lines := lo.Map(strings.Split(strings.Trim(input, "\n "), "\n"), func(d string, _ int) []string {
		return strings.Split(d, " ")
	})

	connections := map[string][]string{}

	for _, line := range lines {
		key := line[0][:len(line[0])-1]
		connections[key] = line[1:]
	}

	return ways(connections, "svr", "fft")*ways(connections, "fft", "dac")*ways(connections, "dac", "out") +
		ways(connections, "svr", "dac")*ways(connections, "dac", "fft")*ways(connections, "fft", "out")
}

func ways(connections map[string][]string, from string, to string) int {
	reached := map[string]*re{}

	reached[from] = &re{
		from:  mapset.NewThreadUnsafeSet[string](),
		to:    mapset.NewThreadUnsafeSet[string](),
		value: 1,
	}

	queue := []string{from}

	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]

		for _, dest := range connections[el] {

			d, ok := reached[dest]
			if ok {
				d.from.Add(el)
				continue
			} else {
				reached[dest] = &re{
					from: mapset.NewThreadUnsafeSet(el),
					to:   mapset.NewThreadUnsafeSet[string](),
				}
			}

			if dest == to {
				continue
			}

			queue = append(queue, dest)
		}
	}

	_, ok := reached[to]
	if !ok {
		return 0
	}

	queue = append(queue, to)
	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]

		d := reached[el]

		d.from.Each(func(s string) bool {
			d := reached[s]
			added := d.to.Add(el)
			if added {
				queue = append(queue, s)
			}
			return false
		})
	}

	queue = append(queue, from)
	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]

		d := reached[el]
		d.to.Each(func(s string) bool {
			e := reached[s]
			e.reached++
			e.value += d.value

			if e.reached == e.from.Cardinality() {
				queue = append(queue, s)
			}

			return false
		})

	}

	return reached[to].value
}
