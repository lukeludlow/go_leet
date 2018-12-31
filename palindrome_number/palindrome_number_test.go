package goleet

import "testing"

func TestIsPalindrome(t *testing.T) {
	in := []int{
		121,
		-121,
		10,
	}
	expected := []bool{
		true,
		false,
		false,
	}
	numCases := 3
	for i := 0; i < numCases; i++ {
		actual := isPalindrome(in[i])
		if actual != expected[i] {
			t.Fatalf("case %d fails: %v\n", i, actual)
		}
	}
}
