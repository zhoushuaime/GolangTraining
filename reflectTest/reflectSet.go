package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u User) Hello(name string,age int) {
	 fmt.Println("hello my name is ",name," age is ",age)
}
func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("o is not pointer or can not set")
		return
	}
	fmt.Printf("%#v\n", v.Kind())
	v = v.Elem()

	if !v.FieldByName("1").IsValid() {
		fmt.Println("isvalid ")
		return
	}
	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("zhouwei")
	}
}

func main() {
	u := User{"zhangsan",20}
	//Set(&u)
	//fmt.Printf("%#v\n", u)
	//u.Hello("zhoushuai")
	v:=reflect.ValueOf(u)
	mv:=v.MethodByName("Hello")
	args:=[]reflect.Value{reflect.ValueOf("joe"),reflect.ValueOf(22)}
	mv.Call(args)
}
