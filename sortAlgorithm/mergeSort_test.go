package sortAlgorithm

import (
	"fmt"
	"testing"
)

func MergeSort(data []int) []int {
	// 归并排序本身也是一个分支思想的排序,因此需要一些递归操作, 首先定义递归头
	dataLength := len(data)

	if dataLength < 2 {
		return data
	}

	middle := dataLength >> 1

	left := data[:middle]
	right := data[middle:]

	return merge(MergeSort(left), MergeSort(right))

}

func merge(left, right []int) []int {
	// 对于归并排序,最重要的就是对分隔开的子数组进行排序后合并,层层递归
	// 对于排序结果需要另外的空间来存储
	var merged []int

	lLength := len(left)
	rLength := len(right)

	// 维护两个分别指向分好块的待排序区间, 起初, 对待排序区间的头元素进行比较, 将较小的元素放入合并空间
	// 直到其中一方完成,这个过程一旦完成,剩余的未排序的元素一定比前面排序的元素都 大
	// 直接将其添加在待排序空间的后面即可
	i, j := 0, 0

	for i < lLength && j < rLength {
		if left[i] <= right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	if i != lLength {
		merged = append(merged, left[i:]...)
	}
	if j != rLength {
		merged = append(merged, right[j:]...)
	}

	return merged

}

func TestMergeSort(t *testing.T) {
	data := []int{3, 1, 4, 5, 9, 0, 2}
	wanted := []int{0, 1, 2, 3, 4, 5, 9}

	got := MergeSort(data)

	if !IsSliceEqual(got, wanted) {
		t.Errorf("got %v but wanted %v", got, wanted)
	}

}
func IsSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func countCharacters(words []string, chars string) int {
	// 利用 map 将单词中包含的字母与单词对应

	cMap := make(map[string]int, len(chars))
	rst := 0

	for _, c := range chars {
		cMap[(string)(c)]++
	}

	fmt.Println("cMap", cMap)
	for _, w := range words {
		wMap := make(map[string]int)
		for _, c := range w {
			wMap[(string)(c)]++
		}
		fmt.Println("wMap:", wMap)
		for k, v := range wMap {
			if count, ok := cMap[(string)(k)]; ok && v <= count {
				rst += v
			} else {
				break
			}
		}

	}

	return rst

}

func TestCountChar(t *testing.T) {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"

	wanted := 6

	got := countCharacters(words, chars)

	if wanted != got {
		t.Errorf("got %v but wanted %v", got, wanted)
	}

}
