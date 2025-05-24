package taskRewards

import (
	"container/heap"
	"sort"
	"strconv"
	"strings"
)

type task struct {
	deadline int
	reward   int
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func MaxTotalReward(n int, dw string) int {
	taskLines := strings.Split(strings.TrimSpace(dw), "\n")
	taskList := make([]task, n)
	for i := 0; i < n; i++ {
		parts := strings.Fields(taskLines[i])
		d, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		taskList[i] = task{deadline: d, reward: w}
	}

	sort.Slice(taskList, func(i, j int) bool {
		if taskList[i].deadline == taskList[j].deadline {
			return taskList[i].reward > taskList[j].reward
		}
		return taskList[i].deadline < taskList[j].deadline
	})

	h := &minHeap{}
	heap.Init(h)

	for _, task := range taskList {
		heap.Push(h, task.reward)
		if h.Len() > task.deadline {
			heap.Pop(h)
		}
	}

	sum := 0
	for _, val := range *h {
		sum += val
	}
	return sum
}
