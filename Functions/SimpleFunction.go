/*
Function is set of statements grouped together and are executed on the function call
- Reusability of the code
- Easy maintainance
- To modularize the problem
- Go requires explicit returns, i.e. it won't automatically return the value of the last expression.

Syntax:
	func name(parameter-list) (result-list) {
		Funcion body
	}

main() and init() are automatically called
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("LOG - Main Function")
	printer()                                                        // Function call
	fmt.Printf("LOG - Address of the printer function :%T", printer) // Print the Type of function
}

// Function with no input arguments and no return value -Function Definition
func printer() {
	fmt.Println("LOG - Printer Function is called")
}
