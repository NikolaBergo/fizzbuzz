package fizzbuzz_test

import (
	"fmt"
	"os"
	"strings"

	"fizzbuzz"
)

// ExampleCorrectInput'demostrates the extra spaces and newlines tolerance of FizzBuzz()
func ExampleCorrectInput () {
	correctInput := "5 4 1 0  -6   -5 \n   50 13456789     \n\n 3 \n\n"

	r := strings.NewReader(correctInput)
	err := fizzbuzz.FizzBuzz(r, os.Stdout)

	fmt.Printf(", ERROR = %+v", err)
	// Output: buzz 4 1 fizzbuzz fizz buzz buzz 13456789 fizz , ERROR = <nil>
}

// ExampleOverflowInput demonstrates detecting too large number on input
func ExampleOverflowInput () {
	overflowInput := "1 100 35 15 -666 12345678909887666565656554433455 1 23 17"

	r := strings.NewReader(overflowInput)
	err := fizzbuzz.FizzBuzz(r, os.Stdout)

	fmt.Printf(", ERROR = %+v", err)
	// Output: 1 buzz buzz fizzbuzz fizz , ERROR = strconv.ParseInt: parsing "12345678909887666565656554433455": value out of range
}

// ExampleNoiIntInput demonstrates detecting non-int value on iput
func ExampleNotIntInput () {
	notIntInput := "1 200 kitty 15"

	r := strings.NewReader(notIntInput)
	err := fizzbuzz.FizzBuzz(r, os.Stdout)

	fmt.Printf(", ERROR = %+v", err)
	// Output: 1 buzz , ERROR = expected integer
}
