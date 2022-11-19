// GCD - Greatest Common Divisor
/*

 */

package main

import "fmt"

func main() {
	fmt.Println("LOG-Start : Main function")
	fmt.Println("Greatest common Divisor is with gcd() ", gcd(5, 10))
	fmt.Println("Greatest common Divisor is with gcd2() ", gcd2(125, 260))
	return
}

func gcd(iparam1 int, iparam2 int) int {
	fmt.Println("LOG-Start : gcd is in progress")
	var irange = 1
	var cgcd = 1

	if iparam1 <= iparam2 {
		irange = iparam1
	} else {
		irange = iparam2
	}

	for i := 1; i <= irange; i++ {
		if iparam1%i == 0 && iparam2%i == 0 {
			cgcd = i
			// fmt.Println("cgcd ", cgcd)
		}
	}

	fmt.Println("LOG-Exit : gcd")
	return cgcd
}

func gcd2(x int, y int) int {
	fmt.Println("gcd2 - another approach to find the gcd")
	fmt.Println("*********************")
	for y != 0 {
		fmt.Println("x and y are ", x, y)
		x, y = y, x%y

	}
	return x
}
