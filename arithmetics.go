/*
Arithmetic operations are as follows
+ -> Addition
- -> Subtraction
* -> Multiplication
/ -> Division
-- -> Decrement
++ -> Increment
% -> Modulo, Remainder 
IMP -> No exponent operator present (^)
*/

package main

import (
	"fmt"
	"math"
)

func arith() {
	x := 7;     //-
	y := 9;      //| -> Addition
	z := x + y; //-

	a := 7
	b := 1000					//-> Division(Convert explicitly to float to get result in decimal points)
	c := float64(a) / float64(b)

	// Similar operations for subtraction, Multiplication, Decrement, Increment, and Modulo

	fmt.Println(z);
	fmt.Println(c);

	fmt.Println(math.Min(4,5)); // .Min results the minimum number between 4 and 5
	fmt.Println(math.Pow(4,2)); // Result of 4 ^ 4 => 16
	fmt.Println(math.Sqrt(100)); // Sqrt results in square root of a number.
	fmt.Println(math.Sqrt(169));
	fmt.Println(math.Floor(4.245)); // Round off a number
}