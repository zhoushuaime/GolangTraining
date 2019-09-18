package linklist

import (
	"fmt"
	"testing"
)

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// TestAddTwoNumbers ...
func TestAddTwoNumbers(t *testing.T) {

	cases := []struct {
		name   string
		input1 *ListNode
		input2 *ListNode
		expect *ListNode // 2-4-3 5-6-4  342 + 465 = 807
	}{{"1 test 1", &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
		&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
		&ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}}},}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := addTwoNumbers(c.input1, c.input2)
			t.Log(got.Val, got.Next.Val, got.Next.Next.Val)
			fmt.Println()
			if !isEqual(got, c.expect) {
				t.Errorf("expect :%v,but got :%v,input data:%v", c.expect, got, c.input1)
			}
		})
	}
}

// isEqual ...
func isEqual(l1 *ListNode, l2 *ListNode) bool {

	if l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil && l2 != nil {
		return false
	}

	if l1 != nil && l2 == nil {
		return false
	}
	return true
}

// addTwoNumbers ...
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{Val: 0, Next: nil}
	n1, n2, tmp := l1, l2, node
	sum := 0
	for n1 != nil || n2 != nil {
		sum /= 10
		if n1 != nil {
			sum += n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			sum += n2.Val
			n2 = n2.Next
		}
		tmp.Next = &ListNode{Val: sum % 10}
		tmp = tmp.Next
	}
	if sum/10 != 0 {
		tmp.Next = &ListNode{Val: 1}
	}
	return node.Next
}
