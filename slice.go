package main

import (
	"fmt"
	"sort"
	"strconv"
)

// Task 1 - The program should create an empty integer slice of size (length) 3.
var numList = make([]int, 3)

// Task 4 - The program should allow the user to add integers to the slice.
// If the slice is full, it should automatically resize to accommodate the new integer.
func add(i, n int) []int {
	if i <= 2 {
		numList[i] = n
	} else {
		numList = append(numList, n)
	}
	return numList
}

// If it does, it returns true; otherwise, it returns false.
func isContains(nl []int, n int) bool {
	for _, i := range nl {
		if i == n {
			return true
		}
	}
	return false
}

func main() {
	var i int
	var str string

	for {
		// Task 2 - During each pass through the loop,
		// the program prompts the user to enter an integer to be added to the slice
		fmt.Print("Please enter an integer: ")
		fmt.Scan(&str)

		// Task 5 - The program should check if the user entered "X" to exit the loop
		// and terminate the program.
		if str == "X" {
			break
		}

		num, _ := strconv.Atoi(str)
		if !isContains(numList, num) {
			nl := add(i, num)
			if i >= 2 {
				// Task 3 - The program should sort the slice in ascending order
				// and print the contents of the slice.
				sort.Ints(nl)
				fmt.Println(nl)
			}
			i++
		}
	}
}
