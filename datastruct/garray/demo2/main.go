package main

import (
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
)

//garray -- 排序数组
func main()  {
	a := garray.NewSortedArray(func(v1, v2 interface{}) int {
		if v1.(int) < v2.(int) {
			return 1
		}
		if v1.(int) > v2.(int) {
			return -1
		}
		return 0
	})

	// 添加数据
	a.Add(2)
	a.Add(3)
	a.Add(1)
	//Slice 说明：得到数组的slice切片数据，注意，如果是在并发安全模式下，返回的是一份拷贝数据，否则返回的是指向数据的指针
	fmt.Println(a.Slice())

	// 添加重复数据
	a.Add(3)
	fmt.Println(a.Slice())

	// 检索数据，返回最后对比的索引位置，检索结果
	// 检索结果：0: 匹配; <0:参数小于对比值; >0:参数大于对比值
	fmt.Println(a.Search(1))

	// 设置不可重复
	a.SetUnique(true)
	fmt.Println(a.Slice())
	a.Add(1)
	fmt.Println(a.Slice())
}