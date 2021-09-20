package binarySearch

func BinarySearch(data []int, target int) int {
	// left, right 都是data 的索引值
	left, right := 0, len(data)-1
	// 二分查找的一个重点就是在左右索引相等时也要进行一次判断.
	// 在左右索引相等时, 即是循环结束的时候, 也是返回当元素不在排序列表中时
	// 需要插入的位置
	for left <= right {
		// 二分查找的核心,就是计算中点值, 每次判断中点值和 target 的差距来计算 target
		// 落在左半区间还是右半区间, 以便左右端点的移动
		// mid 也是索引值,因此在计算时, 需要计算相对的中点值再加上当前区间左端点的索引
		// 才是mid 的绝对索引值
		// 右移比除法更快
		mid := left + (right-left)>>1
		if data[mid] == target {
			return mid
		} else if target < data[mid] { // 当目标值小于 mid 处的值时, 意味着目标值落在左边, 右端点移动至 mid 的前一位
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
