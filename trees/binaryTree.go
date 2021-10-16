package trees

import (
	"data_structure/queue"
	"fmt"
)

type TreeNode struct {
	data  interface{}
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) SetValue(value interface{}) {
	t.data = value
}

func (t *TreeNode) GetValue() interface{} {
	return t.data
}

func (t *TreeNode) IsEmpty() bool {
	if t == nil {
		return true
	} else {
		return false
	}
}

type BinaryTree struct {
	root *TreeNode
	size int
}

// NewBinaryTree 创建一颗空树,没有任何节点
func NewBinaryTree() *BinaryTree {
	tree := new(BinaryTree)
	return tree
}

func (b *BinaryTree) CreateBinaryTree(data ...interface{}) {

	q := queue.NewQueue(len(data))

	b.root = new(TreeNode)

	q.EnQueue(b.root)

	for i := 0; !q.IsEmpty(); i++ {
		e := q.DeQueue()
		if v, ok := e.(*TreeNode); ok && i < len(data) {
			v.data = data[i]
			b.size++
			// 在完全二叉树中, 数元素的个数等于其最后一个节点的编号. 节点的编号是按层序分配的.
			// 根据完全二叉树的性质,对于编号为 i 的节点,若 2 * i > n, 则其无左孩子
			// 当 2 * i + 1 > n 时,其无右孩子.
			if 2*b.size <= len(data) {
				v.left = new(TreeNode)
				q.EnQueue(v.left)
			}
			if 2*b.size+1 <= len(data) {
				v.right = new(TreeNode)
				q.EnQueue(v.right)
			}
		}
	}
}

func (b *BinaryTree) CreateTree(data []int) {
	// 可以使用 channel 作为队列使用, 当作为队列时, channel 需要提前分配好空间,也就是说要创建一个有缓冲的 channel
	q := make(chan *TreeNode, 2)
	// 当创建好 channel 后,如果不使用 for range 循环去遍历 channel 来取出 channel 中的数据, 那么就需要手动使用 defer
	// 去注册 channel 的关闭操作
	defer close(q)

	b.root = new(TreeNode)

	q <- b.root

	for i := 0; i < len(data); i++ {
		if v, ok := <-q; ok {
			v.data = data[i]
			b.size++

			if 2*b.size <= len(data) {
				v.left = new(TreeNode)
				q <- v.left
			}

			if 2*b.size+1 < len(data) {
				v.right = new(TreeNode)
				q <- v.right
			}
		}
	}

}

func PreOrder(node *TreeNode) {
	if !node.IsEmpty() {
		fmt.Printf("%v ", node.GetValue())
		PreOrder(node.left)
		PreOrder(node.right)
	}
}

func InOrder(node *TreeNode) {
	if !node.IsEmpty() {
		InOrder(node.left)
		fmt.Printf("%v ", node.GetValue())
		InOrder(node.right)
	}
}

func PostOrder(node *TreeNode) {
	if !node.IsEmpty() {
		PostOrder(node.left)
		PostOrder(node.right)
		fmt.Printf("%v ", node.GetValue())
	}
}

func LevelOrder(node *TreeNode) {
	q := make(chan *TreeNode, 2)
	defer close(q)
	q <- node
	for len(q) != 0 {
		if v, ok := <-q; ok {
			fmt.Printf("%v ", v.data)
			if v.left != nil {
				q <- v.left
			}
			if v.right != nil {
				q <- v.right
			}
		}
	}

}

func TestBinaryTree() {
	tree := NewBinaryTree()
	fmt.Println("测试 tree 与 tree的根节点是否为空:", tree == nil, tree.root.IsEmpty())
	fmt.Println("测试 tree 的个数在创建根节点后是否为 1(实则建立空树, size应为 0)", tree.size)
	tree.CreateTree([]int{1, 2, 3, 4})
	fmt.Println("先序遍历:")
	PreOrder(tree.root)
	fmt.Println()
	fmt.Println("中序遍历:")
	InOrder(tree.root)
	fmt.Println()
	fmt.Println("层序遍历:")
	LevelOrder(tree.root)

}
