package multiply

import "fmt"

type Pattern struct {
	A        string
	B        string
	Expected string
}

var Patterns = []Pattern{
	{"1.0", "2.0", "2.000000"},
	{"1.1", "2.2", "2.420000"},
	{"-1.2", "2.3", "-2.760000"},
	{"-1.3", "-2.4", "3.120000"},
}

func init() {
	fmt.Println("app1/logics/multiply/multiply_patterns_test.go init")
}
