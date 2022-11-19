/* Fibonnaci series
0,1,1,2,3,5,8,13,21,34......
*/
package main

import "fmt"

func main() {
	fmt.Println("LOG - Main Start")
	fibonnaci(10)
	fmt.Println("LOG - Main End")
	return
}

func fibonnaci(length int) {
	fmt.Println("LOG - fibonnaci Start")
	var x, y = 0, 1
	if length <= 1 {
		fmt.Print(x)
		return
	}
	fmt.Print(x, " ")
	for i := 1; i < length; i++ {
		x, y = y, x+y
		fmt.Print(x, " ")
	}
	fmt.Println()

	fmt.Println("LOG - fibonnaci End")
}
