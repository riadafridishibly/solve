package main

import (
	"container/heap"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type any = interface{}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	pq := make(PriorityQueue, 0, len(lists))
	for _, lis := range lists {
		if lis != nil {
			pq = append(pq, lis)
		}
	}

	heap.Init(&pq)
	out := &ListNode{}
	head := out
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		out.Next = node
		out = out.Next
		if node != nil && node.Next != nil {
			heap.Push(&pq, node.Next)
		}
	}

	return head.Next
}

func main() {
	mergeKLists([]*ListNode{nil, nil, nil})
}
