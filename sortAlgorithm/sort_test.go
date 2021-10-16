package sortAlgorithm

import (
	"math/rand"
	"testing"
	"time"
)

func TestFastSort(t *testing.T) {

	nums1 := []int{}
	nums2 := []int{1}
	nums3 := []int{2, 1, 5}
	nums4 := []int{2, 1, 5, 2, 5, 6, 9, 0, 1, 3}

	FastSort(nums1)
	FastSort(nums2)
	FastSort(nums3)
	FastSort(nums4)

	t.Log(nums1)
	t.Log(nums2)
	t.Log(nums3)
	t.Log(nums4)

	nums5 := []int{4, 4, 2, 1, 4, 5, 6, 0, 9, 8, 7, 8, 2, 3, 6, 1, 0}

	sortColors(nums5)
	t.Log(nums5)

	nums6 := []int{5, 21, 4, 6, 753, 78, 3, 546, 12, 0, 9, 76}
	nums6 = MergeSorts(nums6)
	t.Log(nums6)

}

func FastSort(nums []int) {

	if len(nums) < 2 {
		return
	}

	//快速排序的核心是利用比较交换,找到一个基准值,将数组分成三部分,左边的都小于基准值,基准值和右边的大于基准值的三段区间.
	// 利用 partition 只能找到一次基准值,有了基准值,通过切片操作,就能很轻松的获取它两侧的数组,进入递归缩小规模.
	pivot := partitions(nums, 0, len(nums)-1)
	FastSort(nums[:pivot])
	FastSort(nums[pivot+1:])
	return
}

func partitions(nums []int, left, right int) int {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(nums))

	nums[0], nums[index] = nums[index], nums[0]

	pivot := left

	for left < right {
		if nums[right] >= nums[pivot] {
			right--
			continue
		}
		if nums[left] <= nums[pivot] {
			left++
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[pivot], nums[left] = nums[left], nums[pivot]
	return left
}

// 反向切分快速排序
func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	// 三向切分的快速排序原则是利用两个额外指针,来划分大于等于和小于基准值的三块区间,三块区间的大小和位置会因为选取的基准值不同而有所差别.
	// lt 指向数组的第二个元素是因为第一个元素需要被指定为基准值用于比较.但是一般可以选取数组中随机一个值然后与头值进行交换.
	lt, gt := 0, len(nums)-1

	// 这里使用 nums[0] 而不是直接用索引,是因为使用索引的时候,如果头元素被交换,则无法保证当前的头元素仍然是一开始选定的 pivot
	pivot := nums[0]

	// i<= gt 是因为 gt 是最后一个待判断的元素,只有 gt 也被判断过,才算结束
	// i的自增只在 i 指定的元素元素与 pivot 相等或者 lt < i 时才自增.
	for i := 1; i <= gt; {
		if nums[i] < pivot {
			nums[i], nums[lt] = nums[lt], nums[i]
			lt++
			i++
		} else if nums[i] > pivot {
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		} else {
			i++
		}
	}
	// 结束上述判断后,会获得 lt 与 gt 两个值用来切分待排序数组
	sortColors(nums[:lt])
	sortColors(nums[gt+1:])
}

// 归并排序的核心就是先分, 利用递归无限制的分割数组,直到只剩一个元素为止,然后将两个数组合并,排序的过程只发生在合并期间.因此交归并或者合并排序
func MergeSorts(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	middle := len(nums) / 2
	left := nums[:middle]
	right := nums[middle:]

	nums = merges(MergeSorts(left), MergeSorts(right))
	return nums

}

func merges(left, right []int) []int {
	var rst []int

	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			rst = append(rst, left[0])
			left = left[1:]
		} else {
			rst = append(rst, right[0])
			right = right[1:]
		}
	}

	if len(left) != 0 {
		rst = append(rst, left...)
	}
	if len(right) != 0 {
		rst = append(rst, right...)
	}

	return rst

}
