package goleet

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	in := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{},
	}
	expected := []string{
		"fl",
		"",
		"",
	}
	numCases := 2
	for i := 0; i < numCases; i++ {
		actual := longestCommonPrefix(in[i])
		if actual != expected[i] {
			t.Fatalf("case %d: FAIL \nexpected: %v \nactual: %v", i, expected[i], actual)
		} else {
			fmt.Println("case ", i, " PASS")
		}
	}
}
