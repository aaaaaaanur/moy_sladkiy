package task_01

import "fmt"

func Task_1_2() {
	WithoutAddVar(5, 8)
}

func WithoutAddVar(a, b int) {
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a = %d, b = %d", a, b)

}
