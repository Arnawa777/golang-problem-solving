package main

import (
	"fmt"
)

func main() {
	var arr = [...]int32{2, -3, -2, 0, 3}

	var plus float32 = 0
	var minus float32 = 0
	var zero float32 = 0
	total := float32(len(arr))
	for _, n := range arr {
		if n == 0 {
			zero += 1
		} else if n > 0 {
			plus += 1
		} else if n < 0 {
			minus += 1
		}
	}

	zero = zero / total
	plus = plus / total
	minus = minus / total

	/*
		Print the following  lines, each to  decimals:

		proportion of positive values
		proportion of negative values
		proportion of zeros
	*/

	// var y float64 = float64(zero)
	fmt.Printf("%.6f \n", plus)
	fmt.Printf("%.6f \n", minus)
	fmt.Printf("%.6f \n", zero)

	fmt.Println("Total", total, "Jumlah Zero", zero, "Jumlah Positive", plus, "Jumlah Negtive", minus)

}
