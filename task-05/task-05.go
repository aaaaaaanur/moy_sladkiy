package task_05

import "fmt"

func Task_05() {
	Division(5, 2)
}

func Division(a, d uint) {
	var q uint
	var r uint

	if a == 0 {
		q = 0
		r = 0
	}

	if a > 0 && a < d {
		q = 0
		r = a
	}

	if a > 0 && a == d {
		q = 1
		r = 0
	}

	if a > 0 && a > d {
		for {
			a = a - d
			q += 1

			if a < d {
				r = a
				break
			}
		}
	}

	fmt.Println("Частное: ", q)
	fmt.Println("Остаток: ", r)
}
