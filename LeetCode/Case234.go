package main

import "fmt"

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func isPalindrome(head *ListNode2) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var nextHead *ListNode2
	if nil == fast {
		nextHead = slow
	} else {
		nextHead = slow.Next
	}

	rightHead := reverseList2(nextHead)
	fmt.Println("----------")
	printList(rightHead)

	pre := &ListNode2{
		Val:  0,
		Next: head,
	}
	for curRight, cur := rightHead, head; cur != rightHead && curRight != nil; {
		if cur.Val != curRight.Val {
			return false
		} else {
			pre = cur
			cur = cur.Next
			curRight = curRight.Next
		}
	}
	pre.Next = reverseList2(pre.Next)
	return true
}

func reverseList2(head *ListNode2) *ListNode2 {
	var pre *ListNode2
	cur := head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func main() {
	nums := []int{1, 2, 2, 1}
	arr := make([]ListNode2, len(nums))
	for i := 0; i < len(nums); i++ {
		node := ListNode2{
			Val:  nums[i],
			Next: nil,
		}
		arr[i] = node
	}

	for i := 0; i < len(arr)-1; i++ {
		arr[i].Next = &arr[i+1]
	}

	printList(&arr[0])

	palindrome := isPalindrome(&arr[0])
	fmt.Println(palindrome)

}

func printList(head *ListNode2) {
	fmt.Print("start-> ")
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, " -> ")
		cur = cur.Next
	}

	fmt.Println(" end ")

}
