package app1

import (
	"strconv"

	"github.com/akm/go-dep-examples/lib1"
)

func AddString(a, b string) (string, error) {
	inta, err := strconv.Atoi(a)
	if err != nil {
		return "", err
	}
	intb, err := strconv.Atoi(b)
	if err != nil {
		return "", err
	}
	intRet := lib1.Add(inta, intb)
	return strconv.Itoa(intRet), nil
}

func init() {
	println("app1/add_string.go init")
}
