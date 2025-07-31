/*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数切片的指针作为参数，
在函数内部将该切片中的每个元素都乘以2，然后在主函数中调用该函数并输出修改后的切片。
考察点 ：指针的使用、值传递与引用传递的区别。
*/

package main

import "fmt"

// 定义函数，接收一个整数切片的指针
func doubleElements(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		(*nums)[i] *= 2 // 通过解引用修改切片中的元素
	}
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片：", slice)

	doubleElements(&slice) // 传入切片地址
	fmt.Println("元素乘2后：", slice)
}
