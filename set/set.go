package set

import "fmt"

type Set map[interface{}]struct{}

func NewSet() *Set {
	set := new(Set)
	*set = make(map[interface{}]struct{})
	return set
}

func (s *Set) String() string {
	var str = "{"
	for k, _ := range *s {
		str += fmt.Sprintf("%v, ", k)
	}
	str += "}"
	return str
}

func (s *Set) Add(n interface{}) {
	if s.Has(n) {
		fmt.Println(n, "已存在, 不能添加相同元素到集合中")
	} else {
		(*s)[n] = struct{}{}
	}
}

func (s *Set) Remove(n interface{}) {
	delete(*s, n)
}

func (s *Set) Has(n interface{}) bool {
	if _, ok := (*s)[n]; ok {
		return true
	} else {
		return false
	}
}

func (s *Set) IsEmpty() bool {
	return 0 == len(*s)
}

func (s *Set) Len() int {
	return len(*s)
}

func (s *Set) toSlice() []interface{} {
	var sli []interface{}
	for k, _ := range *s {
		sli = append(sli, k)
	}
	return sli
}
