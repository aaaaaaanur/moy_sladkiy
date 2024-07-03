package task_02

import (
	"fmt"
	"math"
)

func Task_02() {
	Exponentiation(-9, 2)
}

func Exponentiation(a int, n uint) {
	b := math.Pow(float64(a), float64(n))
	fmt.Println(b)
}
