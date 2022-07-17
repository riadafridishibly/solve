package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	var curr *ListNode
	if l1.Val < l2.Val {
		head = l1
		curr = l1
		l1 = l1.Next
	} else {
		head = l2
		curr = l2
		l2 = l2.Next
	}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}
	return head
}

func mergeKLists(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return merge(lists[0], lists[1])
	}

	mid := len(lists) / 2

	return merge(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
}

func main() {
	mergeKLists(nil)
}
