package main

import (
	"fmt"
	"os"

	"github.com/NikolaBergo/fizzbuzz/src"
)

// main represents common user interface
func main() {

	// creating input file in a current directory
	r, err := os.Create("inputdata")
	if err != nil {
		fmt.Print("Enable to create input file\n")
		return
	}
	r.Close()

	fmt.Printf("Welcome to fizzbuzz!\nPlease, write your data to 'inputdata' file and press ENTER\n")

	// scanf is waiting for user to allow program to continue
	var any byte
	fmt.Scanf("%c", &any)

	reader, _ = os.Open("inputdata")
	writer, _ := os.Create("outputdata")

	err = fizzbuzz.FizzBuzz(reader, writer)

	if err != nil {
		fmt.Printf("Ooop-s, something went wrong, check your input, error: %v \n ", err)
	} else {
		fmt.Print("Check out the results in output file!=)\n")
	}

	r.Close()
	w.Close()
}
