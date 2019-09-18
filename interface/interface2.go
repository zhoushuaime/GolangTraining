package main

import "fmt"

type I interface {
	walk()
}

type A struct {
}

func (a A) walk() {

}

type B struct {
}

func (b *B) walk() {

}

func main() {
	var a interface{}
	a = A{}
	_, ok := a.(I)
	a = &A{}
	_, ok = a.(I)
	fmt.Println(ok)

	a = &B{}
	_, ok = a.(I)
	fmt.Println(ok)

}

