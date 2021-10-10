package trees

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrderUsingQueue(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushBack(root)
	queue.PushBack(nil)

	var levelArr []int

	for queue.Len() != 0 {
		front := queue.Front()
		// process the front element in the queue and return
		if front.Value == nil {
			// then it means we have reached end of that level, remove the -100
			queue.Remove(front)

			// Add level arr to global result
			result = append(result, levelArr)

			// make new levelArr
			levelArr = []int{}

			// If there is some valid element present at front of the queue then push back nil at the end as a marker.
			if queue.Front() != nil {
				queue.PushBack(nil)
			}
		} else {
			queue.Remove(front)
			levelArr = append(levelArr, front.Value.(*TreeNode).Val)
		}

		// check if left is there
		if front.Value != nil && front.Value.(*TreeNode).Left != nil {
			queue.PushBack(front.Value.(*TreeNode).Left)
		}

		if front.Value != nil && front.Value.(*TreeNode).Right != nil {
			queue.PushBack(front.Value.(*TreeNode).Right)
		}
	}

	return result
}
