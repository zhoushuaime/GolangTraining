package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"
)

var uniqueId *snowflake.Node

func init() {
	uniqueId, _ = snowflake.NewNode(1)
}
func NewUniqueId() string {
	return fmt.Sprint(uniqueId.Generate())
}
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(NewUniqueId())
	}
	fmt.Println(time.Now().Local())
	fmt.Println(time.Now())
}
