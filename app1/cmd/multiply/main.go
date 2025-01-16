package main

import (
	"fmt"
	"os"

	"app1/logics/multiply"
)

func main() {
	v1 := os.Args[1]
	v2 := os.Args[2]
	res, err := multiply.Multiply(v1, v2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}

func init() {
	println("app1/cmd/multiply/main.go init")
}
