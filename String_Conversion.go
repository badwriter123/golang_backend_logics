package main

import (
	"fmt"
	"strconv"
)

func conv(){
	// x := "1234hello";
	// y, err := strconv.Atoi(x)
	// fmt.Println(y, err);

	x := "1111011"
	y, err := strconv.ParseInt(x, 2, 0)
	fmt.Println(y, err);
}