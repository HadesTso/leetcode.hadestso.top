package main

import "fmt"

func main() {
	fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
}

func countStudents(students []int, sandwiches []int) int {
	n := len(students)
	circular, square := 0, 0
	for _, student := range students {
		if student == 0 {
			circular++
		} else {
			square++
		}
	}

	for _, sandwiche := range sandwiches {
		if sandwiche == 0 {
			if circular > 0 {
				n--
				circular--
			} else {
				break
			}
		} else {
			if square > 0 {
				n--
				square--
			} else {
				break
			}
		}
	}

	return n
}
