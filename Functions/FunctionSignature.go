/*
Function Signature: Taking a function with 2 input arguments of type int an 1 output argument of type int

	1. func name(a int, b int) int {
		// Function body
		return value
	}

	2. func name(a int, b int) (z int) {
		z=a*b
		return
	}

	3. func name(int,int) int {
		return 0
	}

	4. func name(a int, _ int) int {
		return a
	}
*/

package main

import "fmt"

func main() {
	fmt.Println("LOG : Main Function")
	var val1, val2 int
	val1 = 10
	val2 = 21
	addition := add(val1, val2)
	subtraction := subtract(val1, val2)
	multiplication := multiply(10, 0)
	returnZero := zero(10, 10)
	fmt.Println("LOG : Back to Main Function")
	fmt.Println("LOG : Addition of values is ", addition)
	fmt.Printf("Type of add function %T \n", add)
	fmt.Println("LOG : Subtraction of values is ", subtraction)
	fmt.Printf("Type of subtract function %T \n", subtract)
	fmt.Println("LOG : Multiplication of value is ", multiplication)
	fmt.Printf("Type of multiply function %T \n", multiply)
	fmt.Println("LOG : Zero Function result is ", returnZero)
	fmt.Printf("Type of zero function %T\n", zero)
	return
}

func add(val1 int, val2 int) int {
	fmt.Println("LOG : add function is called")
	var result int
	result = val1 + val2
	return result
}

func subtract(val1 int, val2 int) (result int) {
	fmt.Println("LOG : Subtract funtion is called ")
	result = val2 - val1
	return
}

func multiply(val1 int, _ int) (result int) {
	fmt.Println("LOG : Multiply function is called ")
	result = val1 * 10
	return
}

func zero(int, int) int {
	fmt.Println("LOG : Zero function is called ")
	return 0
}
