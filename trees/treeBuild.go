package trees

import "fmt"

func BuildTree(preOrder, inOrder []int) *TreeNode {
	inPos := make(map[interface{}]int)
	for i, v := range inOrder {
		inPos[v] = i
	}
	return CreateTreeWithPreIn(preOrder, 0, len(preOrder)-1, 0, inPos)
}

func CreateTreeWithPreIn(pre []int, pStart, pEnd int, inStart int, inPos map[interface{}]int) *TreeNode {
	if pStart > pEnd {
		return nil
	}

	// 第一步找根
	root := new(TreeNode)
	root.data = pre[pStart]
	rootIndex := inPos[root.data]

	// 计算根左边的元素个数, 也是根在先序遍历中偏移多少个后,得到右子树的长度.
	leftLen := rootIndex - inStart

	root.left = CreateTreeWithPreIn(pre, pStart+1, pStart+leftLen, inStart, inPos)
	root.right = CreateTreeWithPreIn(pre, pStart+leftLen+1, pEnd, rootIndex+1, inPos)
	return root
}

func CreateTree() {
	preOrder := []int{1, 2, 3, 4, 5, 6, 7, 8}
	inOrder := []int{3, 4, 2, 1, 7, 6, 5, 8}
	root := BuildTree(preOrder, inOrder)
	InOrder(root)
	fmt.Println()
	PostOrder(root)

}
