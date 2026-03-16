package main

import "fmt"

func arr() {
// 	var arr [2]bool
// 	fmt.Println(arr)

	// arr := [2][2]int{{1,2}, {0,0}};

	// arr := [...][2]int{{1,2}, {2,2}, {3,2}};
	// arr[0] = [2]int{10, 11};
	// fmt.Println(len(arr));

	// Looping through the array
	// for value := range arr {
	// 	fmt.Println(value)
	// }

	// Looping through nested arrays
	// arr1 := [...][2]int{{1,2}, {2,4}, {2,2}};

	// for _, nested := range arr1 {
	// 	for _, value := range nested {
	// 		fmt.Println(value);
	// 	}
	// }

	arr := [...][2]int{{1,2}, {2,2}, {3,2}}
	test(arr)
	fmt.Println(arr);
}

func test(arr [3][2]int) {
	arr[0] = [2]int{100, 100}
}