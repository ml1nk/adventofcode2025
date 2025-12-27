package lgs

import (
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
	"github.com/samber/lo"
)

func Solve02(input string) int {
	machines := lo.Map(strings.Split(strings.Trim(input, "\n "), "\n"), func(d string, _ int) []string {
		return strings.Split(d, " ")
	})

	arr := make([]int, len(machines))
	for i, t := range machines {

		joltage := lo.Map(strings.Split(t[len(t)-1][1:len(t[len(t)-1])-1], ","), func(d string, _ int) int {
			return toNum(d)
		})

		buttons := lo.Map(t[1:len(t)-1], func(d string, _ int) []int {
			return lo.Map(strings.Split(d[1:len(d)-1], ","), func(d string, _ int) int {
				return toNum(d)
			})
		})
		arr[i] = solveMachine(i, joltage, buttons)
		// println("")
		// print(i)
		// print(" r ")
		// println(arr[i])
	}

	return lo.Sum(arr)
}

func solveMachine(i int, joltage []int, buttons [][]int) int {

	lp := golp.NewLP(len(joltage), len(buttons))

	for i := range buttons {
		lp.SetInt(i, true)
	}

	for i, p := range joltage {
		res := lo.Map(buttons, func(a []int, _ int) float64 {
			if slices.Contains(a, i) {
				return 1.0
			}
			return 0.0
		})

		lp.AddConstraint(res, golp.EQ, float64(p))
	}

	lp.SetObjFn(lo.Map(buttons, func(p []int, _ int) float64 {
		return 1.0
	}))

	ty := lp.Solve()
	if ty > 0 {
		panic(ty)
	}

	vars := lp.Variables()
	res := lo.Sum(vars)

	return int(res)
}

func toNum(num string) int {
	n, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
