package main

import "github.com/craftizmv/DS/pkg/twopointers"

/*func main() {
	//tp := t.NewTP()
	////input := []int16{1, 2, 3, 4, 5}
	////rA := tp.ReverseArray(input)
	////fmt.Printf("val is +%v", rA)
	//
	//input2 := []int{1, 2, 3, 4, 5}
	//result := tp.FindSumOfThree(input2, 4)
	//println("Sum Of Three Result", result)
}*/

// main - sum of 3
/*func main() {
	arr := []int{738, 286, 814, 742, 910, -551, -500, 22, -794, -639, -167, -555, -117, -978, -741, -373, -999, 455, 248, 321, -569, -228, -969, -409, -251, 98, -54, 114, -897, -523, 554, -372}
	if t.HasTripletWithTargetSum(arr, 0) {
		fmt.Println("There is a triplet with zero sum.")
	} else {
		fmt.Println("There is no triplet with zero sum.")
	}
}*/

// Linked list - code to check working
func main() {
	ll := new(twopointers.LinkedList)
	ll.CreateLinkedList([]int{5, 4, 3, 2, 1, 12, 14})
	ll.RemoveNodeNthPos(5)
	ll.DisplayLinkedList()

}
