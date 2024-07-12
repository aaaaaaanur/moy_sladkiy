package task_06

import "fmt"

func Task_6_2() {
	fmt.Println(Swap_2([]int{5, 6, 53, 1, 90, 39, 2}))
}

func Swap_2(x []int) []int {
	b := 0
	i := 1

	for {
		if len(x)-i < b {
			break
		}

		if len(x)-i == b {
			x[len(x)-i] = x[b]
		}

		if len(x)-i > b {
			x[len(x)-i] = x[len(x)-i] + x[b]
			x[b] = x[len(x)-i] - x[b]
			x[len(x)-i] = x[len(x)-i] - x[b]
		}

		b++
		i++
	}
	return x
}
