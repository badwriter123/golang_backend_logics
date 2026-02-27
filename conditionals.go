// Following are the comparison operators
// < (less than)
// > (Greater than)
// <= >= (Less-than-or-equal-to and Greater-than-or-equal-to)
// == (Double equal to) "===" is not present in golang as it is statically typed, meaning if we are comparing two types it does not work anyways
// != (Not Equal to)
package main

import(
	"fmt"
)

func conditions(){
	x := uint(8)
	y := 10
	z := x < uint(y)
	fmt.Println(z)

}