package task_01

import "fmt"

func Task_1_1() {
	AddVar(5, 8)

}

func AddVar(a, b int) {
	var c int
	c = a
	a = b
	b = c

	fmt.Printf("a = %d, b = %d", a, b)
}
