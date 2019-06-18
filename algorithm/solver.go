package algorithm

import (
	"algorithms/greedy/mazeSolver/model"
	"container/heap"
)

func FindPath(start, end *model.Node) *model.Item {
	var queue model.PriorityQueue
	items := make(map[*model.Node]*model.Item)
	startItem := model.CreateItem(start)
	startItem.SetWeight(0)
	startItem.SetPreviousItem(nil)
	items[start] = startItem
	add2Queue(&queue, startItem, items)

	for len(queue) > 0 {
		currentItem := (heap.Pop(&queue)).(*model.Item)

		currentItem.SetVisited()
		currentNode := currentItem.GetNode()

		if currentNode == end {
			break
		}

		for i := 0; i < 4; i++ {

			var tmpItem *model.Item
			var exist bool
			if tmpItem, exist = items[currentNode.GetConnection(i)]; !exist {
				tmpItem = model.CreateItem(currentNode.GetConnection(i))
			}
			weight, prev := relaxAndUpdatePrevNode(currentItem, tmpItem, i)
			tmpItem.SetWeight(weight)
			tmpItem.SetPreviousItem(prev)
			add2Queue(&queue, tmpItem, items)

		}

	}
	return items[end]

}

func add2Queue(queue *model.PriorityQueue, item *model.Item, items map[*model.Node]*model.Item) {
	if item.GetNode() != nil && !item.IsVisited() {
		items[item.GetNode()] = item
		heap.Push(queue, item)

	}
}
func relaxAndUpdatePrevNode(currentItem, tmpItem *model.Item, i int) (float64, *model.Item) {
	if currentItem.GetNode().GetWeight(i)+currentItem.GetWeight() < tmpItem.GetWeight() {
		return currentItem.GetNode().GetWeight(i) + currentItem.GetWeight(), currentItem
	}
	return tmpItem.GetWeight(), tmpItem.GetPreviousItem()

}
