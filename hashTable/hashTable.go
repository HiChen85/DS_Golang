package hashTable

import (
	"fmt"
	"unsafe"
)

// 哈希表, map 的应用
type Map interface {
	AddItem(k string, v interface{}) bool
}

// 这样的定义方式表示说 HashTable 是一个带有后面类型特性的新类型,这两个类型在类型系统中是不同的类型
//
type HashTable map[string]interface{}

func NewHashTable() HashTable {
	return make(HashTable)
}

func (h HashTable) AddItem(k string, v interface{}) bool {
	h[k] = v
	if _, ok := h[k]; ok {
		fmt.Println("完成添加")
		return true
	} else {
		return false
	}
}

func (h HashTable) GetValue(key string) interface{} {
	if v, ok := h[key]; ok {
		return v
	} else {
		return nil
	}

}

func TestInterface(m Map) {
	fmt.Printf("%#p\n", m)
	fmt.Printf("%v\n", m)
}

type Cache = HashTable

func Test() {
	//m1 := NewHashTable()
	//TestInterface(m1)
	m2 := make(map[string]interface{}) // 等价与 make(map[string]interface{})
	fmt.Printf("%#p\n", m2)
	m2["123"] = 123
	fmt.Printf("m2 :%#v\n", m2)
	m3 := *(*HashTable)(unsafe.Pointer(&m2))
	TestInterface(m3)

	cache := NewHashTable()
	url := "https://www.baidu.com"
	cache.AddItem(url, "百度页面")
	fmt.Println(cache.GetValue(url))
	fmt.Println(cache[url])

}
