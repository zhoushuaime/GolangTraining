package main

func main() {
}

type IContext interface {
	Set(key string, value interface{})
}

type Con struct {
	Key   string
	Value interface{}
}

func (c *Con) Set(key string, value interface{}) {
	c.Key = key
	c.Value = value
}

func WithContext() *Con {
	return &Con{}
}
func NewContext() IContext {
	return WithContext()
}

