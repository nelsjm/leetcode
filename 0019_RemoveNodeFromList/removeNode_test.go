package removeNodeFromList

import (
	"fmt"
	"testing"
)

func Test_RemoveNthNode(t *testing.T) {
	head := buildNodes([]int{1})
	v := removeNthFromEnd(head, 1)
	if v == nil {
		return
	}

	i := 0
	for v != nil && i < 100 {
		fmt.Printf("%v -> ", v.Val)
		v = v.Next
		i++
	}
	fmt.Println("")

}

func buildNodes(in []int) *ListNode {
	var current *ListNode = nil
	var head *ListNode = nil
	for _, v := range in {
		n := &ListNode{
			Val: v,
		}

		if current == nil {
			current = n
			head = n
			continue
		}

		current.Next = n
		current = current.Next
	}

	return head
}

func extractValues(h *ListNode) []int {
	vals := make([]int, 0)

	for h != nil {
		vals = append(vals, h.Val)
		h = h.Next
	}

	return vals
}

func areTheSame(l []int, r []int) bool {
	if len(l) != len(r) {
		return false
	}

	return true
}
