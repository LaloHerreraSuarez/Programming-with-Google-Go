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

	// Task 2 - You will need to define and use a function called GenDisplaceFn() which takes three float64
	// arguments, acceleration a, initial velocity vo, and initial displacement so
	fn := GenDisplaceFn(acceleration, velocity, displacement)

	fmt.Print("Please enter elapsed(time): ")
	fmt.Scan(&time)

	fmt.Printf("Displacement: %f after time: %f\n", fn(time), time)
}

// Task 3 - GenDisplaceFn() should return a function which computes displacement as a function of time,
// assuming the given values acceleration, initial velocity, and initial displacement.
// The function returned by GenDisplaceFn() should take one float64 argument t, representing time,
// and return one float64 argument which is the displacement travelled after time t.
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (a*math.Pow(t, 2))/2 + v0*t + s0
	}
	return fn
}
