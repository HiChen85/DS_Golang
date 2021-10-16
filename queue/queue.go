package queue

import "fmt"

type Queue struct {
	data  []interface{}
	front int
	rear  int
	size  int
}

func NewQueue(n int) *Queue {
	q := new(Queue)
	q.data = make([]interface{}, n+1)
	q.front = 0
	q.rear = 0
	q.size = len(q.data) - 1
	return q
}

func (q *Queue) PrintQueue() {
	if !q.IsEmpty() {
		fmt.Printf("[ ")
		for i := q.front; (i)%len(q.data) != q.rear; i++ {
			fmt.Printf("%v ", q.data[i])
		}
		fmt.Printf("]\n")
	}
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsFull() bool {
	if (q.rear+1)%len(q.data) == q.front {
		return true
	} else {
		return false
	}
}

func (q *Queue) IsEmpty() bool {
	if q.rear == q.front {
		return true
	} else {
		return false
	}
}

func (q *Queue) EnQueue(e interface{}) {

	if q.IsFull() {
		fmt.Println("队列已满")
		return
	} else {
		q.data[q.rear] = e
		// 循环队列中添加元素后, 队尾指针移动的方向是 (r + 1) % 队列的数组长度
		q.rear = (q.rear + 1) % len(q.data)
	}
}

func (q *Queue) DeQueue() interface{} {
	if q.IsEmpty() {
		return 0
	} else {
		value := q.data[q.front]
		q.data[q.front] = 0
		// 出队操作完成后, 对头指针也需要按照 (f + 1) % 队列的数组长度 移动
		q.front = (q.front + 1) % len(q.data)
		return value
	}
}

//func TestQueue() {
//	d := []int{2, 13, 4, 5, 6}
//	q := NewQueue(len(d))
//
//	for i := 0; i < len(d); i++ {
//		q.EnQueue(d[i])
//	}
//	if q.IsFull() {
//		fmt.Println("队列已装满")
//		fmt.Println("队列长度为", q.Size())
//	}
//	q.PrintQueue()
//	for i := 0; i < len(d); i++ {
//		fmt.Println("出队元素:", q.DeQueue())
//	}
//	if q.IsEmpty() {
//		fmt.Println("队列已清空")
//	}
//
//	fmt.Println("火车车厢重排")
//	train := []int{5, 8, 1, 6, 9, 7, 3, 2, 4}
//	q1 := SortTrain(train)
//	fmt.Println("重排后:")
//	q1.PrintQueue()
//
//}

//func SortTrain(train []int) *Queue {
//
//	out := NewQueue(len(train))
//
//	smallestBox := findSmallestBox(train)
//	fmt.Println("最小车厢号:", smallestBox)
//
//	tracks := make([]*Queue, 3)
//	for i := 0; i < len(tracks); i++ {
//		tracks[i] = NewQueue(len(train))
//	}
//
//	for i := len(train) - 1; i >= 0; i-- {
//		fmt.Println("当前车厢:", train[i])
//
//		bestTrack := -1
//		bestLast := 0 // 取最佳缓冲轨道最后的车厢
//
//		for j := range tracks {
//			if !tracks[j].IsEmpty() {
//				cBox := tracks[j].data[tracks[j].rear-1] // 拿到当前轨道中的车厢号
//				if train[i] > cBox && cBox > bestLast {
//					bestLast = cBox
//					bestTrack = j
//				}
//			} else {
//				if bestTrack == -1 {
//					bestTrack = j
//				}
//			}
//		}
//
//		if bestTrack == -1 {
//			fmt.Println("无可用轨道, 当前车厢无法完成排序")
//		} else {
//			tracks[bestTrack].EnQueue(train[i])
//		}
//
//		fmt.Println("出队前:")
//		for i := range tracks {
//			tracks[i].PrintQueue()
//		}
//
//		if train[i] == smallestBox {
//			out.EnQueue(tracks[bestTrack].DeQueue())
//		}
//
//		smallerTracks := findMin(tracks)
//		fmt.Println("<<", smallerTracks, ">>")
//		for smallerTracks != -1 && !out.IsEmpty() {
//			if tracks[smallerTracks].data[tracks[smallerTracks].front] > out.data[out.rear-1] {
//				out.EnQueue(tracks[smallerTracks].DeQueue())
//				smallerTracks = findMin(tracks)
//			} else {
//				break
//			}
//
//		}
//		fmt.Println("Out:")
//		out.PrintQueue()
//		fmt.Println("<<<<<<<<<<<<<<<<<<")
//
//	}
//
//	return out
//
//}
//
//func findSmallestBox(train []int) int {
//	min := train[0]
//	for i := range train {
//		if train[i] < min {
//			min = train[i]
//		}
//	}
//
//	return min
//}
//
//func findMin(tracks []*Queue) int {
//	min := 0
//	minIndex := -1
//	for i := range tracks {
//		if !tracks[i].IsEmpty() && min == 0 {
//			min = tracks[i].data[tracks[i].front]
//			minIndex = i
//		}
//
//		if !tracks[i].IsEmpty() && tracks[i].data[tracks[i].front] < min {
//			min = tracks[i].data[tracks[i].front]
//			minIndex = i
//		}
//	}
//	return minIndex
//}
