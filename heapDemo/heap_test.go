package heapDemo

import (
	"container/heap"
	"fmt"
	"testing"
)

// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example demonstrates an integer heap built using the heap interface.

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func TestExample_intHeap(t *testing.T) {
	h := &IntHeap{3, 2, 6, 8, 1, 4, 0, 5, 9, 7}
	fmt.Println("Ori:", h)
	heap.Init(h)
	fmt.Println("After Init", h)
	//heap.Push(h, 3)
	//fmt.Println("After InitPush", h)
	fmt.Printf("maximum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	//Output:
	//minimum: 1
	//1 2 3 5
	//fmt.Println(CreateMaxHeap([]int(*h)))
	//fmt.Println(h)

}

// 堆排序的过程就是通过改变数组内元素的位置, 使其任意节点都满足 i >= 2i+1 和 2i+2 处的元素
// 则此时,返回的就是一个大根堆, 堆顶就是最大值.
func CreateMaxHeap(data []int) int {

	n := len(data)

	// 从最后一个带有子节点的根开始扫描,看其是否是一个大(小)根堆.如果是则返回循环,不是
	// 就调整元素,是这棵树成为一个大根堆
	// 最后一个带有子节点的根在数组中的位置可以根据完全二叉树的性质, 即 n/2-1 来确定.
	for i := n/2 - 1; i >= 0; i-- {
		parentIdx := i
		for {
			// 根据完全二叉树的性质, 计算当前根的左节点在数组中的索引
			left := 2*parentIdx + 1
			// 当左子节点的索引大于数组长度,则结束当前轮排序,意味着该节点无左子节点
			if left >= n || left < 1 { //left < 1 一般发生在 int 溢出后
				break
			}
			// 当构建大根堆时, 就用 max, 小根堆用 min
			// 先将 max 定为左节点在数组中的索引, 然后判断右节点
			max := left
			// 此处条件为,判断右子树的根是否存在, 若存在且比左根大,则 max = 该索引
			if right := left + 1; right < n && data[right] > data[left] {
				max = right
			}
			// 不论上述条件是否执行, 在拿到左右子树中较大的索引后, 判断它与其父节点的关系
			if data[parentIdx] > data[max] {
				break
			} else {
				data[parentIdx], data[max] = data[max], data[parentIdx]
				// 当父节点于子节点发生交换后,为保证其下面的树结构仍满足条件,因此
				// 父节点指针还得下沉到交换的地方,重新进行排查
				parentIdx = max
			}

		}
	}
	// 构建完成后,返回最大值
	return data[0]

}
