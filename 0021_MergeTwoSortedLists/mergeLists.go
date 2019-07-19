package mergeTwoSortedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := l1
	r := l2

	var head *ListNode = nil
	var last *ListNode = nil

	for l != nil || r != nil {
		if l == nil {
			last.Next = r
			last = last.Next
			r = r.Next
		}

		if r == nil && l != nil {
			last.Next = l
			last = last.Next
			l = l.Next
		}

		if l.Val < r.Val {
			last.Next = l
			last = last.Next
			l = l.Next
		} else {
			last.Next = r
			last = last.Next
			r = r.Next
		}

		if head == nil {
			head = last
		}
	}

	return head
}
