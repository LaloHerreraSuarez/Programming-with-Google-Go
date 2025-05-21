package main

import (
	"fmt"
	"math"
)

func main() {
	var acceleration, velocity, displacement, time float64

	// Task 1 - Write a program which first prompts the user to enter
	// values for acceleration, initial velocity, and initial displacement
	fmt.Print("Please enter acceleration: ")
	fmt.Scan(&acceleration)

	fmt.Print("Please enter velocity: ")
	fmt.Scan(&velocity)

	fmt.Print("Please enter displacement: ")
	fmt.Scan(&displacement)

	fmt.Print("Please enter elapsed(time 1): ")
	fmt.Scan(&time)

	// Task 2 - You will need to define and use a function called GenDisplaceFn() which takes three float64
	fn := GenDisplaceFn(acceleration, velocity, displacement)
	fmt.Printf("Displacement after %f seconds\n", time)
	fmt.Println(fn(time))
	
	fmt.Print("Please enter elapsed(time 2): ")
	fmt.Scan(&time)

	// Task 2 - You will need to define and use a function called GenDisplaceFn() which takes three float64
	fn2 := GenDisplaceFn(acceleration, velocity, displacement)
	fmt.Printf("Displacement after %f seconds\n", time)
	fmt.Println(fn2(time))
}

// Task 3 - GenDisplaceFn() should return a function which computes displacement as a function of time
func GenDisplaceFn(acceleration, velocity, displacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return 0.5*acceleration*math.Pow(time, 2) + velocity*time + displacement
	}
}
