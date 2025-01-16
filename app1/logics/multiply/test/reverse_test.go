package multiplytest

import (
	"testing"

	"app1/logics/multiply"
	multiplytesthelper "app1/logics/multiply/testhelper"
)

func TestReverse(t *testing.T) {
	for _, p := range multiplytesthelper.Patterns {
		t.Run(p.A+"*"+p.B, func(t *testing.T) {
			actual, err := multiply.Multiply(p.A, p.B)
			if err != nil {
				t.Errorf("Multiply(%s, %s) got error: %v", p.A, p.B, err)
			}
			if actual != p.Expected {
				t.Errorf("Multiply(%s, %s) = %s; want %s", p.A, p.B, actual, p.Expected)
			}
		})
	}
}

func init() {
	println("app1/logics/multiply/test/reverse_test.go init")
}
