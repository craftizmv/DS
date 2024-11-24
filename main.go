package main

import (
	"fmt"
	"github.com/craftizmv/DS/pkg/arrays"
)

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
//func main() {
//	ll := new(two_pointers.LinkedList)
//	ll.CreateLinkedList([]int{5, 4, 3, 2, 1, 12, 14})
//	ll.RemoveNodeNthPos(5)
//	ll.DisplayLinkedList()
//}

// check 3 sum problem and merge array problem
func main() {
	//interval := [][]int{
	//	{1, 3},
	//	{2, 4},
	//	{4, 5},
	//	{6, 10},
	//}
	//result := merge_interval.MergeIntervals(interval)
	//fmt.Printf("Result +%v", result)

	//result, err := arrays.ThreeSumBruteForceTriplet([]int{-1, 0, 1, 2, -1, -4}, 0)
	//if err != nil {
	//	fmt.Print("Error : ", err)
	//	return
	//}
	//fmt.Println("Result", result)

	//result, err := arrays.ThreeSumOptimisedTwoPointer([]int{-1, 0, 1, 2, -1, -4}, 0)
	//if err != nil {
	//	fmt.Print("Error : ", err)
	//	return
	//}
	//fmt.Println("Result", result)
	//
	//result, err = arrays.ThreeSumOptimisedTwoPointer([]int{1, 2, 3}, 6)
	//if err != nil {
	//	fmt.Print("Error : ", err)
	//	return
	//}
	//fmt.Println("Result", result)

	result := arrays.ThreeSumOptimisedTwoPointer([]int{1, 2, 3, 4, 5}, 9)
	//if err != nil {
	//	fmt.Print("Error : ", err)
	//	return
	//}
	fmt.Println("Result", result)

	//result, err = arrays.ThreeSumOptimisedTwoPointer([]int{2, -1, 0, 1, 2}, 0)
	//if err != nil {
	//	fmt.Print("Error : ", err)
	//	return
	//}
	//fmt.Println("Result", result)

}

// main - for rate limiter
//func main() {
//	//rate_limiter.TickerBasedRateLimiter()
//	//rate_limiter.BurstyRateLimiter()
//	rate_limiter.UberRateLimiter()
//}

// producer consumer
//func main() {
//	random_practice.SimpleProducerConsumer()
//}

//func main() {
//	//result := basics.ReverseDigits(8749348034803)
//	//fmt.Println("Result is ", result)
//
//	tests.FilterCSV("")
//}

// worker pool main - unbuffered.
//func main() {
//	workerPool := NewWorkerPool(5)
//	workerPool.Start()
//
//	for i := 0; i < 10; i++ {
//		//fmt.Println("Submitting task with ID : ", i)
//		workerPool.Submit(func() (interface{}, error) {
//			return someExpensiveOperation(i), nil
//		})
//
//		//fmt.Println("Finished submitting task with ID : ", i)
//	}
//
//	fmt.Println("Going to listen for the result channel")
//	for i := 0; i < 10; i++ {
//		var res = workerPool.GetResult()
//		fmt.Printf("Result: %v", res)
//	}
//}

// buffered worker pool
//func main() {
//	workerPool := worker_pool_buffered.NewWorkerPool(5)
//	workerPool.Start()
//
//	// IF WE USE 11 INSTEAD OF 10 TASKS THEN AGAIN IT WILL END UP IN DEADLOCK.
//	// IF WE USE 10 WITH 5 WORKER AND RESULT CHANNEL AS BUFFERED WITH 5, THEN
//	// AS OTHER 5 TASK WILL BE PICKED UP HENCE, IT WILL UNBLOCK LINE NUMBER 109
//	// AND HENCE IT WILL START READING RESULTS.
//
//	// IDEALLY - WE NEED TO IMPLEMENT IT IN SUCH A WAY THAT... tasks can be listened/consumed
//	// in async way. - TODO : write a gist for that.
//
//	for i := 0; i < 10; i++ {
//		//fmt.Println("Submitting task with ID : ", i)
//		workerPool.Submit(func() (interface{}, error) {
//			return someExpensiveOperation(i), nil
//		})
//
//		//fmt.Println("Finished submitting task with ID : ", i)
//	}
//
//	fmt.Println("Going to listen for the result channel")
//	for i := 0; i < 10; i++ {
//		var res = workerPool.GetResult()
//		fmt.Printf("Result: %v", res)
//		fmt.Println("   ")
//	}
//}
//
//func someExpensiveOperation(i int) interface{} {
//	time.Sleep(1 * time.Second)
//	return true
//}

// morris traversal
//func main() {
//	root := &traversals.Node{
//		Data: 1,
//	}
//	root.Left = &traversals.Node{
//		Data: 2,
//	}
//	root.Right = &traversals.Node{
//		Data: 6,
//	}
//	root.Right.Left = &traversals.Node{
//		Data: 9,
//	}
//	root.Right.Right = &traversals.Node{
//		Data: 7,
//	}
//
//	root.Left.Left = &traversals.Node{
//		Data: 3,
//	}
//	root.Left.Right = &traversals.Node{
//		Data: 4,
//	}
//	root.Left.Right.Left = &traversals.Node{
//		Data: 8,
//	}
//	root.Left.Right.Right = &traversals.Node{
//		Data: 5,
//	}
//
//	bt := traversals.NewBinaryTree(root)
//	err := bt.MorrisInOrderTraversal(root)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}
