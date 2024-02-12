package main

import (
	"fmt"
)

/*
We offer a product that allows businesses that use Textio to send pairs of messages to couples. It is mostly used by flower shops and movie theaters.

Complete the `sendSMSToCouple` function. It should send 2 messages, first to the customer, then to the customer's spouse.

1. Use `sendSMS()` to send the `msgToCustomer`. If an error is encountered, return `0.0` and the error.
2. Do the same for the `msgToSpouse`
3. If both messages are sent successfully, return the total cost of the messages added together.

*When you return a non-nil error in Go, it's conventional to return the "zero" values of all other return values.*
*/

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	// send message to customer
	cost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}

	// send message to spuse
	cost2, err2 := sendSMS(msgToSpouse)
	if err2 != nil {
		return 0.0, err2
	}

	return cost + cost2, nil
}

// don't edit below this line

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}

func test(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("========")
	fmt.Println("Message for customer:", msgToCustomer)
	fmt.Println("Message for spouse:", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}

func main() {
	test(
		"Thanks for coming in to our flower shop today!",
		"We hope you enjoyed your gift.",
	)
	test(
		"Thanks for joining us!",
		"Have a good day.",
	)
	test(
		"Thank you.",
		"Enjoy!",
	)
	test(
		"We loved having you in!",
		"We hope the rest of your evening is absolutely fantastic.",
	)
}
