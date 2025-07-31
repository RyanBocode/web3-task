/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/

package main

import (
	"fmt"
	"math"
)

// 定义 Shape 接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义 Rectangle 结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle 实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle 实现 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 定义 Circle 结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle 实现 Shape 接口的 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// 创建 Rectangle 实例
	rect := Rectangle{Width: 4, Height: 5}
	// 创建 Circle 实例
	circ := Circle{Radius: 3}

	// 定义 Shape 类型变量并赋值
	var s Shape

	// 使用 Rectangle
	s = rect
	fmt.Println("矩形:")
	fmt.Printf("面积: %.2f\n", s.Area())
	fmt.Printf("周长: %.2f\n", s.Perimeter())

	// 使用 Circle
	s = circ
	fmt.Println("\n圆形:")
	fmt.Printf("面积: %.2f\n", s.Area())
	fmt.Printf("周长: %.2f\n", s.Perimeter())
}
