package stack

import "fmt"

type stack struct {
	data   []interface{}
	top    int
	bottom int
}

func NewStack() *stack {
	fmt.Println("建立一个空栈...")
	s := new(stack)
	s.top = -1
	s.bottom = s.top
	return s
}

func (s *stack) PushBack(e interface{}) {
	s.data = append(s.data, e)
	s.top++
}

func (s *stack) PopStack() interface{} {
	if s.top >= 0 {
		value := s.data[s.top]
		s.data = s.data[:s.top]
		s.top--
		return value
	} else {
		s.top = -1
		//fmt.Println("已到栈底")
		return nil
	}
}

func (s *stack) IsEmpty() bool {
	if s.top == s.bottom {
		return true
	} else {
		return false
	}
}

func (s *stack) Top() interface{} {
	if !s.IsEmpty() {
		return s.data[s.top]
	} else {
		return nil
	}
}

func TestNewStack() {
	s := NewStack()
	fmt.Println(s.data == nil) //结果为 true
	fmt.Println(s.top, s.bottom)

	for i := 0; i < 5; i++ {
		s.PushBack(i + 1)
		fmt.Println("当前栈内元素", s.data)
	}
	for i := 0; i < 6; i++ {
		fmt.Println("当前栈顶元素", s.PopStack())
		fmt.Println("当前栈内元素", s.data, len(s.data))
	}

}

// 利用栈来完成一个bracket matching

func BracketMatch(str string) []string {
	s := NewStack()
	var rst []string

	for i := 0; i < len(str); i++ {
		t := (string)(str[i])
		if t == "(" {
			s.PushBack(t)
		} else if t == ")" {
			if !s.IsEmpty() {
				c := s.PopStack()
				if v, ok := c.(string); ok {
					rst = append(rst, v, t)
				}
			} else {
				rst = append(rst, t)
			}
		}
	}
	for !s.IsEmpty() {
		e := s.PopStack()
		if v, ok := e.(string); ok {
			rst = append(rst, v)
		}
	}
	return rst
}

func TestBracketMatch() {
	data := "(a+b)*c+())()))"
	fmt.Println(BracketMatch(data))
}

func hanota(n int, A, B, C int) {
	t := make(map[int]*stack)
	t[A] = NewStack()
	for i := n - 1; i >= 0; i-- {
		t[A].PushBack(i)
	}
	t[B] = NewStack()
	t[C] = NewStack()
	fmt.Println("A 塔:")
	for i := len(t[A].data) - 1; i >= 0; i-- {
		fmt.Println(t[A].data[i])
	}
	move(len(t[A].data), t[A], t[B], t[C])
	fmt.Println("移动后...")
	fmt.Printf("A塔:%v \n", t[A].data)
	fmt.Println("C 塔:")
	for i := len(t[C].data) - 1; i >= 0; i-- {
		fmt.Println(t[C].data[i])
	}
}

func move(n int, A, B, C *stack) {
	if n > 0 {
		move(n-1, A, C, B)
		d := A.Top()
		A.PopStack()
		C.PushBack(d)
		move(n-1, B, A, C)
	}
}

func hanota2(A []int, B []int, C []int) []int {
	if len(A) > 0 {
		B = hanota2(A[1:], C, B)
		fmt.Println("B长度:", len(B))
		C = append(C, A[0])
		A = A[:len(A)-1]
		C = hanota2(B, A, C)
	}
	return C
}

func TestHanota() {
	//hanota(3, 0, 1, 2)
	A := []int{1, 0}
	var B, C []int
	hanota2(A, B, C)
}
