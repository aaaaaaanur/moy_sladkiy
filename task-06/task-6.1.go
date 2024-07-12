package task_06

import "fmt"

func Task_6_1() {
	Swap([]int{5, 6, 53, 1, 90, 39})
}

func Swap(x []int) {
	var y []int

	for i := 1; ; i++ {
		if len(x)-i < 0 {
			break
		}

		y = append(y, x[len(x)-i])
	}
	fmt.Println(y)
}
