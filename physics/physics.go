package main

import "fmt"

func main() {
	fmt.Println("The program finds the capacity of three condensers connected in series.")
	fmt.Print("Enter capacity of three condensers:")
	var c1, c2, c3 float64
	fmt.Scan(&c1, &c2, &c3)
	invc := 1/c1 + 1/c2 + 1/c3
	fmt.Println("The total capacity of the three condensers connected in series is", 1/invc)
}
