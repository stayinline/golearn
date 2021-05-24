package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	newHead := lists[0]
	for i := 1; i < len(lists); i++ {
		newHead = mergeTwoLists(lists[i], newHead)
	}
	return newHead
}

//递归法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}
	var cur *ListNode
	if l1.Val < l2.Val {
		cur = l1
		cur.Next = mergeTwoLists(l1.Next, l2)
	} else {
		cur = l2
		cur.Next = mergeTwoLists(l1, l2.Next)
	}
	return cur
}

//循环法
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}
	var pre = ListNode{
		Val:  0,
		Next: nil}
	cur := &pre
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	} else if l2 == nil {
		cur.Next = l1
	}
	return pre.Next
}

func main() {

}
