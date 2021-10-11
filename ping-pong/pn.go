package main

import "fmt"

func main() {
	fmt.Println("The program prints Ping, Pong or a number.")
	var n float64
	fmt.Scan(&n)
	fmt.Println("Enter a number:")
	if n % 2 {
		fmt.Println("Ping")
	} else if n % 5 {
		fmt.Println("Pong")
	} else if n%2 && n%5 {
		fmt.Println("Ping Pong")
	}
}
