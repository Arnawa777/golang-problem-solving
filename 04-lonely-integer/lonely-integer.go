package main

import (
	"fmt"
)

/*
Given an array of integers, where all elements but one occur twice, find the unique element.

# Example {2, 3, 2, 3, 10}

The unique element is 10.

int a[n]: an array of integers
*/
func main() {
	var arr = [...]int32{2, 3, 2, 3, 10}
	var n = len(arr)

	var res int32 = 0

	for i := 0; i < n; i++ {
		res = res ^ arr[i]
		fmt.Println("Array :", arr[i], "Res is", res)
	}

	fmt.Println(res)
}
