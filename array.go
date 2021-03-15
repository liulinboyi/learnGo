package main

import (
	"fmt"
	"strconv"
)

func unique(arr []int) []int {
	// hash 去重
	hash := make(map[string]int)
	a := arr
	res := make([]int, 0, len(a))
	for i := 0; i < len(a); i++ {
		index := strconv.Itoa(a[i])
		value := hash[index]
		if value != 1 {
			hash[index] = 1
			res = append(res, a[i])
		}
	}
	return res
}

func main() {
	{
		var arr [5]int     // 一维
		var arr2 [5][5]int // 二维

		fmt.Println(arr)
		fmt.Println(arr2)
		fmt.Println("---------------------------")
	}
	{
		var arr = [5]int{1, 2, 3, 4, 5}
		// 或 arr := [5]int{1, 2, 3, 4, 5}
		fmt.Println(arr)
		fmt.Println("---------------------------")

	}
	{
		arr := [5]int{1, 2, 3, 4, 5}
		for i := 0; i < len(arr); i++ {
			arr[i] += 100
		}
		fmt.Println(arr) // [101 102 103 104 105]
		fmt.Println("---------------------------")

	}
	{
		//数组的长度不能改变，如果想拼接2个数组，或是获取子数组，需要使用切片。切片是数组的抽象。 切片使用数组作为底层结构。切片包含三个组件：容量，长度和指向底层数组的指针,切片可以随时进行扩展
		slice1 := make([]float32, 0)          // 长度为0的切片
		slice2 := make([]float32, 3, 5)       // [0 0 0] 长度为3容量为5的切片
		fmt.Println(len(slice2), cap(slice2)) // 3 5
		fmt.Println(slice1)
		fmt.Println("---------------------------")

	}
	{
		slice2 := make([]float32, 3, 5) // [0 0 0] 长度为3容量为5的切片
		// 添加元素，切片容量可以根据需要自动扩展
		slice2 = append(slice2, 1, 2, 3, 4)   // [0, 0, 0, 1, 2, 3, 4]
		fmt.Println(len(slice2), cap(slice2)) // 7 12
		// 子切片 [start, end)
		sub1 := slice2[3:]  // [1 2 3 4]
		sub2 := slice2[:3]  // [0 0 0]
		sub3 := slice2[1:4] // [0 0 1]
		// 合并切片
		combined := append(sub1, sub2...) // [1, 2, 3, 4, 0, 0, 0]
		fmt.Println(sub3)
		fmt.Println(combined)
		fmt.Println("---------------------------")

	}

	{
		a := []int{1, 2, 3, 4, 5, 1, 2, 4, 7}
		res := unique(a)
		fmt.Println("去重之前：", a)
		fmt.Println("去重之后：", res)
	}

}
