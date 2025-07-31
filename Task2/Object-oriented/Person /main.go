/*使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/

package main

import "fmt"

// 定义 Person 结构体
type Person struct {
	Name string
	Age  int
}

// 定义 Employee 结构体，组合 Person
type Employee struct {
	Person     // 结构体嵌入（组合）
	EmployeeID string
}

// 为 Employee 实现 PrintInfo 方法
func (e Employee) PrintInfo() {
	fmt.Println("员工信息：")
	fmt.Println("姓名：", e.Name)
	fmt.Println("年龄：", e.Age)
	fmt.Println("员工编号：", e.EmployeeID)
}

func main() {
	// 创建 Employee 实例
	emp := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: "E12345",
	}

	// 调用方法输出员工信息
	emp.PrintInfo()
}
