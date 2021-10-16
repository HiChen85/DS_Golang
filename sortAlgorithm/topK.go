package sortAlgorithm

// 求数组中第 n 大的数
// 一个数组中第 n 大的数, 就是在从小到大排序后,索引为 len() - n 的数

func FindKNumber(nums []int, k int) int {
	target := len(nums) - k

	left, right := 0, len(nums)-1

	index := partition(nums, left, right)

	for target != index {
		if target < index {
			index = partition(nums, left, index)
		} else if target > index {
			index = partition(nums, index+1, right)
		}
	}
	return nums[index]
}
