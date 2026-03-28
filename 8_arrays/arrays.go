package main

import "fmt"

//numbered sequence of specific length
func main(){

	//by default the elemnt in an integer array are 0
	// var nums [4]int

	// nums[0] = 1


	//fmt.Println(len(nums))
	// fmt.Println(nums)

	// var vals [4]bool
	// vals[2] = true
	//by default the elemnt in a boolean array are false
	// fmt.Println(vals)

	// to declare it in single line
	// nums := [3]int{1 , 2, 3}
	// fmt.Println(nums)

	// 2d arrays
	nums := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(nums)
}