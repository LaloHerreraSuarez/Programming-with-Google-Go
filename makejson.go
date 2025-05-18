package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Create a JSON object from the map and print a Person JSON object.
func jsonPerson(name, address string) {
	person := map[string]string{
		"name":    name,
		"address": address,
	}

	// Task 3 - The program should should use Marshal() to create a JSON object from the map.
	barr, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error creating JSON object:", err)
	}

	// Task 4 - The program should print the JSON object to the console.
	// The JSON object should be in the format {"name": "name", "address": "address"}.
	fmt.Println(string(barr))
}

func main() {
	// Task 1 - The program should prompt the user to enter a name and address.
	fmt.Print("Please enter name: ")
	name := bufio.NewScanner(os.Stdin)
	name.Scan()

	fmt.Print("Please enter address: ")
	address := bufio.NewScanner(os.Stdin)
	address.Scan()

	// Task 2 - The program should create a JSON object with the name and address.
	// The JSON object should be in the format {"name": "name", "address": "address"}.
	jsonPerson(name.Text(), address.Text())
}
