package main

import (
	"fmt"
)

func main() {
	var num int
	var numbers []int

	// Task 1 - The program should prompt the user to type in a sequence of up to 10 integers.
	for i := 0; i < 10; i++ {
		fmt.Print("Please enter a integer: ")
		fmt.Scan(&num)
		numbers = append(numbers, num)
	}

	// Task 2 - As part of this program, you should write a function called BubbleSort()
	// which takes a slice of integers as an argument and returns nothing
	BubbleSort(numbers)

	// Task 6 - The program should print the sorted integers in ascending order.
	for i, number := range numbers {
		if i == len(numbers) {
			fmt.Println(number)
		}
		fmt.Printf("%d ", number)
	}
}

// Task 3 - The BubbleSort() function should modify the slice so that the elements are in sorted order.
// The program should sort the integers in ascending order using the bubble sort algorithm.
func BubbleSort(num []int) {
	for i := len(num); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			for num[j] > num[j+1] {
				// Task 4 - The BubbleSort() function should use a helper function called Swap()
				// which swaps the position of two adjacent elements in the slice.
				Swap(num, j)
			}
		}
	}
}

// Task 5 - The Swap() function should take two arguments: the slice of integers and the index of the first element to swap.
// The Swap() function should not check for out-of-bounds errors.
func Swap(num []int, index int) {
	tmp := num[index+1]
	num[index+1] = num[index]
	num[index] = tmp
}
