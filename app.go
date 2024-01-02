package main

import "fmt"

func main() {
	fmt.Print("lets Calculate \n")
	var firstNum int
	var secondNum int
	fmt.Print("enter the first number: ")
	fmt.Scan(&firstNum)
	fmt.Print("enter the second number: ")
	fmt.Scan(&secondNum)
	calculatedValue := calculate(firstNum, secondNum)
	fmt.Println("the calculated value is", calculatedValue)
}

func calculate(firstNum, secondNum int) int {
	return firstNum * secondNum
}
