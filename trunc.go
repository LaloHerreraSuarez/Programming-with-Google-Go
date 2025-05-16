package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64

	fmt.Print("Please enter a floating point number: ")
	fmt.Scan(&number)
	fmt.Printf("%.0f", math.Trunc(number))
}

// fmt.Printf("%.0f", math.Trunc(num)) is used to truncate the decimal part of the number.
// The math.Trunc function returns the integer part of the number, and %.0f formats it as a floating-point number with no decimal places.
// The result is printed to the console.
