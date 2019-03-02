package addTwoNumbers

import "testing"

func Test_AddNumbers(t *testing.T) {
	//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	//Output: 7 -> 0 -> 8

	tc1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	tc2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	expects := ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 8,
			},
		},
	}

	result := addTwoNumbers(&tc1, &tc2)

	if result == nil {
		t.Fatal("expected a result but got nil")
	}

	expectsCurrent := &expects
	resultCurrent := result

	for expectsCurrent != nil {
		if resultCurrent == nil {
			t.Fatalf("expected val %v but got nil", expectsCurrent.Val)
		}

		if resultCurrent.Val != expectsCurrent.Val {
			t.Fatalf("expected val %v but got %v", expectsCurrent.Val, resultCurrent.Val)
		}

		expectsCurrent = expectsCurrent.Next
		resultCurrent = resultCurrent.Next
	}
}

func Test_AddTwoNumbersFiveAndFive(t *testing.T) {
	//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	//Output: 7 -> 0 -> 8

	tc1 := ListNode{
		Val: 5,
	}

	tc2 := ListNode{
		Val: 5,
	}

	expects := ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
		},
	}

	result := addTwoNumbers(&tc1, &tc2)

	if result == nil {
		t.Fatal("expected a result but got nil")
	}

	expectsCurrent := &expects
	resultCurrent := result

	for expectsCurrent != nil {
		if resultCurrent == nil {
			t.Fatalf("expected val %v but got nil", expectsCurrent.Val)
		}

		if resultCurrent.Val != expectsCurrent.Val {
			t.Fatalf("expected val %v but got %v", expectsCurrent.Val, resultCurrent.Val)
		}

		expectsCurrent = expectsCurrent.Next
		resultCurrent = resultCurrent.Next
	}
}
