package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program prints Ping, Pong or a number.")
	var a int
	fmt.Println("Enter a number:")
	fmt.Scan(&a)
	if a%2 == 0 {
		if a%5 == 0 {
			fmt.Println("PingPong")
		} else {
			fmt.Println("Ping")
		}
	} else if a%5 == 0 {
		fmt.Println("Pong")
	} else {
		fmt.Println(a)
	}
}
