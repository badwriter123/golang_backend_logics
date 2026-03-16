package main

import "fmt"

func arr() {
// 	var arr [2]bool
// 	fmt.Println(arr)

	// arr := [2][2]int{{1,2}, {0,0}};

	arr := [...][2]int{{1,2}, {2,2}, {3,2}};
	arr[0] = [2]int{10, 11};
	fmt.Println(len(arr));
}