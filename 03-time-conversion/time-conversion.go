package main

/*
Given a time in -hour AM/PM format, convert it to military (24-hour) time.

Note: - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
- 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.

Example

s= '12:01:00PM'
Return '12:01:00'.

s= '12:01:00AM'
Return '00:01:00'.
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	str = "07:05:45PM"
	// pm := strings.HasSuffix(str, "PM")
	// am := strings.HasSuffix(s, "AM")
	hour := str[:2]
	fmt.Println(hour)
	ints, _ := strconv.Atoi(hour)

	// fmt.Println(reflect.TypeOf(ints), ints)
	if str[8:9] == "P" && ints != 12 {
		ints = ints + 12
		sad := strconv.Itoa(ints)
		fmt.Println(sad + str[2:8]) //1 - 11 PM
	} else if str[8:9] == "P" && ints == 12 {
		fmt.Println(str[:8]) //12 PM
	} else if str[8:9] == "A" && ints == 12 {
		hour = "00"
		fmt.Println(hour + str[2:8]) // 12 AM
	} else if str[8:9] == "A" && ints != 12 {
		fmt.Println(hour + str[2:8]) // 1 - 11 AM
	}
	// fmt.Println(ints, str[1:9])
	// fmt.Println("get:", str[:1])
	// fmt.Println("get:", str[8:9])
	// fmt.Println("get:", str[1:2])
	// fmt.Println("output")
}
