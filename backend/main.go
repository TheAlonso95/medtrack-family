package main

import (
	"fmt"
	"github.com/monorepo/backend/pkg/calculator"
)

func main() {
	result := calculator.Add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)
}