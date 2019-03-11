package fizzbuzz

import (
	"fmt"
	"io"
)
// FizzBuzz() reads int64 from reader. If this number is devided by 3 it writes "fizz" to writer,
// if this number is devided by 5 it writes "buzz", if this number is devided both 3 and 5
// it writes "fizzbuzz" to writer
// all data are assumed to be within int64 size, numbers can be separated by any quantity of ' ' and '\n'
// if returned error != nil, it is not guaranteed that all data were
// read, written and processed correctly

func FizzBuzz(r io.Reader, w io.Writer) error {
	var number int64
	var err error

	// reads number until EOF or error
	for {
		_, err = fmt.Fscanf(r, "%d", &number)

		// case when newline or extra space in data allow program continue on reading
		for err != nil && err.Error() == moveOn {
            _, err = fmt.Fscanf(r, "%d", &number)
        }

		// case when EOF reached or another error occured
		if err != nil {
			break
		}

		_, err = fmt.Fprintf(w, "%s ", checkFizzBuzz(number))
		if err != nil {
			break
		}
	}

	// technically, EOF signals that everything has processed seccessfully
	// that's why it's not reported to user
	if err.Error() == "EOF" {
		return nil
	}

	return err
}

const (
	minus = '-'
	fizz  = "fizz"
	buzz  = "buzz"
	moveOn = "unexpected newline"
)

// checkFizzBuzz() checks wheather the number is devided by 5 or 3
func checkFizzBuzz(number int64) string {
	var index int
	var sumDigits int
	var digit int

	stringNumber := fmt.Sprintf("%d", number)
	result := stringNumber
	lastDigitIndex := len(stringNumber) - 1

	// passing the minus if number is < 0
	if stringNumber[0] == minus {
		index++
	}

	// calculate the sum of the number's digits
	for index = 0; index <= lastDigitIndex; index++ {
		fmt.Sscanf(stringNumber[index:index+1], "%d", &digit)
		sumDigits += digit
	}

	// checks wheather the sum of digits is devided by 3
	for sumDigits > 0 {
		sumDigits -= 3
	}
	if sumDigits == 0 {
		result = fizz
		// checks wheather the number is devided by 5
		if stringNumber[lastDigitIndex] == '5' || stringNumber[lastDigitIndex] == '0' {
			result += buzz
		}
		return result
	}

	// case when the number is not devided by 3 and is checked for division by 5
	if stringNumber[lastDigitIndex] == '5' || stringNumber[lastDigitIndex] == '0' {
		result = buzz
	}

	// returns the number in case when result value still had not been modifyed by this step
	return result
}
