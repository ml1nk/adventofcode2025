package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"reflect"
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
	println("02")

	_, _ = fmt.Println("example1: ", solve(example, false))
	_, _ = fmt.Println("example2: ", solve(example, true))

	_, _ = fmt.Println("puzzle1: ", solve(puzzle, false))
	_, _ = fmt.Println("puzzle2: ", solve(puzzle, true))
}

func solve(input string, second bool) int {
	entries := strings.Split(strings.Trim(input, "\n "), ",")
	data := lo.Map(entries, func(d string, _ int) leftandright {

		e := strings.Split(d, "-")

		left, err := strconv.Atoi(e[0])
		if err != nil {
			log.Fatal(err)
		}

		right, err := strconv.Atoi(e[1])
		if err != nil {
			log.Fatal(err)
		}

		return leftandright{
			left:  left,
			right: right,
		}
	})

	res := 0

	for _, d := range data {

		pos := d.left

		for pos <= d.right {

			if (!second && isInvalid1(pos)) || (second && isInvalid2(pos)) {
				res += pos
			}

			pos++
		}
	}

	return res
}

func isInvalid1(num int) bool {
	str := strconv.Itoa(num)
	return len(str)%2 == 0 && str[0:len(str)/2] == str[len(str)/2:]
}

func isInvalid2(num int) bool {
	str := strconv.Itoa(num)

	max := int(math.Floor(float64(len(str)) / 2.0))

	size := 1

	for size <= max {

		splitted := lo.Chunk([]byte(str), size)

		if len(splitted) < 2 {
			break
		}

		res := lo.EveryBy(splitted, func(e []byte) bool {
			return reflect.DeepEqual(e, splitted[0])
		})

		if res {
			return true
		}

		size++
	}

	return false
}
