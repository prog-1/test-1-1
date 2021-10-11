package main

import "fmt"

// Replace the comments below with your explanations. Try to provide a
// "high-level" idea, instead of writing literally what each line does.
//
// TODO: What does the program do?
func main() {
	var num uint
	fmt.Scan(&num)

	// TODO: What do the variables a, b and c store?
	//проверяет делится ли число без остатка
	a, b, c := num%10, num/10%10, num/100%10

	// TODO: What do the three `if ... { ... }` statements below do?
	if a > b {
		a, b = b, a
		//swaps a and b
	}
	if a > c {
		a, c = c, a
		//swaps a and c
	}
	if b > c {
		b, c = c, b
		//swaps c and b
	}

	// TODO: What do the three lines below do?
	//These lines calculate (find) num
	num = a
	num = num*10 + b
	num = num*10 + c

	fmt.Println(num)
}
