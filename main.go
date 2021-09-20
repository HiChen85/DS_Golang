package main

// 使用 go mod 方法导入自定义包时, 需要使用 go mod init 时
// 输入的项目名称, 这样编译器才能从当前目录寻找对应的包
import (
	dl "data_structure/doubleLinkedList"
)

func main() {
	l := dl.New()
	arr := []string{"张三", "李四", "王五", "张麻子", "黄四郎"}
	//for i := 0; i < len(arr); i++ {
	//	l.Append(arr[i])
	//}
	//l.PrintValue()
	//fmt.Println("------")
	//fmt.Println(l.Len())
	//fmt.Println("------")
	//l.Insert("汤师爷", "王五")
	//l.PrintValue()
	//fmt.Println("------")
	//fmt.Println(l.Len())
	//fmt.Println("------")
	//l.Delete("汤师爷")
	//l.PrintValue()

	for _, v := range arr {
		l.Append(v)
	}
	l.PrintForward()
	l.PrintBack()

}
