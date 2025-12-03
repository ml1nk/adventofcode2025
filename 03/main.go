package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

//go:embed data-example.txt
var example string

//go:embed data-puzzle.txt
var puzzle string

type leftandright struct {
	left  int
	right int
}

func main() {
	println("03")

	_, _ = fmt.Println("example1: ", solve01(example))
	_, _ = fmt.Println("example2: ", solve02(example))

	_, _ = fmt.Println("puzzle1: ", solve01(puzzle))
	_, _ = fmt.Println("puzzle2: ", solve02(puzzle))
}

func solve01(input string) int {
	entries := strings.Split(strings.Trim(input, "\n "), "\n")

	banks := lo.Map(entries, func(d string, _ int) []int {
		return lo.Map([]byte(d), func(d byte, _ int) int {
			r, err := strconv.Atoi(string(d))
			if err != nil {
				log.Fatal(err)
			}
			return r
		})
	})

	res := 0

	for _, d := range banks {

		index := 0

		first := lo.Reduce(d[0:len(d)-1], func(agg int, d int, in int) int {
			if d > agg {
				index = in
				return d
			} else {
				return agg
			}
		}, 0)

		second := lo.Reduce(d[index+1:], func(agg int, d int, in int) int {
			if d > agg {
				return d
			} else {
				return agg
			}
		}, 0)

		res += first*10 + second
	}

	return res
}

func solve02(input string) int {
	entries := strings.Split(strings.Trim(input, "\n "), "\n")

	banks := lo.Map(entries, func(d string, _ int) []int {
		return lo.Map([]byte(d), func(d byte, _ int) int {
			r, err := strconv.Atoi(string(d))
			if err != nil {
				log.Fatal(err)
			}
			return r
		})
	})

	res := 0

	for _, d := range banks {
		curIndex := 0
		subRes := 0

		for i := range 12 {
			myIndex := 0
			num := lo.Reduce(d[curIndex:len(d)-11+i], func(agg int, d int, in int) int {
				if d > agg {
					myIndex = in + 1
					return d
				} else {
					return agg
				}
			}, 0)
			curIndex += myIndex
			subRes = subRes*10 + num
		}

		res += subRes
	}

	return res
}
