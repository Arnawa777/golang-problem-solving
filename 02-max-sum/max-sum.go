package main

import (
	"fmt"
)

/*
For example, arr[]={1,3,5,7,9}. Our minimum sum is 1+3+5+7 = 16 and our maximum sum is 3+5+7+9 = 24 . We would print

16 24

Our initial numbers are 1 ,2 ,3 ,4 , and 5. We can calculate the following sums using four of the five integers:

If we sum everything except 1, our sum is 2 + 3 + 4+ 5 = 14 .
If we sum everything except 2, our sum is 1 + 3 +4 + 5 = 13 .
If we sum everything except 3, our sum is 1+ 2+ 4 + 5 = 12.
If we sum everything except 4, our sum is 1+2+3+ 5 = 11.
If we sum everything except 5, our sum is 1+2+3+4 = 10.

Hints: Beware of integer overflow! Use 64-bit Integer.”
*/

func main() {
	var arr = [...]int32{2, 3, 9, 0, 6}

	var sum int64
	for i := 0; i < len(arr); i++ {
		sum += int64(arr[i])
	}
	fmt.Println("sum outside", sum)

	var min int64
	var max int64

	for i := 0; i < len(arr); i++ {

		fmt.Println("it's i ", i)
		/*
			jika min kosong atau sum <= min
			Baris 10 sampai 12 logika nya seperti ini jika sum-n[i] lebih kecil dari min atau min sama dengan 0 maka update variabel min menjadi sum-n[i].
		*/
		if sum-int64(arr[i]) <= min || min == 0 {
			min = sum - int64(arr[i])

			fmt.Println("it's SUM ", sum)
			fmt.Println("it's Min ", min)
		}

		/*
			Baris 10 sampai 12 logika nya seperti ini jika sum-n[i] lebih beasar dari max atau max sama dengan 0 maka update variabel max menjadi sum-n[i].
		*/
		if sum-int64(arr[i]) > max || max == 0 {
			max = sum - int64(arr[i])

			fmt.Println("it's Max ", max)
		}
	}

	fmt.Println(min, max)
}
