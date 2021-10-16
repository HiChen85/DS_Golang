package sortAlgorithm

import "fmt"

// 选择排序的一个核心点就是在每次循环遍历时,找到当前列表中的最小值
// 每找到一个最小值,就将其记录,与当前对比的元素进行替换.
// 由于给定的数据中可能无法将每次搜寻的最小值 pop,因此, 还需要进行替换
// 根据 golang 的特性,在每次找到最小值后, 就可以利用并行赋值的规则来
// 交换档次循环的头元素与找到的最小值

func SelectSort(data []int) {
	for i := range data {
		smallest := data[i]
		smallIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < smallest {
				smallest = data[j]
				smallIndex = j
			}
		}
		// 在每次循环找到最小值之后, 将最小值和当前循环比较的开头数据交换,保证每次
		// 循环都将最小值放在列表的头部
		data[i], data[smallIndex] = smallest, data[i]
	}
}

// Sum1 利用递归求和
func Sum1(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	} else {
		return data[0] + Sum1(data[1:])
	}
}

// CountNum 利用递归计算切片中的元素个数
func CountNum(data []int) int {
	num := 0
	if len(data) == 0 {
		return 0
	} else {
		num++
		return num + CountNum(data[1:])
	}
}

// FindMax 利用递归找到切片中的最大值
// 此函数的核心思想实际是从最后两个数据开始,
// 比较后得到一个较大的值返回,然后再跟前一个比较.
// 这个倒序的过程其实是因为递归的特性决定的
func FindMax(data []int) int {
	if len(data) == 1 {
		fmt.Println("递归已到最底端,数据为", data)
		return data[0]
	} else {
		max := data[0]
		rst := FindMax(data[1:])
		fmt.Println("当前待比较数据", max, rst)
		if max < rst {
			return rst
		} else {
			return max
		}
	}
}
