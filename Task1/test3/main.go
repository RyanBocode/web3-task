/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{} // 用来存放左括号

	// 括号对应关系
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch) // 左括号入栈
		case ')', ']', '}':
			// 栈空或栈顶不是匹配的左括号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1] // 弹出栈顶
		}
	}

	return len(stack) == 0 // 最终栈为空才有效
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([])"))   // true
	fmt.Println(isValid("([)]"))   // false
}
