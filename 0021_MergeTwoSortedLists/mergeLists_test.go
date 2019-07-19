package mergeTwoSortedLists

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getList(in string) *ListNode {
	var head *ListNode
	var current *ListNode

	for i, c := range in {
		v, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}

		ln := &ListNode{
			Val: v,
		}

		if i == 0 {
			head = ln
			current = ln
			continue
		}

		current.Next = ln
		current = current.Next
	}
	return head
}
