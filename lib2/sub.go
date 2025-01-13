package lib2

import (
	"github.com/akm/go-dep-examples/lib1"
)

func SubtractString(a, b string) (string, error) {
	return lib1.AddString(a, "-"+b)
}

func init() {
	println("lib2/sub.go init")
}
