package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id, _ := GetUUID()
	fmt.Print(id)
}

// GetUUID ID生成
func GetUUID() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
