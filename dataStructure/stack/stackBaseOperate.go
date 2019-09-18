package main

import (
	"errors"
	"fmt"
)

/*
@Time : 2018/1/13 20：13
@Author : joshua
@File : linklist_test
@Software: GoLand
*/


// Stacker ...
type Stacker interface {
	InitStack(s int)
	Push(data interface{}) error
	Pop() interface{}
	IsEmpty() bool
	IsFull() bool
	GetTop() (interface{}, error)
	ClearStack()
}

// Stack Stack
type Stack struct {
	Top       int
	StackSize int
	Elements  []interface{}
}

// InitStack 初始化栈
func (stack *Stack) InitStack(s int) {
	stack.StackSize = s
	stack.Elements = make([]interface{}, s)
	//stack.Elements = make([]interface{}, s)
	stack.Top = -1
}

// Push 入栈
func (s *Stack) Push(data interface{}) error {
	if s.IsFull() {
		return errors.New("stack is full")
	}
	s.Top++
	s.Elements[s.Top] = data
	return nil
}

// Pop 出栈
func (s *Stack) Pop() interface{} {
	t := s.Top
	s.Top--
	s.StackSize --
	return s.Elements[t]
}

// IsEmpey 检查栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

// IsFull 检查栈是否满了
func (s *Stack) IsFull() bool {
	return s.Top == s.StackSize-1
}

// GetTop 栈定元素
func (s *Stack) GetTop() (interface{}, error) {
	//if s.Top != -1 {
	//	return s.Elements[s.Top]
	//}
	if s.IsEmpty() {
		return s.Elements[s.Top], nil

	}
	//return -1
	return nil, errors.New("stack is empty")
}

// ClearStack 清空栈
func (s *Stack) ClearStack() {
	s.Top = -1
	s.Elements = s.Elements[0:0]
	s.StackSize = 0
}

// StackLen StackLen
func (s *Stack) StackLen() int {
	return len(s.Elements)
}
func main() {
	s := Stack{}
	s.InitStack(3)
	str := "ABCDEFG"

	for i := 0; i < len(str); i++ {
		if err := s.Push(string(str[i])); err != nil {
			fmt.Println("err:", err)
		}
	}
	//s.Push(10)
	//s.Push(20)
	//s.Push(30)
	//s.Push(40)
	//fmt.Println(len(s.Elements))
	//str := "[([][])]"
	//strs := strings.Split(str, "")
	//for _, v := range strs {
	//fmt.Print(v," ")
	//if (s.GetTop() == "[" && v == "]") || (s.GetTop() == "(" && v == ")") {
	//	s.Pop()
	//} else {
	//	s.Push(v)
	//	fmt.Print(s.GetTop(), " ")
	//}
	//}
	//fmt.Println()
	//fmt.Println(s.IsEmpey())
	//fmt.Println(s.Top)
	for !s.IsEmpty() {
		fmt.Print(s.Pop(), " ")
	}
	fmt.Println()
	fmt.Println(s.StackLen())
	s.ClearStack()
	fmt.Println("clear:", s.StackLen())
}
