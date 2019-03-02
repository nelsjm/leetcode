package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return buildNode(l1, l2, 0)
}

func buildNode(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	v1 := 0
	var l1Next *ListNode
	if l1 != nil {
		v1 = l1.Val
		l1Next = l1.Next
	}

	v2 := 0
	var l2Next *ListNode
	if l2 != nil {
		v2 = l2.Val
		l2Next = l2.Next
	}

	sum := v1 + v2 + carry
	nextCarry := 0
	if sum >= 10 {
		sum = sum - 10
		nextCarry = 1
	}

	currentNode := ListNode{
		Val:  sum,
		Next: buildNode(l1Next, l2Next, nextCarry),
	}

	return &currentNode
}
