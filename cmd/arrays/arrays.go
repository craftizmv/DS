package main

import "github.com/craftizmv/DS/pkg/arrays"

func main() {
	//a := []int{-5, 4, 6, -1, 2, 3, -6, 7, -2}
	a := []int{-5, -4, -6, -7, -2, -3, -6, -7, -2}
	sum := arrays.FindSubarraySum(a)
	println(sum)
}
