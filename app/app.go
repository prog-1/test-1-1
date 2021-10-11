package main

import "fmt"

// Replace the comments below with your explanations. Try to provide a
// "high-level" idea, instead of writing literally what each line does.
//
// in this programm user enters number and with that number some math manipulations
func main() {
	var num uint
	fmt.Scan(&num)

	// a, b, and c stores number, that user enter, devided by some numbers
	a, b, c := num%10, num/10%10, num/100%10

	// if statements sort 2 numbers and change their values in ascending order
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}

	// get num other numbers values
	num = a
	num = num*10 + b
	num = num*10 + c

	fmt.Println(num)
}
