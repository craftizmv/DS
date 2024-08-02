package main

import (
	t "github.com/craftizmv/DS/pkg/twopointers"
)

func main() {
	tp := t.NewTP()
	//input := []int16{1, 2, 3, 4, 5}
	//rA := tp.ReverseArray(input)
	//fmt.Printf("val is +%v", rA)

	input2 := []int{1, 2, 3, 4, 5}
	result := tp.FindSumOfThree(input2, 9)
	println("Sum Of Three Result", result)

}
