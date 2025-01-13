package main

import (
	"app1"
	"fmt"
	"os"
)

func main() {
	v1 := os.Args[1]
	v2 := os.Args[2]
	res, err := app1.AddString(v1, v2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}

func init() {
	println("app1/cmd/cli/main.go init")
}
