package sortAlgorithm

import "fmt"

// QuickSort 快速排序
// 快速排序使用了分而治之的思想, 在选取了基准值的基础上,将整个数组划分为两部分
// 而每一部分又可以看做是一个新的待排序数组,从而进行新一轮的快速排序,因此使用
// 递归. 需要注意的是确定递归结束的条件. 快速排序实在原有数组上的排序.
func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}

	left, right := 0, len(data)-1

	pivot := partition(data, left, right)
	fmt.Println("当前 pivot", data[pivot])

	QuickSort(data[:pivot])
	QuickSort(data[pivot+1:])

}

func partition(data []int, left, right int) int {
	pivot := left
	for left < right {
		if data[right] > data[pivot] {
			right--
			continue
		}
		if data[left] <= data[pivot] {
			left++
			continue
		}
		data[left], data[right] = data[right], data[left]
	}
	data[left], data[pivot] = data[pivot], data[left]
	return left

}

// QuickSort2 三向切分快速排序
// 利用左右双移动指针, 将与 pivot 相同的元素全部集中于
// 指定的动态双指针的中间,从而减少递归的次数, 提高快速排序的性能
func QuickSort2(data []int) {
	if len(data) <= 1 {
		return
	}

	//lt, gt := 0, len(data)-1
	//
	////当 pivot 直接赋值为数据元素时,更换位置并不会改变 pivot 的值
	//pivot := data[0]
	//
	//for i := lt + 1; i <= gt; {
	//	if data[i] < pivot {
	//		data[i], data[lt] = data[lt], data[i]
	//		lt++
	//		i++
	//	} else if data[i] > pivot {
	//		data[gt], data[i] = data[i], data[gt]
	//		gt--
	//	} else {
	//		i++
	//	}
	//}
	lt, gt := partition2(data)
	QuickSort2(data[:lt])
	QuickSort2(data[gt+1:])
}

func partition2(data []int) (int, int) {
	lt, gt := 0, len(data)-1
	pivot := data[lt]
	for i := lt + 1; i <= gt; {
		if data[i] < pivot {
			data[i], data[lt] = data[lt], data[i]
			lt++
			i++
		} else if data[i] > pivot {
			data[i], data[gt] = data[gt], data[i]
			gt--
		} else {
			i++
		}
	}
	return lt, gt
}
