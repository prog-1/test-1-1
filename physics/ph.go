package main

import "fmt"

func main(){

	fmt.Println("The program finds the capacity of three condensers connected in series.")
	var C1, C2, C3 float64
	fmt.Println("Enter capacity of three condensers:")
	fmt.Scan(&C1, &C2, &C3)
	var C float64 
	1/C := 1/C1 + 1/C2 + 1/C3
	fmt.Println("The total capacity of the three condensers connected in series is", 1/C)
}