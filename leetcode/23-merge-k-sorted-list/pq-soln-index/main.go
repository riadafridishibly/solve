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

type Item struct {
	root  *ListNode
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].root.Val < pq[j].root.Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]

	// Modification
	ret := item.root

	if item.root.Next == nil {
		old[n-1] = nil  // avoid memory leak
		item.index = -1 // for safety
		*pq = old[0 : n-1]
	} else {
		item.root = item.root.Next
		pq.update(item)
	}

	return ret
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item) {
	heap.Fix(pq, item.index)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	pq := make(PriorityQueue, 0, len(lists))
	for i, lis := range lists {
		if lis != nil {
			pq = append(pq, &Item{
				root:  lis,
				index: i,
			})
		}
	}

	heap.Init(&pq)
	out := &ListNode{}
	head := out
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		out.Next = node
		out = out.Next
	}

	return head.Next
}

func main() {
	mergeKLists([]*ListNode{nil, nil, nil})
}
