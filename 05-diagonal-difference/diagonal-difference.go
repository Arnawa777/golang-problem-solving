package main

import (
	"fmt"
)

func main() {

	var arr = [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12}}

	var d1, d2 int32 = 0, 0
	n := len(arr)
	fmt.Println(n)

	for i := 0; i < n; i++ {

		// d1 += arr[i][i]
		// d2 += arr[i][n-i-1]

		for j := 0; j < n; j++ {
			fmt.Println("i", i)
			fmt.Println("j", j)

			// finding sum of primary diagonal
			if i == j {
				fmt.Println("d1", d1)
				d1 += arr[i][j]
			}

			// finding sum of secondary diagonal
			if i == n-j-1 {
				d2 += arr[i][j]
				fmt.Println("d2", d2)
			}
		}
	}

	fmt.Println("d1", d1)
	fmt.Println("d2", d2)
}
