package main

import (
	"fmt"
	"math"
)

func Calculation(Number1, Number2 float32, Operator string) float32 {
	var GetValue float32
	switch Operator {
	case "+":
		GetValue = Number1 + Number2
	case "-":
		GetValue = Number1 - Number2
	case "*":
		GetValue = Number1 * Number2
	case "/":
		GetValue = Number1 / Number2
	default:
		break
	}

	return GetValue
}

func Power_Function(Base, Power int) {
	var Temp = 1
	for i := 1; i <= Power; i++ {
		Temp = Temp * int(Base)
	}

	fmt.Println("Power: ", Temp)

}

func main() {
	var Num1, Num2, Taken float32
	var Operator string
	var Base, Power int
	var Sqrt_Value float32

	for {

		fmt.Print("Enter the value of First Number: ")
		fmt.Scan(&Num1)
		fmt.Print("Enter teh value of Second Number: ")
		fmt.Scan(&Num2)
		fmt.Print("Enter Operator: ")
		fmt.Scan(&Operator)
		Taken = Calculation(Num1, Num2, Operator)
		fmt.Print("Calculated Value: ", Taken, "\n")

		fmt.Println("Power")
		fmt.Print("Enter the value of Base: ")
		fmt.Scan(&Base)
		fmt.Print("Enter the value of Power: ")
		fmt.Scan(&Power)
		Power_Function(Base, Power)

		fmt.Println("Check Power Function")
		fmt.Print("Enter the value of Base: ")
		fmt.Scan(&Base)
		fmt.Print("Enter the value of Power: ")
		fmt.Scan(&Power)
		PowerAnswer := math.Pow(float64(Base), float64(Power))
		fmt.Println("Power checking value: ", PowerAnswer)

		fmt.Print("Enter a value of sqrt: ")
		fmt.Scan(&Sqrt_Value)

		SqrtAnswer := math.Sqrt(float64(Sqrt_Value))
		fmt.Printf("Sqrt Answer: %.3v\n", SqrtAnswer)

	}
}
