package main

import (
	"fmt"
	"os"

	"github.com/akm/go-dep-examples/lib2"
)

func main() {
	v1 := os.Args[1]
	v2 := os.Args[2]
	res, err := lib2.SubtractString(v1, v2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}

func init() {
	println("app1/cmd/sub/main.go init")
}
