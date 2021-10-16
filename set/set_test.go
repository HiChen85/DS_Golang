package set

import (
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet()
	set.Add("你好")
	t.Log(set)
	set.Add("你好")
}
