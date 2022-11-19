/*
	1. Called function can not access the data of calling function directly
	2. Calling Function has to pass the data to called function as a input argument
	3. Variables mustbe passed explicitly
*/
package main

import "fmt"

func main() {
	fmt.Println("LOG : Main function")
	var Name = "Umar"
	fmt.Println("Name of the Person is ", Name)
	fmt.Println("Address of the Name is ", &Name)
	Read(Name)
	fmt.Println("Name of the Person is ", Name)
	return
}

func Read(name string) {
	// fmt.Println("Trying to read the name of the person in Read() ", Name)
	// Can not access the Name from calling function
	fmt.Println("Reading the name ", name)
	name = "Rehan"
	fmt.Println("Reading the name ", name)
	fmt.Println("Address of the name is", &name)

	fmt.Println()
	hello()
	hello2("Umar")
	hello3("Umar", "Dafedar")
}

func hello() {
	fmt.Println("Hello !!! Welcome to the function")
}

func hello2(name string) {
	fmt.Println("Hello !!!", name)
}

func hello3(name string, surname string) {
	fmt.Println("Hello !!!", name, surname)
}
