package task_03

import "fmt"

func Task_03() {
	Multiplication(12, 1)
}

func Multiplication(a, b uint) {
	var n uint
	var i uint
	if a == 0 || b == 0 {
		n = 0
		fmt.Println(n)
	}
	if a > 0 && b > 0 {
		for i = 1; i <= b; i++ {
			n += a
		}
		fmt.Println(n)
	}
}
