package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

//go:embed data-example.txt
var example string

//go:embed data-puzzle.txt
var puzzle string

func main() {
	println("06")

	_, _ = fmt.Println("example1: ", solve01(example))
	_, _ = fmt.Println("example2: ", solve02(example))

	_, _ = fmt.Println("puzzle1: ", solve01(puzzle))
	_, _ = fmt.Println("puzzle2: ", solve02(puzzle))
}

func solve01(input string) int {

	t := lo.Map(strings.Split(strings.Trim(input, "\n "), "\n"), func(d string, _ int) []string {
		res := strings.Trim(regexp.MustCompile(`\s+`).ReplaceAllString(d, " "), " ")
		return strings.Split(res, " ")
	})

	y := len(t)

	ops := t[y-1]

	start := lo.Map(t[0], func(s string, _ int) int {
		return toNum(s)
	})

	t = t[1 : y-1]

	for _, d := range t {
		for i, e := range d {
			switch ops[i] {
			case "+":
				start[i] += toNum(e)
			case "*":
				start[i] *= toNum(e)
			}
		}
	}

	return lo.Reduce(start, func(a int, b int, _ int) int {
		return a + b
	}, 0)
}

func solve02(input string) int {

	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	ops := strings.Split(lines[len(lines)-1], "")

	t := lo.Map(lines[:len(lines)-1], func(d string, _ int) []string {
		return strings.Split(d, "")
	})

	res := 0

	curOps := ""
	curRes := 0

	y := len(t)
	x := len(t[0])

	for xx := range x {
		curNum := ""

		for yy := range y {
			val := t[yy][xx]
			if val == " " {
				continue
			}

			curNum += val
		}

		if ops[xx] != " " {
			res += curRes
			curRes = toNum(curNum)
			curOps = ops[xx]
		} else {
			if len(curNum) == 0 {
				continue
			}

			switch curOps {
			case "+":
				curRes += toNum(curNum)
			case "*":
				curRes *= toNum(curNum)
			}
		}

		if xx == x-1 {
			res += curRes
		}

	}

	return res
}

func toNum(num string) int {
	n, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
