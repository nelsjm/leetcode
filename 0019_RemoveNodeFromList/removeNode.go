package removeNodeFromList

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	if n == 0 {
		return head
	}

	nodes := make([]*ListNode, 0)
	current := head

	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	newVals := nodes[:len(nodes)-n]
	newVals = append(newVals, nodes[len(nodes)-n+1:]...)

	for i := 0; i < len(newVals); i++ {
		var next *ListNode = nil

		if i+1 < len(newVals) {
			next = newVals[i+1]
		}

		newVals[i].Next = next
	}

	if len(newVals) == 0 {
		return nil
	}

	return newVals[0]
}
