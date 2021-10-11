package main

import "fmt"

func main() {
	fmt.Println("The program prints Ping, Pong or a number.")
	fmt.Print("Enter a number: ")
	var num int
	fmt.Scan(&num)
	if num%2 == 0 && num%5 == 0 {
		fmt.Println("PingPong")
	} else if num%2 == 0 {
		fmt.Println("Ping")
	} else if num%5 == 0 {
		fmt.Println("Pong")
	} else {
		fmt.Println(num)
	}
}
