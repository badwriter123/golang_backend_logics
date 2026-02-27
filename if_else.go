package main

import (
	"fmt"
)

func logics() {
	x := 2;
	if x > 2 {
		fmt.Println("run");
	} else if x < 2 {
		fmt.Println("Hold")
	} else {
		fmt.Println("Completed...")
	}
}