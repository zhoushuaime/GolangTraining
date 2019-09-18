
package buildIgnore

import (
	"io"
)

/*
@Time : 2018/11/11 11:21 
@Author : zhoushuai
@File : colorWriter
@Software: GoLand
*/
type ColorWriter struct {
	w io.WriteCloser
}

type C interface {
	f()
}
func (c *C)f() {
	const (
		_ = iota
		Valid
		InValid
	)
}
func main(){
	var _ C_ = new (ColorWriter_)
}
