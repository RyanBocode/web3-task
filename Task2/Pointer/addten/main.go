/*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*/

package main

import "fmt"

// 定义函数，接收一个整数指针参数
func addTen(num *int) {
	*num += 10 // 通过指针修改实际值
}

func main() {
	value := 5 // 定义一个整数变量
	fmt.Println("原始值：", value)

	addTen(&value) // 传递变量的地址
	fmt.Println("增加10后：", value)
}
