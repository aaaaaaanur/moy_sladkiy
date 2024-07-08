package task_04

import "fmt"

func Task_04() {
	Addition(4, 20)
}

func Addition(a, b uint) {
	c := a
	var i uint

	for i = 1; i <= b; i++ {
		c = c + 1
	}

	fmt.Println(c)
}
