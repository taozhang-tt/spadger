package xheap

import "container/heap"

type PriorityQueue struct {
	priorityQueue
}

func NewPriorityQueue(cap int, min ...bool) *PriorityQueue {
	return &PriorityQueue{
		priorityQueue: priorityQueue{
			min: len(min) == 0 || min[0],
			pqe: make([]*priorityQueueElement, 0, cap),
		},
	}
}

func (pq *PriorityQueue) Len() int {
	return pq.priorityQueue.Len()
}

func (pq *PriorityQueue) Cap() int {
	return pq.priorityQueue.Cap()
}

func (pq *PriorityQueue) Empty() bool {
	return pq.Len() == 0
}

func (pq *PriorityQueue) Top() (value interface{}, score int) {
	v := pq.priorityQueue.pqe[0]
	return v.value, v.score
}

func (pq *PriorityQueue) Pop() (value interface{}, score int) {
	v := heap.Pop(&pq.priorityQueue).(*priorityQueueElement)
	return v.value, v.score
}

func (pq *PriorityQueue) Push(value interface{}, score int) {
	heap.Push(&pq.priorityQueue, &priorityQueueElement{
		value: value,
		score: score,
	})
}

type priorityQueue struct {
	min bool
	pqe []*priorityQueueElement
}

type priorityQueueElement struct {
	value interface{}
	score int
}

func (pq priorityQueue) Len() int {
	return len(pq.pqe)
}

func (pq priorityQueue) Cap() int {
	return cap(pq.pqe)
}

func (pq priorityQueue) Less(i, j int) bool {
	if pq.min {
		return pq.pqe[i].score < pq.pqe[j].score
	} else {
		return pq.pqe[j].score < pq.pqe[i].score
	}
}

func (pq priorityQueue) Swap(i, j int) {
	pq.pqe[i], pq.pqe[j] = pq.pqe[j], pq.pqe[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	pq.pqe = append(pq.pqe, x.(*priorityQueueElement))
}

func (pq *priorityQueue) Pop() interface{} {
	n := pq.Len()
	r := pq.pqe[n-1]
	pq.pqe = pq.pqe[0 : n-1]
	return r
}

type PriorityQueue64 struct {
	priorityQueue64
}

func NewPriorityQueue64(cap int, min ...bool) *PriorityQueue64 {
	return &PriorityQueue64{
		priorityQueue64: priorityQueue64{
			min: len(min) == 0 || min[0],
			pqe: make([]*priorityQueueElement64, 0, cap),
		},
	}
}

func (pq *PriorityQueue64) Len() int {
	return pq.priorityQueue64.Len()
}

func (pq *PriorityQueue64) Cap() int {
	return pq.priorityQueue64.Cap()
}

func (pq *PriorityQueue64) Empty() bool {
	return pq.Len() == 0
}

func (pq *PriorityQueue64) Top() (value interface{}, score int64) {
	v := pq.priorityQueue64.pqe[0]
	return v.value, v.score
}

func (pq *PriorityQueue64) Pop() (value interface{}, score int64) {
	v := heap.Pop(&pq.priorityQueue64).(*priorityQueueElement64)
	return v.value, v.score
}

func (pq *PriorityQueue64) Push(value interface{}, score int64) {
	heap.Push(&pq.priorityQueue64, &priorityQueueElement64{
		value: value,
		score: score,
	})
}

type priorityQueue64 struct {
	min bool
	pqe []*priorityQueueElement64
}

type priorityQueueElement64 struct {
	value interface{}
	score int64
}

func (pq priorityQueue64) Len() int {
	return len(pq.pqe)
}

func (pq priorityQueue64) Cap() int {
	return cap(pq.pqe)
}

func (pq priorityQueue64) Less(i, j int) bool {
	if pq.min {
		return pq.pqe[i].score < pq.pqe[j].score
	} else {
		return pq.pqe[j].score < pq.pqe[i].score
	}
}

func (pq priorityQueue64) Swap(i, j int) {
	pq.pqe[i], pq.pqe[j] = pq.pqe[j], pq.pqe[i]
}

func (pq *priorityQueue64) Push(x interface{}) {
	pq.pqe = append(pq.pqe, x.(*priorityQueueElement64))
}

func (pq *priorityQueue64) Pop() interface{} {
	n := pq.Len()
	r := pq.pqe[n-1]
	pq.pqe = pq.pqe[0 : n-1]
	return r
}
