package main

import "fmt"

/*
There are three print functions
1. Print
2. Println
3. Printf

- These are defined, in package fmt
- if fmt package is not imported, you can use print() and println() functions

*/

func main() {
	// ###  Print()
	// 1. No, new lines added
	// 2. if you are trying to print multiple values, they are not seperated by space by default
	fmt.Print()
	fmt.Print("Printing first comment,")
	fmt.Print("Printing the Second comment!!!")
	fmt.Print("\n")
	//Result : Printing first comment,Printing the Second comment!!!

	/// ### Println()
	// 1. new line is added to the comment
	// 2. If you are trying to print multiple values, then these are seperated by space by default
	fmt.Println("Println() :", "First comment")
	fmt.Println("Println() : Second comment")
	// Result : Println() : First comment

	// ### Printf()
	/* 	1. Used to format the data before printing it to the console/Std output
	2. No new line is added
	v – formats the value in a default format
	d – formats decimal integers
	g – formats the floating-point numbers
	b – formats base 2  numbers
	o – formats base 8 numbers
	t – formats true or false values
	s – formats string values
	*/
	intvalue := 999
	floatvalue := 3.32423444
	var ToF bool = true
	var myString = "Hello Anam"
	fmt.Printf("Printing integer value, %d \n", intvalue)
	fmt.Printf("Print the Binary eqivalent : %b \n", intvalue)
	fmt.Printf("Print the floating value : %f \n", floatvalue)
	fmt.Printf("Print the floating value : %g \n", floatvalue)
	fmt.Printf("Print the floating value : %v \n", floatvalue)
	fmt.Printf("Print the boolean value : %v \n", ToF)
	fmt.Printf("Print the string value : %v \n", myString)
}
