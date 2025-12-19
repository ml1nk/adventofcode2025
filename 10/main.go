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
	println("10")

	_, _ = fmt.Println("example1: ", solve01(example))
	//	_, _ = fmt.Println("example2: ", solve02(example, true))
	_, _ = fmt.Println("puzzle1: ", solve01(puzzle))
	//_, _ = fmt.Println("puzzle2: ", solve02(puzzle, false)) // tested with false and true, autodetection is trivial
}

func solve01(input string) int {
	machines := lo.Map(strings.Split(strings.Trim(input, "\n "), "\n"), func(d string, _ int) []string {
		return strings.Split(d, " ")
	})

	return lo.Reduce(machines, func(p int, t []string, _ int) int {
		lights := lo.Map(strings.Split(t[0][1:len(t[0])-1], ""), func(d string, _ int) bool {
			return d == "#"
		})

		buttons := lo.Map(t[1:len(t)-1], func(d string, _ int) []int {
			return lo.Map(strings.Split(d[1:len(d)-1], ","), func(d string, _ int) int {
				return toNum(d)
			})
		})

		return p + solveMachine(lights, buttons)
	}, 0)
}

func solveMachine(lights []bool, buttons [][]int) int {
	state := make([]bool, len(lights))
	if slices.Equal(state, lights) {
		return 0
	}

	reached := mapset.NewSet[string]()
	reached.Add(toString(state))

	depth := 0

	states := [][]bool{state}
	nextStates := [][]bool{}

	for len(states) > 0 {
		depth++

		for _, state := range states {
			for _, b := range buttons {
				newState := applyButton(state, b)

				id := toString(newState)
				if reached.ContainsOne(id) {
					continue
				}
				reached.Add(id)
				if slices.Equal(newState, lights) {
					return depth
				}
				nextStates = append(nextStates, newState)
			}
		}
		states = nextStates
	}

	return depth
}

func toString(in []bool) string {
	return lo.Reduce(in, func(a string, b bool, _ int) string {
		if b {
			return a + "1"
		}
		return a + "0"
	}, "")
}

func applyButton(state []bool, buttons []int) []bool {
	state = slices.Clone(state)
	for _, button := range buttons {
		state[button] = !state[button]
	}
	return state
}

func toNum(num string) int {
	n, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
