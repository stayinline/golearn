package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if nil == head || n <= 0 {
		return head
	}
	pre := &ListNode{0, head}
	quick, slow := head, pre

	for i := 0; quick != nil && i < n; i++ {
		quick = quick.Next
	}
	for ; quick != nil; quick = quick.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return pre.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
func main() {

}
