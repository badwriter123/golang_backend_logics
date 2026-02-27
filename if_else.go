package main

import (
	"fmt"
)

func logics() {

	// If Else logics
	// x := 2;
	// if x > 2 {
	// 	fmt.Println("run");
	// } else if x < 2 {
	// 	fmt.Println("Hold")
	// } else {
	// 	fmt.Println("Completed...")
	// }

	// Switch statements -> Default break is used to avoid falling to other cases. To purposefully execute multiple cases, use "fallthrough" keyword
	// a := 1;
	
	// switch  {
	// 	case a == 1:
	// 		fmt.Println("one");
	// 		fallthrough
	// 	case a <= 1 :
	// 		fmt.Println("two");
	// 		// fallthrough
	// 	default:
	// 		fmt.Println("default run...")		
	// }

	// Another example of fallthrough
	a := -2;

	switch {
		case a < 1:
			fmt.Println("a is less than 1..");
			fallthrough
		case a < 0:
			fmt.Println("a is less than 0..");
			// fallthrough
		case a < -3:
			fmt.Println("a is equal to -2...");
		default:
			fmt.Println("Default run...")		
	}
}