package main

import "fmt"

func main() {
	fmt.Print("Suvrodeb Howlader\n")
	var Number int

	fmt.Print("Take a Number: ")
	fmt.Scan(&Number)

	if Number%2 == 0 {
		fmt.Print(Number, " is a Even Number")
	} else {
		fmt.Print(Number, " is a odd Number")
	}
}
