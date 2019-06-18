package model

import (
	"math"
)

type Item struct {
	start  *Node
	prev   *Item
	weight float64
	index  int
}

func (item Item) IsVisited() bool {
	return item.start.Visited
}
func (item *Item) SetPreviousItem(n *Item) {
	item.prev = n
}
func (item Item) GetPreviousItem() *Item {
	return item.prev
}
func (item *Item) SetVisited() {
	item.start.Visited = true
}
func (item Item) GetWeight() float64 {
	return item.weight
}
func (item Item) GetNode() *Node {
	return item.start
}

func CreateItem(n *Node) *Item {

	return &Item{start: n, weight: math.MaxFloat64}
}
func (item *Item) SetWeight(a float64) {
	item.weight = a
}

type PriorityQueue []*Item

func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)

}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func CreateQueue(length int) PriorityQueue {
	return make(PriorityQueue, length)
}
