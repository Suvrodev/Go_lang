package main

import "fmt"

func main() {
	fmt.Print("Suvrodeb Howlader")
	var Name, Department string
	var Age float32

	fmt.Print("\nEnter Your Name: ")
	fmt.Scan(&Name)
	fmt.Print("Enter Your Age: ")
	fmt.Scan(&Age)
	print("Enter Your Department: ")
	fmt.Scan(&Department)

	fmt.Println("Your Name: ", Name)
	fmt.Println("Your Age: ", Age)
	fmt.Println("Your Department: ", Department)
}
