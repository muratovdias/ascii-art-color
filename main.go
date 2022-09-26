package main

import (
	"fmt"
	"student/domain"
)

func main() {
	result := domain.Print()
	if result != nil {
		fmt.Println(result)
	}
}
