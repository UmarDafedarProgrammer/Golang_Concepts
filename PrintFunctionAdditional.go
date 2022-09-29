/*
Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
The format 'verbs' are derived from C's but are simpler.

General::
-----------------------------------------------------------
%v	the value in a default format, when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
%T	a Go-syntax representation of the type of the value
%%	a literal percent sign; consumes no value

Boolean::
-------------------------------------------------------------
%t	the word true or false

Integer::
--------------------------------------------------------------
%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"

Floating-point and complex constituents::
---------------------------------------------------------------
%b	decimalless scientific notation with exponent a power of two,
	in the manner of strconv.FormatFloat with the 'b' format,
	e.g. -123456p-78
%e	scientific notation, e.g. -1.234456e+78
%E	scientific notation, e.g. -1.234456E+78
%f	decimal point but no exponent, e.g. 123.456
%F	synonym for %f
%g	%e for large exponents, %f otherwise.
%G	%E for large exponents, %F otherwise
%x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
%X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

String and slice of bytes (treated equivalently with these verbs)::
-------------------------------------------------------------------
%s	the uninterpreted bytes of the string or slice
%q	a double-quoted string safely escaped with Go syntax
%x	base 16, lower-case, two characters per byte
%X	base 16, upper-case, two characters per byte

Pointer::
-------------------------------------------------------
%p	base 16 notation, with leading 0x
The %b, %d, %o, %x and %X verbs also work with pointers,
formatting the value exactly as if it were an integer.


The default format for %v is:
------------------------------------------------------
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p

Precision for floating point number
-----------------------------------------------------
%f     default width, default precision
%9f    width 9, default precision
%.2f   default width, precision 2
%9.2f  width 9, precision 2
%9.f   width 9, precision 0

*/

package main

import "fmt"

// Main Function
func main() {
	// Declare a variable
	var myvar int = 10

	// Printing int values
	fmt.Printf("Printing integer value using %%d %d \n", myvar)
	// %v	the value in a default format. When printing structs, the plus flag (%+v) adds field names
	fmt.Printf("Printing integer value using %%v %v \n", myvar)
	// %#v	a Go-syntax representation of the value
	// %%	a literal percent sign; consumes no value
	fmt.Printf("Printing integer value using %%#v %#v \n", myvar)
	fmt.Printf("Printing integer value using %%+v %+v \n", myvar)
	// %T	a Go-syntax representation of the type of the value
	fmt.Printf("Printing Type of myvar using %%T %T \n", myvar)

	// Print the boolean values true/false
	var mybool bool = true
	// %t	the word true or false
	fmt.Printf("Printing the boolean value using %%t %t \n", mybool)
}
