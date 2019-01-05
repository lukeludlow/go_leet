package goleet

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	in := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"]",
		"[",
		"[[[[]]]]",
		")))",
	}
	expected := []bool{
		true,
		true,
		false,
		false,
		true,
		false,
		false,
		true,
		false,
	}
	numCases := 9
	for i := 0; i < numCases; i++ {
		actual := isValid(in[i])
		if actual != expected[i] {
			t.Fatalf("case %d: FAIL \nexpected: %v \nactual: %v", i, expected[i], actual)
		} else {
			fmt.Println("case ", i, " PASS")
		}
	}
}
