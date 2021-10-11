package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program finds the capacity of three condensers connected in series.")
	var C1, C2, C3 float64
	fmt.Println("Enter capacity of three condensers:")
	fmt.Scan(&C1, &C2, &C3)
	Cn := 1/C1 + 1/C2 + 1/C3
	C := 1 / Cn
	fmt.Println("The total capacity of the three condensers connected in series is", C)
}
