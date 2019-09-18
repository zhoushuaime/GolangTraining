package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)
var uniqueId *snowflake.Node
var uniqueIdErr error

func init() {
	uniqueId, uniqueIdErr = snowflake.NewNode(1)
}

// NewUniqueId ...
func NewUniqueId() (string, error) {
	return fmt.Sprint(uniqueId.Generate()), uniqueIdErr
}