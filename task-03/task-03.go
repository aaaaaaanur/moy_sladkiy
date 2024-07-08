package task_03

import "fmt"

func Task_03() {
	Multiplication(30, 3)
}

func Multiplication(a, b uint) {
	var c uint
	var i uint

	if a == 0 || b == 0 {
		c = 0
		fmt.Println(c)
	}

	if a > 0 && b > 0 {
		for i = 1; i <= b; i++ {
			c += a
		}
		fmt.Println(c)
	}
}
