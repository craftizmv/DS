package main

import "fmt"

func main() {
	abc := []int{1, 2, 3}
	modify(abc)

	//fmt.Println(abc)

	abc = nil
	fmt.Println(len(abc))
}

func modify(a []int) {
	//a = append(
	//	a,
	//	4,
	//)

	//fmt.Println(len(a))

	a[1] = 4

	//fmt.Println(a)
	//
	//return a
}
