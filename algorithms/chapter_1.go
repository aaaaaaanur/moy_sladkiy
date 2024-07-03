package algorithms

import "fmt"

func Chapter_1() {
	BinarySearch([]int{5, 6, 10, 40, 65, 104, 194, 318}, -13)
}

func BinarySearch(list []int, item int) {
	low := 0
	high := len(list) - 1

	for {
		if low > high {
			fmt.Println("None")
			break
		}

		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			fmt.Println(mid)
		} else if guess > item {
			high = mid - 1
		} else if guess < item {
			low = mid + 1
		}

		if guess == item {
			break
		}
	}
}
