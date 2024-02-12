package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

/*
Our users are frequently trying to run custom analytics queries on their message deliverability metrics,
and end up writing a bad query that tries to divide a number by zero. It's become such a problem, 
that we think it would be best to make a specific type of error for division by zero.

Update the code so that the `divideError` type implements the `error` interface. 
Its `Error()` method should just return a string formatted in the following way:

```
can not divide DIVIDEND by zero
```

Where `DIVIDEND` is the actual dividend of the `divideError`. 
Use the `%v` [verb](https://pkg.go.dev/fmt#hdr-Printing) to format the dividend as a float.
*/

func (d divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", d.dividend)
}

// don't edit below this line

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		// We convert the `divideError` struct to an `error` type by returning it
		// as an error. As an error type, when it's printed its default value
		// will be the result of the Error() method
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func test(dividend, divisor float64) {
	defer fmt.Println("====================================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", dividend, divisor)
	quotient, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
}

func main() {
	test(10, 0)
	test(10, 2)
	test(15, 30)
	test(6, 3)
}
