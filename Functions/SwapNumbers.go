package main

import "fmt"

func main() {
	fmt.Println("LOG - main function started")
	x, y := 10, 15
	fmt.Println("x,y ", x, y)
	x, y = swap(x, y)
	fmt.Println("After swap - x,y ", x, y)

	fmt.Println("Simple Logic ")
	x = x + y
	y = x - y
	x = x - y
	fmt.Println("values after swap -x, y", x, y)

	var string1, string2 string = "Umar", "Dafedar"
	fmt.Println("String1, String2 ", string1, string2)
	string1, string2 = swapString(string1, string2)
	fmt.Println("After swap String1, String2 ", string1, string2)
	fmt.Println("LOG - main function ending")
}

func swap(a int, b int) (int, int) {
	fmt.Println("LOG - swap function Started")
	fmt.Println("LOG - swap function Ending")
	return b, a
}

func swapString(a string, b string) (string, string) {
	fmt.Println("LOG - swap function for strings")
	return b, a
}
