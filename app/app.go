package main

import "fmt"

// Replace the comments below with your explanations. Try to provide a
// "high-level" idea, instead of writing literally what each line does.
//
// TODO: What does the program do?
//write the new num variable, which depends on formula
func main() {
	var num uint
	fmt.Scan(&num)

	// TODO: What do the variables a, b and c store?
	a, b, c := num%10, num/10%10, num/100%10
	// show and write which a, b, c variables need to be, after this changes. We write num. and after that program change our variables with formule, which we have

	// TODO: What do the three `if ... { ... }` statements below do?
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	// Show, how need to work programm if we have this statements, which are differents with mormal state.

	// TODO: What do the three lines below do?
	//
	num = a
	num = num*10 + b
	num = num*10 + c
	// score how to write answer (in println) by this formula
	fmt.Println(num)
}
