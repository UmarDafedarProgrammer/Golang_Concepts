package main

import "fmt"

func main() {
	fmt.Println("LOG - Main Start")
	calculator()
	fmt.Println("LOG - Main End")
}

func calculator() {
	var num1, num2 float32
	var operator string
	fmt.Println("Enter the First argument ")
	fmt.Scanln(&num1)
	fmt.Println("Enter the Second argument ")
	fmt.Scanln(&num2)
	fmt.Println("Enter the operation")
	fmt.Scanln(&operator)

	for ok := true; ok; {
		ok = false
		switch operator {
		case "+":
			num1 = num1 + num2
		case "-":
			num1 = num1 - num2
		case "*":
			num1 = num1 * num2
		case "/":
			if num2 != 0 {
				num1 = num1 / num2
			} else {
				fmt.Println("invalid denominator value")
				fmt.Println("Final Result is ", num1)
				return
			}
		default:
			{
				fmt.Println("Requested operator is not supported, yet")
				return
			}
		}
		fmt.Println(num1)
		fmt.Scanln(&operator)
		if operator == "+" || operator == "-" || operator == "*" || operator == "/" {
			fmt.Println("Enter the value")
			fmt.Scanln(&num2)
			ok = true
		}
	}
	fmt.Println("Final Result is ", num1)
}
