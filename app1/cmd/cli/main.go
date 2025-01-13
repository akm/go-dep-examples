package main

import (
	"fmt"

	"github.com/akm/go-dep-examples/lib1"
)

func main() {
	fmt.Printf("%d + %d = %d\n", 1, 2, lib1.Add(1, 2))
}
