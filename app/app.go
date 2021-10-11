package main

import "fmt"

// Replace the comments below with your explanations. Try to provide a
// "high-level" idea, instead of writing literally what each line does.
//
// This program if all numbers in variable num goes from biggest to smallest flips number to all numbers go from smallest to biggest.
func main() {
	var num uint
	fmt.Scan(&num)

	/* the variable number is divided into three numbers -
	the first digit of the number, the second digit of the number and the third*/
	a, b, c := num%10, num/10%10, num/100%10

	// Then with if statements program finds which is the smallest biggest etc.
	// the smallest number will be in a, biggest in c and not smallest and not biggest in b
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}

	// In this three lines program assumes number back in order to all number go from smallest to biggest
	num = a
	num = num*10 + b
	num = num*10 + c

	fmt.Println(num)
}
