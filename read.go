package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Task 1 - Your program will define a name struct which has two fields,
// fname for the first name, and lname for the last name.
type Person struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

// Open and read a file line by line, and for each line,
func openFile(fileName string) {

	fd, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	reader := bufio.NewReaderSize(fd, 1024)
	var persons []Person
	for {

		// Task 3 - Your program will read the file line by line and parse each line into a name struct.
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error: ", err)
		}

		lines := strings.Split(string(line), " ")
		if !isPrefix {
			// Task 4 - Your program will check if the first name and last name are less than 20 characters.
			p := Person{Fname: lines[0], Lname: lines[1]}
			if len(p.Fname) > 20 || len(p.Lname) > 20 {
				fmt.Println(fmt.Sprintf("Over size 20 characters: %+v", p))
				continue
			}

			// Task 5 - Your program will store the name structs in a slice.
			persons = append(persons, p)
		}
	}

	// Task 6 - Your program will marshal the slice of name structs into JSON format.
	p, err := json.Marshal(&persons)
	if err != nil {
		fmt.Println("Error: JSON marshal failed: ", err)
	}
	fmt.Printf(string(p))
}

func main() {
	var fileName string

	// Task 2 - Your program will read a file name from the command line.
	fmt.Print("Please enter input file name:")
	fmt.Scan(&fileName)
	log.Print(fileName)
	openFile(fileName)
}
