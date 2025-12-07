package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/samber/lo"
)

//go:embed data-example.txt
var example string

//go:embed data-puzzle.txt
var puzzle string

func main() {
	println("04")

	_, _ = fmt.Println("example1: ", solve01(example))
	_, _ = fmt.Println("example2: ", solve02(example))

	_, _ = fmt.Println("puzzle1: ", solve01(puzzle))
	_, _ = fmt.Println("puzzle2: ", solve02(puzzle))
}

func solve01(input string) int {
	matrix := lo.Map(strings.Split(strings.Trim(input, "\n "), "\n"), func(d string, _ int) []bool {
		return lo.Map([]byte(d), func(d byte, _ int) bool {
			return d == '@'
		})
	})

	res := 0

	for i, e := range matrix {
		for p, _ := range e {
			if isValid(matrix, i, p) {
				res++
			}
		}
	}

	return res
}

func solve02(input string) int {
	matrix := lo.Map(strings.Split(strings.Trim(input, "\n "), "\n"), func(d string, _ int) []bool {
		return lo.Map([]byte(d), func(d byte, _ int) bool {
			return d == '@'
		})
	})

	res := 0

	oldres := -1
	for oldres != res {
		oldres = res

		for i, e := range matrix {
			for p, _ := range e {
				if isValid(matrix, i, p) {
					matrix[i][p] = false
					res++
				}
			}
		}
	}

	return res
}

func isValid(m [][]bool, i int, p int) bool {
	if !m[i][p] {
		return false
	}

	paper := 0

	di := -1
	for di <= 1 {
		dp := -1
		for dp <= 1 {

			ni := i + di
			np := p + dp

			if (di != 0 || dp != 0) && ni >= 0 && np >= 0 && ni < len(m) && np < len(m[0]) && m[ni][np] {
				paper++
			}

			dp++
		}
		di++
	}

	return paper < 4
}
