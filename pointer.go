package main

import "fmt"

func add(num int) {
	num += 1
}

func realAdd(num *int) {
	*num += 1
}

func main() {
	{
		//指针即某个值的地址，类型定义时使用符号*，对一个已经存在的变量，使用 & 获取该变量的地址。
		str := "Golang"
		var p *string = &str // p 是指向 str 的指针
		*p = "Hello"
		fmt.Println(str) // Hello 修改了 p，str 的值也发生了改变
	}
	{
		// 一般来说，指针通常在函数传递参数，或者给某个类型定义新的方法时使用。Go 语言中，参数是按值传递的，如果不使用指针，函数内部将会拷贝一份参数的副本，对参数的修改并不会影响到外部变量的值。如果参数使用指针，对参数的传递将会影响到外部变量。
		num := 100
		add(num)
		fmt.Println(num) // 100，num 没有变化

		realAdd(&num)
		fmt.Println(num) // 101，指针传递，num 被修改
	}

}
