package main

import "math"

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{
		Val:  math.MinInt64,
		Next: head}
	// pre记录上一个节点，cur记录当前节点，
	pre, cur := head, head.Next
	for cur != nil {
		if pre.Val <= cur.Val {
			pre = pre.Next
		} else {
			//从头开始找
			finder := newHead
			for finder.Next.Val <= cur.Val {
				finder = finder.Next
			}
			pre.Next = cur.Next
			cur.Next = finder.Next
			finder.Next = cur
		}
		cur = pre.Next
	}
	return newHead.Next
}

func main() {

}
