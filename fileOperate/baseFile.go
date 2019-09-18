package main

import (
	"io/ioutil"
)

func main(){
	//<-ch
	//file, err := ioutil.ReadFile("exit/exit.go")
	//bytes := make([]byte,1020)
	//f, err := os.Open("sync/Decnum.go")
	//check(err)
	//_, err = f.Read(bytes)
	//check(err)
	//fmt.Println(string(bytes))
	//if err != nil {
	//	fmt.Println("open file error")
	//	return
	//}
	//check(err)
	//fmt.Println(string(file))
	b:=[]byte("this is a test file")
	err:=ioutil.WriteFile("fileOperate/test.txt",b,0644)
	check(err)

}
func check(err error) {
	if err != nil {
		panic(err)
	}
}

