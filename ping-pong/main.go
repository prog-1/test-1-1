package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program prints Ping, Pong or a number.")
	var a float64
	fmt.Println("Enter a number:")
	fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Println("Ping")
	} else if a%5 == 0 {
		fmt.Println("Pong")
	} else if a%5 == 0 && a%2 == 0 {
		fmt.Println("PingPong")
	}
	fmt.Println(a)
}
