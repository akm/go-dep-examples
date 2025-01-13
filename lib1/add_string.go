package lib1

import (
	"strconv"
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
	intRet := Add(inta, intb)
	return strconv.Itoa(intRet), nil
}

func init() {
	println("app1/add_string.go init")
}
