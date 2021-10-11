package main

import "fmt"

func main() {

	fmt.Println("The program finds the capacity of three condensers connected in series.")
	var x, y, z float64
	fmt.Print("Enter capacity of three condensers:")
	fmt.Scan(&x, &y, &z)
	invc := 1/x + 1/y + 1/z
	fmt.Println("The total capacity of the three condensers connected in series is", 1/invc)
}
