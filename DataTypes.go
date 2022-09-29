/*
* Data Type specifies the type of data variable can hold
* It also specifies, the size of memory which needs to be allocated for the variable

A> Data Types are divided into 4 categories
---------------------------------------------
1. Basic Type
2. Aggregate Type
3. Reference Type
4. Interface Type

	a> Basic data types are further categorised into 3 types
	--------------------------------------------------------
	1. Numbers
	2. Booleans
	3. Strings

		a.1> Numbers
		--------
		Numbers are again categorised into 3 sub categories
 		1. integers
 		2. floats-point numbers
 		3. complex numbers

			a.1.1> Integers - Signed/unsigned
			-----------------------------
			int - 4/8 bytes, depends up on the processor
			uint/uintptr - unsigned int
			int8 - 8-bit signed integer
			int16 - 16 bit signed integer
			int32/rune - 32 bit signed integer - rune in Go is a data type that stores codes that represent Unicode characters
			int64 - 64 bit signed integer
			uint8/byte - 8 bit unsigned integer
			unit16 - 16 bit unsigned integer
			unit32 - 32 bit unsigned integer
			unit64 - 64 bit unsigned integer
			uintptr - For representing pointer variable

			a.1.2> Floating-point numbers
			-------------------------
			float32 - 32 bit floating point
			float64 - 64 bit floating point

			a.1.3> Complex Numbers
			--------------------------------
			The complex numbers are divided into two parts are shown in the below table.
			float32 and float64 are also part of these complex numbers.
			The in-built function creates a complex number from its imaginary and real part and in-built imaginary
			and real function extract those parts.
			complex64 	- Complex numbers which contain float32 as a real and imaginary component
			complex128	- Complex numbers which contain float64 as a real and imaginary component

		a.1.2> Boolean
		The boolean data type represents only one bit of information either true or false.
		The values of type boolean are not converted implicitly or explicitly to any other type.
		Default value is false

		a.1.3> Strings
		The string data type represents a sequence of Unicode code points. Or in other words,
		we can say a string is a sequence of immutable bytes, means once a string is created you cannot change
		that string.
		A string may contain arbitrary data, including bytes with zero value in the human-readable form.

*/
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var mybool bool
	fmt.Printf("Printing value from boolean variable %t \n", mybool)
	var Vint8 int8
	fmt.Printf("int8 - %d \n", Vint8)
	var Vint16 int16
	fmt.Printf("int16 - %d \n", Vint16)
	var Vint32 int32
	fmt.Printf("int32 - %d \n", Vint32)
	var Vint64 int64
	fmt.Printf("int64 - %d \n", Vint64)
	MyVar := 10
	fmt.Printf("var data type is %T \n", MyVar)

	var Vuint8 uint8
	fmt.Printf("uint8 - %d \n", Vuint8)
	var Vuint16 uint16
	fmt.Printf("int16 - %d \n", Vuint16)
	var Vuint32 int32
	fmt.Printf("int32 - %d \n", Vuint32)
	var Vuint64 uint64
	fmt.Printf("uint64 - %d \n", Vuint64)
	MyVar1 := -10
	fmt.Printf("var1 data type is %T \n", MyVar1)

	var Vrune rune
	fmt.Printf("rune - %d \n", Vrune)
	fmt.Println("Sizeof rune is ", reflect.TypeOf(Vrune).Size())

	var Vbyte byte
	fmt.Printf("byte - %d \n", Vbyte)
	fmt.Println("Sizeof byte is ", reflect.TypeOf(Vbyte).Size())

	var Vptr uintptr
	fmt.Printf("uintptr - %d \n", Vptr)
	fmt.Println("Sizeof uintptr is ", reflect.TypeOf(Vptr).Size())

	// boolean variables
	var Vbool bool
	fmt.Printf("Value of bool is - %t \n", Vbool)
	fmt.Println("Size of boolean is ", reflect.TypeOf(Vbool).Size())

	// Float precisions and variables
	// By default 5 decimal precision
	VarFloat := 3.12345667
	fmt.Printf("Value of VarFloat is %f \n", VarFloat) //3.123457
	fmt.Printf("Value of VarFloat is %v \n", VarFloat) // 3.12345667
	fmt.Println("Size of the VarFloat is ", reflect.TypeOf(VarFloat).Size())

	var float1 float64 = 3.14e+39
	fmt.Printf("float1 value is %f ", float1)
	fmt.Println("Size of the float1 is ", reflect.TypeOf(float1).Size())

	var float2 float32 = 3.14e+12
	fmt.Printf("float2 value is %f ", float2)
	fmt.Println("Size of the float2 is ", reflect.TypeOf(float2).Size())
	fmt.Printf("float2 value is %e \n", float2)

	float2 = 66.786556
	fmt.Printf("float2 value -default width and default precision %f \n", float2)
	fmt.Printf("float2 value -default width and precision 2:%.2f \n", float2)
	fmt.Printf("float2 value -default width 8 and precision 2:%8.2f \n", float2)
	fmt.Printf("float2 value -default width 8 and precision 2:%g \n", float2)

	// Complex Numbers
	var a float32 = 3
	var b float32 = 5

	//Initialize-1
	c := complex(a, b)

	//Initialize-2
	var d complex64
	d = 4 + 5i

	//Print Size
	fmt.Printf("c's size is %d bytes\n", unsafe.Sizeof(c))
	fmt.Printf("d's size is %d bytes\n", unsafe.Sizeof(d))

	//Print type
	fmt.Printf("c's type is %s\n", reflect.TypeOf(c))
	fmt.Printf("d's type is %s\n", reflect.TypeOf(d))

	//Operations on complex number
	fmt.Println(c+d, c-d, c*d, c/d)
}
