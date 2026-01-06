package buildstronghold

import (
	"fmt"

	"github.com/lossdev/stack"
)

func SortStackDemo() {
	st := stack.NewStack(stack.Int)
	st.Push(4)
	st.Push(3)
	st.Push(7)
	st.Push(1)
	st.Push(8)

	SortStack(st)

	for st.Size() != 0 {
		val, _ := st.Pop()
		fmt.Println(val)
	}
}

func SortStack(st *stack.Stack) {
	// bc
	if st.Size() == 0 {
		return
	}

	if top, e := st.Pop(); e != nil {
		// the error is possible when stack is empty ...
		// just return from this recursive path.
		return
	} else {
		// calling the rec. func.
		SortStack(st)

		// insert the numbers ...
		sortedInsert(st, top.(int))
	}

}

// TODO - Handler error condition if any.
// Goal insert the element at right location.
func sortedInsert(st *stack.Stack, element int) {
	// 1. If stack is empty of top is lesser than element then insert
	top, _ := st.Peek()
	if st.Size() == 0 || element > top.(int) {
		// simply insert the element
		_ = st.Push(element)
		return
	}

	// 2. pop the element.
	st.Pop()

	// 3. Again attempt to insert.
	sortedInsert(st, element)

	// 4. push the popped element at this level.
	st.Push(top)
}
