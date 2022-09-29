package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Println("String Vriable")
	var string1 string = "Umar"
	var string2 string = "Anam"
	fmt.Printf("print the string 1 :%s \n", string1)
	fmt.Printf("print the string 2 :%s \n", string2)

	fmt.Println("Size of the string1 :", unsafe.Sizeof(string1))
	fmt.Println("Size of the string2 :", unsafe.Sizeof(string2))

	fmt.Println(string1 + string2)
	fmt.Printf("Type is %T \n", string1)

	character := 'c'
	fmt.Println("character is ", character)
	fmt.Println("Type of character is", reflect.TypeOf(character))
	fmt.Printf("Lets print the character %c \n", character)
	fmt.Printf("Lets print the quoted character %q \n", character)

	character1 := "abc"
	fmt.Println("character1 is ", character1)
	fmt.Println("Type of character1 is", reflect.TypeOf(character1))
	fmt.Printf("Lets print the character1 %s \n", character1)
	fmt.Printf("Lets print the quoted character1 %q \n", character1)

	fmt.Printf("Lets print the Width 6, right justify %6s b\n", character1)
	fmt.Printf("Lets print the Width 6, left justify %-6s b \n", character1)
	fmt.Printf("Lets print Hex dump of byte values %x b \n", character1)
	fmt.Printf("Lets print Hex dump of byte values with spaces % x b \n", character1)
}
