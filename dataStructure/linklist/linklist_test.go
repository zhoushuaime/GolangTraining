package linklist

import (
	"fmt"
	"sync"
	"testing"
)

/*
@Time : 2018/1/12 17:55
@Author : joshua
@File : linklist_test
@Software: GoLand
*/

// LinkLister ...
type LinkLister interface {
	AddNode(Node)
	RemoveNode(Node)
	SetValue(interface{})
}

// Node ...
type Node struct {
	Data int
	Next *Node
	lock sync.Mutex
}

func (n *Node) SetValue(data int) {
	n.lock.Lock()
	defer n.lock.Unlock()
	n.Data = data
}

// AddNode ...
func (n *Node) AddNode() {

}
func (n *Node) GetValue()  interface{} {
	return n.Data
}

func TestLink(t *testing.T) {
	head := new(Node)
	head.SetValue(-1)
	head.Next = nil
	// 第一个节点
	firstNode := new(Node)
	firstNode.SetValue(0)
	firstNode.Next = nil
	head.Next = firstNode

	secord := new(Node)
	secord.SetValue(1)
	secord.Next = nil
	firstNode.Next = secord

	for head != nil {
		fmt.Println(head.GetValue())
		head = head.Next
	}

	m := make(map[string]string)
	m["one"] = "1"
	m["two"] = "2"
	m["three"] = "3"

	for k, _ := range m {
		switch k {
		case "one", "three":
			delete(m, k)

		}
	}
	fmt.Println("===========m=========")
	fmt.Println(m)

}

