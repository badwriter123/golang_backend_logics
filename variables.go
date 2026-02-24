package main

import (
	"fmt"
)

// Two types to declare variables
// 1. var -> mutable
// 2. const -> immutable
func variables1(){
	var x string = "Hello World1" // string type
	var y  uint = 22; // unsigned Intiger type -> can be of uint32, uint64, uint16, uint8 etc. Cannot hold negative value since it is unsigned.
	const z uint = 32; // value of const cannot be changed / mutated.

	// explicit assignment
	a := 22; // explicitly assignes type to be int.
	fmt.Println(x);
	fmt.Println(y);
	fmt.Println(z);
	fmt.Println(a); 
	fmt.Printf("%T",a); // printing type of 'a' variables results in type 'int'.
	
}