package goleet

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	in := []string{
		"III",
		"IV",
		"IX",
		"LVIII",
		"MCMXCIV",
	}
	expected := []int{
		3,
		4,
		9,
		58,
		1994,
	}
	numCases := 5
	for i := 0; i < numCases; i++ {
		actual := romanToInt(in[i])
		if actual != expected[i] {
			t.Fatalf("case %d: FAIL \nexpected: %v \nactual: %v", i, expected[i], actual)
		} else {
			fmt.Println("case ", i, " PASS")
		}
	}
}
