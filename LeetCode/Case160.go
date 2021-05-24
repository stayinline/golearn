package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	lenA := getLength(headA)
	lenB := getLength(headB)

	curA, curB := headA, headB
	for lenA != lenB {
		if lenA > lenB {
			curA = curA.Next
			lenA--
		}
		if lenA < lenB {
			curB = curB.Next
			lenB--
		}
	}

	for curA != curB {
		curB = curB.Next
		curA = curA.Next
	}
	return curB
}

func getLength(head *ListNode) int {
	l := 0
	for cur := head; cur != nil; {
		l++
		cur = cur.Next
	}
	return l
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}
		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}
	return curA
}

func main() {

}
